package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/phoebetronic/orderbook-kraken/pkg/orderbook"
	"github.com/sacOO7/gowebsocket"
)

func main() {
	fmt.Println()
	fmt.Println("|===========================================|")
	fmt.Println("|            wss://ws.kraken.com            |")
	fmt.Println("|===========================================|")
	fmt.Println()

	{
		api := "wss://ws.kraken.com"
		sub := "{ \"event\":\"subscribe\", \"subscription\":{\"name\":\"book\"},\"pair\":[\"ETH/USD\"] }"

		OpenAndStreamWebSocketSubscription(api, sub)
	}

	fmt.Println()
	fmt.Println("|===========================================|")
	fmt.Println("|            wss://ws.kraken.com            |")
	fmt.Println("|===========================================|")
	fmt.Println()
}

func OpenAndStreamWebSocketSubscription(api, sub string) {
	var cli gowebsocket.Socket
	{
		cli = gowebsocket.New(api)
	}

	var obk *orderbook.Orderbook
	{
		obk = orderbook.New()
	}

	var sig chan os.Signal
	{
		sig = make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
	}

	cli.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Received connect error - ", err)
	}

	cli.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to server")
	}

	cli.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		var err error

		{
			eth := strings.Contains(message, "ETH/USD")
			snp := (strings.Contains(message, `"as"`) || strings.Contains(message, `"bs"`))
			upd := (strings.Contains(message, `"a"`) || strings.Contains(message, `"b"`))

			if !(eth && (snp || upd)) {
				return
			}
		}

		var raw orderbook.Raw
		{
			message = strings.TrimPrefix(message, `[560,`)
			message = strings.TrimSuffix(message, `,"book-10","ETH/USD"]`)

			err = json.Unmarshal([]byte(message), &raw)
			if err != nil {
				panic(err)
			}
		}

		{
			err = obk.Middleware(raw.Response())
			if err != nil {
				panic(err)
			}
		}

		var byt []byte
		{
			byt, err = json.Marshal(obk)
			if err != nil {
				panic(err)
			}
		}

		{
			fmt.Printf("%s\n", byt)
		}
	}

	cli.OnPingReceived = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Received ping - " + message)
	}

	cli.OnPongReceived = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Received pong - " + message)
	}

	cli.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Socket Closed")
	}

	{
		cli.Connect()
		cli.SendText(sub)
	}

	for range sig {
		cli.Close()
		break
	}
}
