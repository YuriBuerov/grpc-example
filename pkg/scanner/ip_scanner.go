package scanner

import (
	"net"

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
	var resp api.ScanIPResponse

	addr, err := net.LookupIP(in.DomainName)
	if err != nil {
		s.logger.Log("error", err)
		return nil, err
	}

	resp.IpAddresses = make([]*api.ScanIPResponse_ScanIPEntry, len(addr))
	for i, v := range addr {
		resp.IpAddresses[i] = &api.ScanIPResponse_ScanIPEntry{IpAddr: v.String()}
	}

	return &resp, nil
}
