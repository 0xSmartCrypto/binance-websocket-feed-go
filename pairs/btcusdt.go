package pairs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/0xSmartCrypto/binance-websocket-feed-go/db"
	"github.com/0xSmartCrypto/binance-websocket-feed-go/features"
	"github.com/adshao/go-binance/v2"
)

type BtcFeature interface {
	features.VolumeMA | features.SMA
}

func BtcUsdt (
	ctx context.Context, 
	kline *binance.WsKline, 
	baseMA *features.VolumeMA, 
	quoteMA *features.VolumeMA,
	client *db.PrismaClient,
) {
	// Find or create pair
	pairId, err := findOrCreatePair(ctx, client, db.ValidPairBtcusdt)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Record OHLCV
	start := time.UnixMilli(kline.StartTime)
	open, _ := strconv.ParseFloat(kline.Open, 64)
	high, _ := strconv.ParseFloat(kline.High, 64)
	low, _ := strconv.ParseFloat(kline.Low, 64)
	close, _ := strconv.ParseFloat(kline.Close, 64)
	baseVolume, _ := strconv.ParseFloat(kline.Volume, 64)
	quoteVolume, _ := strconv.ParseFloat(kline.QuoteVolume, 64)
	numberTrades := kline.TradeNum
	takerBuyBaseVolume, _ := strconv.ParseFloat(kline.ActiveBuyVolume, 64)
	takerBuyQuoteVolume, _ := strconv.ParseFloat(kline.ActiveBuyQuoteVolume, 64)

	// Feature: Calculate Volume Moving Average
	baseMA.Add(baseVolume)
	quoteMA.Add(quoteVolume)
	
	meta := map[string]float64{
		"BaseVolumeMA": baseMA.Value(),
		"QuoteVolumeMA": quoteMA.Value(),
	}
	metaJson, err := json.Marshal(meta)
	if err != nil {
		fmt.Println(err)
		return
	}

	created, err := client.Kline.CreateOne(
		db.Kline.Pair.Link(
			db.Pair.ID.Equals(pairId),
		),
		db.Kline.Start.Set(start),
		db.Kline.Open.Set(open),
		db.Kline.High.Set(high),
		db.Kline.Low.Set(low),
		db.Kline.Close.Set(close),
		db.Kline.Trades.Set(int(numberTrades)),
		db.Kline.BaseVolume.Set(baseVolume),
		db.Kline.QuoteVolume.Set(quoteVolume),
		db.Kline.TakerBuyBaseVolume.Set(takerBuyBaseVolume),
		db.Kline.TakerBuyQuoteVolume.Set(takerBuyQuoteVolume),
		db.Kline.Meta.Set(metaJson),
	).Exec(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Print("Kline created: ", created.ID)
}

func findOrCreatePair(ctx context.Context, client *db.PrismaClient, symbol db.ValidPair) (string, error) {
	pair, err := client.Pair.FindFirst(
		db.Pair.Symbol.Equals(symbol),
	).Exec(ctx)

	if pair != nil {
		return pair.ID, nil
	}

	created, _ := client.Pair.CreateOne(
		db.Pair.Symbol.Set(symbol),
		db.Pair.Base.Set("BTC"),
		db.Pair.Quote.Set("USDT"),
	).Exec(ctx)

	if created == nil {
		return "0", err
	}

	return created.ID, nil
}