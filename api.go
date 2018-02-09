package cryptomarket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// CryptoMoney is some datas about a specific currency.
type CryptoMoney struct {
	Rank       string `json:"rank"`
	Name       string `json:"name"`
	Price      string `json:"price_usd"`
	Total      string `json:"market_cap_usd"`
	ChangeWeek string `json:"percent_change_7d"`
	ChangeDay  string `json:"percent_change_24h"`
	ChangeHour string `json:"percent_change_1h"`
}

func (c CryptoMoney) String() string {
	return fmt.Sprintf("%s   %s   %s   %s   %s   %s   %s",
		fmtRank(c.Rank),
		fmtName(c.Name),
		fmtPrice(c.Price),
		fmtTotal(c.Total),
		fmtChange(c.ChangeWeek),
		fmtChange(c.ChangeDay),
		fmtChange(c.ChangeHour),
	)
}

const (
	colorRed     = "\x1b[31m"
	colorGreen   = "\x1b[32m"
	colorDefault = "\x1b[39m"
)

func fmtRank(v string) string {
	return fmt.Sprintf("%-4s", v)
}

func fmtName(v string) string {
	if len(v) > 20 {
		v = v[:20]
	}
	return fmt.Sprintf("%-20s", v)
}

func fmtPrice(v string) string {
	value, _ := strconv.ParseFloat(v, 64)
	return fmt.Sprintf("%-10.2f", value)
}

func fmtTotal(v string) string {
	value, _ := strconv.ParseFloat(v, 64)
	return fmt.Sprintf("%-15.0f", value)
}

func fmtChange(v string) string {
	color := colorGreen
	value, _ := strconv.ParseFloat(v, 64)
	if value < 0 {
		color = colorRed
	}
	return fmt.Sprintf("%s%-+10.2f%s", color, value, colorDefault)
}

// CryptoMarket is datas about some currency.
type CryptoMarket []CryptoMoney

// GetCryptoMarketDatas retrieve top crypto currency datas from coinmarketcap.
func GetCryptoMarketDatas(limit int) (CryptoMarket, error) {
	var cm CryptoMarket

	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/?limit=%d", limit)
	netclient := &http.Client{
		Timeout: time.Second * 5,
	}
	response, err := netclient.Get(url)
	if err != nil {
		return CryptoMarket{}, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&cm)
	if err != nil {
		return CryptoMarket{}, err
	}
	return cm, nil
}

func (c CryptoMarket) String() string {
	s := fmt.Sprintf(
		"%-4s   %-20s   %-10s   %-15s   %-10s   %-10s   %-10s%s",
		"", "Currency", "Price ($)", "Total ($)", "1 Week (%)", "1 Day (%)", "1 Hour (%)",
		"\n----   --------------------   ----------   ---------------   ----------   ----------   ----------",
	)
	for _, cm := range c {
		s += fmt.Sprintf("\n%s", cm.String())
	}
	return s
}
