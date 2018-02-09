package main

import (
	"flag"
	"fmt"

	"github.com/kahlys/cryptomarket"
)

var (
	fTop = flag.Int("top", 3, "top ranked money")
)

func main() {
	flag.Parse()

	market, err := cryptomarket.GetCryptoMarketDatas(*fTop)
	if err != nil {
		panic(err)
	}

	fmt.Println(market)
}
