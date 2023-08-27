package pairs

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func XrpUsdt (kline *binance.WsKline) {
	fmt.Println("XRPUSDT here", kline)
}