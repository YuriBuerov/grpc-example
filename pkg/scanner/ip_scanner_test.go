package scanner_test

import (
	"context"
	"os"
	"testing"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/scanner"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/require"
)

func TestScanIPAddr(t *testing.T) {
	ipScanner := scanner.NewIPScanner(log.NewLogfmtLogger(os.Stdout))

	res, err := ipScanner.ScanIPAddr(context.Background(), &api.ScanIPRequest{DomainName: "stackoverflow.com"})

	checkIPs := func(ips []*api.ScanIPResponse_ScanIPEntry) bool {
		for _, ip := range ips {
			if ip.IpAddr == "151.101.1.69" {
				return true
			}
		}
		return false
	}

	require.NoError(t, err)
	require.True(t, checkIPs(res.IpAddresses))
}
