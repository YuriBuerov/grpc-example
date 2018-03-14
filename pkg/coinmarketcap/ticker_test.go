package coinmarketcap_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/coinmarketcap"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/require"
)

func TestGetCCurrencies(t *testing.T) {
	cTicker := coinmarketcap.NewCTicker(log.NewLogfmtLogger(os.Stdout), http.DefaultClient)

	res, err := cTicker.GetCCurrencies(context.Background(), &api.GetCCurrenciesRequest{Limit: 3})

	check := func(currencies []*api.GetCCurrenciesResponse_CCurrency) bool {
		for _, c := range currencies {
			if c.Symbol == "BTC" {
				return true
			}
		}
		return false
	}

	require.NoError(t, err)
	require.True(t, check(res.Currencies))
}
