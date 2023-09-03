//go:generate go run github.com/steebchen/prisma-client-go generate

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/0xSmartCrypto/binance-websocket-feed-go/db"
	"github.com/0xSmartCrypto/binance-websocket-feed-go/features"
	"github.com/0xSmartCrypto/binance-websocket-feed-go/pairs"
	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

type ValidPair string

const (
	BTCUSDT ValidPair = "BTCUSDT"
	ETHUSDT ValidPair = "ETHUSDT"
	XRPUSDT ValidPair = "XRPUSDT"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Environment", os.Getenv("ENV"))

	client := db.NewClient()
	if err:= client.Prisma.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
    if err := client.Prisma.Disconnect(); err != nil {
      panic(err)
    }
  }()
 
  ctx := context.Background()

	// Feature: Calculate Volume Moving Average
	btcBaseMA := features.VolumeMA{
		Period: 8,
	}
	btcQuoteMA := features.VolumeMA{
		Period: 8,
	}
	
	handler := func(event *binance.WsKlineEvent) {
		if (event.Kline.IsFinal) {
			switch (event.Symbol) {
				case string(BTCUSDT):
					pairs.BtcUsdt(ctx, &event.Kline, &btcBaseMA, &btcQuoteMA, client)
				// case string(ETHUSDT):
				// 	pairs.EthUsdt(&event.Kline)
				// case string(XRPUSDT):
				// 	pairs.XrpUsdt(&event.Kline)
			}
		}
	}

	errHandler := func(err error) {
		fmt.Println(err)
	}

	fmt.Println("Subscribing to Binance Websocket Feed ...")
	
	symbolIntervalMap := make(map[string]string)
	symbolIntervalMap[string(BTCUSDT)] = "1m"
	// symbolIntervalMap[string(ETHUSDT)] = "1m"
	// symbolIntervalMap[string(XRPUSDT)] = "1m"

	doneC, _, err := binance.WsCombinedKlineServe(
		symbolIntervalMap,
		handler, 
		errHandler,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC

	
}
