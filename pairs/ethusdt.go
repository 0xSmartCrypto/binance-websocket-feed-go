package pairs

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func EthUsdt (kline *binance.WsKline) {
	fmt.Println("ETHUSDT here", kline)
}