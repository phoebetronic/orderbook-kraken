package orderbook

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"sort"
	"strings"
	"sync"
	"time"
)

// https://docs.kraken.com/websockets/#message-book
type Orderbook struct {
	ask map[json.Number]json.Number
	bid map[json.Number]json.Number
	mut sync.Mutex
}

func New() *Orderbook {
	return &Orderbook{}
}

func (o *Orderbook) Checksum() string {
	var err error

	// Since the checksum calculation requires asks and bids to be sorted by
	// price, we need to allocate our internal state from a map to a slice of
	// price/volume pairs. Below we cast prices only once so that we do not have
	// to add additional compute during sorting.

	var ask []privol
	var bid []privol

	for k, v := range o.ask {
		var pri float64
		{
			pri, err = k.Float64()
			if err != nil {
				panic(err)
			}
		}

		{
			ask = append(ask, privol{Num: k, Flo: pri, Vol: v})
		}
	}

	for k, v := range o.bid {
		var pri float64
		{
			pri, err = k.Float64()
			if err != nil {
				panic(err)
			}
		}

		{
			bid = append(bid, privol{Num: k, Flo: pri, Vol: v})
		}
	}

	// As per the Kraken documentation, processing order is important. First,
	// the top ten ask price levels should be processed, sorted by price from
	// low to high. Then, the top ten bid price levels should be processed,
	// sorted by price from high to low.

	{
		sort.SliceStable(ask, func(i, j int) bool { return ask[i].Flo < ask[j].Flo })
		sort.SliceStable(bid, func(i, j int) bool { return bid[i].Flo > bid[j].Flo })
	}

	// The final concatenation adds the left-trimmed strings of price and volume
	// respectively, for each top pair from our sorted ask and bid slices. The
	// result is then hashed via CRC32, Golang specified using the IEEE
	// polynomial. Note that the code below removes non-essential order book
	// entries that got pushed out along the price levels. Without removing out
	// of scope price levels the checksum calculation does not work in its
	// current form. The implication here is that Orderbook.Checksum modifies
	// the internal order book state, instead of only reading from it. A cleaner
	// design could be achieved by sacrificing compute performance.

	var con string

	for i, x := range ask {
		if i <= 9 {
			con += trmlft(x.Num.String())
			con += trmlft(x.Vol.String())
		} else {
			delete(o.ask, x.Num)
		}
	}

	for i, x := range bid {
		if i <= 9 {
			con += trmlft(x.Num.String())
			con += trmlft(x.Vol.String())
		} else {
			delete(o.bid, x.Num)
		}
	}

	return fmt.Sprintf("%d", crc32.ChecksumIEEE([]byte(con)))
}

func (o *Orderbook) Empty() bool {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	return len(o.ask) == 0 && len(o.bid) == 0
}

func (o *Orderbook) MarshalJSON() ([]byte, error) {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	return json.Marshal(&struct {
		Ask map[json.Number]json.Number `json:"ask"`
		Bid map[json.Number]json.Number `json:"bid"`
		Tim time.Time                   `json:"tim"`
	}{
		Ask: o.ask,
		Bid: o.bid,
		Tim: time.Now().UTC().Round(time.Second),
	})
}

func (o *Orderbook) Middleware(upd Response) error {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	if upd.IsSnapshot {
		o.Snapshot(upd)
	} else {
		{
			o.Update(upd)
		}

		if upd.CheckSum != o.Checksum() {
			return fmt.Errorf("current order book checksum (%s) must match desired order book checksum (%s)", o.Checksum(), upd.CheckSum)
		}
	}

	return nil
}

func (o *Orderbook) Snapshot(upd Response) {
	{
		o.ask = map[json.Number]json.Number{}
		o.bid = map[json.Number]json.Number{}
	}

	for _, x := range upd.Asks {
		o.ask[x.Price] = x.Volume
	}

	for _, x := range upd.Bids {
		o.bid[x.Price] = x.Volume
	}
}

func (o *Orderbook) Update(upd Response) {
	var err error

	for _, x := range upd.Asks {
		var vol float64
		{
			vol, err = x.Volume.Float64()
			if err != nil {
				panic(err)
			}
		}

		if vol == 0 {
			delete(o.ask, x.Price)
		} else {
			o.ask[x.Price] = x.Volume
		}
	}

	for _, x := range upd.Bids {
		var vol float64
		{
			vol, err = x.Volume.Float64()
			if err != nil {
				panic(err)
			}
		}

		if vol == 0 {
			delete(o.bid, x.Price)
		} else {
			o.bid[x.Price] = x.Volume
		}
	}
}

type privol struct {
	Num json.Number
	Flo float64
	Vol json.Number
}

func trmlft(str string) string {
	x := strings.TrimLeft(strings.Replace(str, ".", "", -1), "0")
	return x
}
