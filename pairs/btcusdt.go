package pairs

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func BtcUsdt (kline *binance.WsKline) {
	fmt.Println("BTCUSDT here", kline)
}