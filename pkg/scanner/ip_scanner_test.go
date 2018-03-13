package scanner_test

import (
	"context"
	"os"
	"testing"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/scanner"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanIPAddr(t *testing.T) {
	ipScanner := scanner.NewIPScanner(log.NewLogfmtLogger(os.Stdout))

	res, err := ipScanner.ScanIPAddr(context.Background(), &api.ScanIPRequest{IpAddr: "stackoverflow.com"})

	require.NoError(t, err)
	assert.Equal(t, "151.101.129.69", res.DomainName)
}
