package scanner

import (
	"github.com/YuriBuerov/grpc-example/api"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

type IPScanner struct {
	logger log.Logger
}

func NewIPScanner(logger log.Logger) *IPScanner {
	return &IPScanner{
		logger: logger,
	}
}

func (s *IPScanner) ScanIPAddr(ctx context.Context, in *api.ScanIPRequest) (*api.ScanIPResponse, error) {
	return nil, nil
}
