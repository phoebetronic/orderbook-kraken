package orderbook

import "encoding/json"

type Raw struct {
	A  [][]string `json:"a"`
	As [][]string `json:"as"`
	B  [][]string `json:"b"`
	Bs [][]string `json:"bs"`
	C  string     `json:"c"`
}

func (r Raw) Response() Response {
	var ask []Object
	var bid []Object

	for _, x := range r.A {
		ask = append(ask, Object{
			Price:     json.Number(x[0]),
			Volume:    json.Number(x[1]),
			Time:      json.Number(x[2]),
			Republish: republish(x),
		})
	}

	for _, x := range r.As {
		ask = append(ask, Object{
			Price:     json.Number(x[0]),
			Volume:    json.Number(x[1]),
			Time:      json.Number(x[2]),
			Republish: republish(x),
		})
	}

	for _, x := range r.B {
		bid = append(bid, Object{
			Price:     json.Number(x[0]),
			Volume:    json.Number(x[1]),
			Time:      json.Number(x[2]),
			Republish: republish(x),
		})
	}

	for _, x := range r.Bs {
		bid = append(bid, Object{
			Price:     json.Number(x[0]),
			Volume:    json.Number(x[1]),
			Time:      json.Number(x[2]),
			Republish: republish(x),
		})
	}

	return Response{
		Asks:       ask,
		Bids:       bid,
		CheckSum:   r.C,
		IsSnapshot: len(r.As) != 0 && len(r.Bs) != 0 && r.C == "",
	}
}

type Response struct {
	Asks       []Object
	Bids       []Object
	CheckSum   string
	IsSnapshot bool
}

type Object struct {
	Price     json.Number
	Volume    json.Number
	Time      json.Number
	Republish bool
}

func republish(raw []string) bool {
	if len(raw) == 4 && raw[3] == "r" {
		return true
	}

	return false
}
