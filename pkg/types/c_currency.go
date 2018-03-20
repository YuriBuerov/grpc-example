package types

import (
	"github.com/YuriBuerov/grpc-example/api"
)

// CCurrency object represent crypto currency from coinmarketcap API response
type CCurrency struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Rank             int       `json:"rank,string"`
	PriceUSD         float64   `json:"price_usd,string"`
	PriceBTC         float64   `json:"price_btc,string"`
	PercentChange24H float64   `json:"percent_change_24h,string"`
}

// CCurrencies set of CCurrency objects
type CCurrencies []CCurrency

// ToProtoCCurrency convert CCurrency to proto currency entry, take a look on api.proto
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
