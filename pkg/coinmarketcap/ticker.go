package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/types"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const baseURL = "https://api.coinmarketcap.com/v1/ticker/"

// CTicker handler
type CTicker struct {
	logger  log.Logger
	client  *http.Client
	baseURL string
}

// NEwCTicker handler initializer
func NewCTicker(logger log.Logger, c *http.Client) *CTicker {
	return &CTicker{
		logger:  logger,
		client:  c,
		baseURL: baseURL,
	}
}

// GetCCurrencies handle func, which CTicker handler have to implement. (take a look on api.proto and api.pb.go)
func (t *CTicker) GetCCurrencies(ctx context.Context, in *api.GetCCurrenciesRequest) (*api.GetCCurrenciesResponse, error) {
	// Just simple functionality to show GPRC handler example
	var resp api.GetCCurrenciesResponse

	url := fmt.Sprintf("%s?limit=%d", t.baseURL, in.Limit)
	r, err := t.client.Get(url)
	if err != nil {
		t.logger.Log("error", err)
		return nil, status.Errorf(codes.Internal, "coinmarketcap api call error: %s", err.Error())
	}
	defer r.Body.Close()

	var cCurrencies types.CCurrencies
	if err := json.NewDecoder(r.Body).Decode(&cCurrencies); err != nil {
		t.logger.Log("error", err)
		return nil, status.Errorf(codes.Internal, "coinmarketcap api decode resp error: %s", err.Error())
	}

	resp.Currencies = make([]*api.GetCCurrenciesResponse_CCurrency, len(cCurrencies))
	for i, c := range cCurrencies {
		resp.Currencies[i] = types.ToProtoCCurrency(&c)
	}

	return &resp, nil
}
