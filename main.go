//go:generate go run github.com/steebchen/prisma-client-go generate

package main

import (
	"fmt"

	"example.com/trading/pairs"
	"github.com/adshao/go-binance/v2"
)

type ValidPair string

const (
	BTCUSDT ValidPair = "BTCUSDT"
	ETHUSDT ValidPair = "ETHUSDT"
	XRPUSDT ValidPair = "XRPUSDT"
)

func main() {
	handler := func(event *binance.WsKlineEvent) {
		if (event.Kline.IsFinal) {
			switch (event.Symbol) {
				case string(BTCUSDT):
					pairs.BtcUsdt(&event.Kline)
				case string(ETHUSDT):
					pairs.EthUsdt(&event.Kline)
				case string(XRPUSDT):
					pairs.XrpUsdt(&event.Kline)
			}
		}
	}

	errHandler := func(err error) {
		fmt.Println(err)
	}

	fmt.Println("Subscribing to Binance Websocket Feed ...")

	doneC, _, err := binance.WsKlineServe(string(BTCUSDT), "1m", handler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
