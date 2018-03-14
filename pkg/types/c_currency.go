package types

import (
	"github.com/YuriBuerov/grpc-example/api"
)

type CCurrency struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Rank             int       `json:"rank,string"`
	PriceUSD         float64   `json:"price_usd,string"`
	PriceBTC         float64   `json:"price_btc,string"`
	PercentChange24H float64   `json:"percent_change_24h,string"`
}

type CCurrencies []CCurrency

func ToProtoCCurrency(c *CCurrency) *api.GetCCurrenciesResponse_CCurrency {
	return &api.GetCCurrenciesResponse_CCurrency{
		Id:          c.ID,
		Name:        c.Name,
		Symbol:      c.Symbol,
		Rank:        uint32(c.Rank),
		PriceUSD:    c.PriceUSD,
		DailyChange: c.PercentChange24H,
	}
}
