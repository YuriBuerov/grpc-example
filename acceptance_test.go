package grpc_example_test

import (
	"net"
	"os"
	"testing"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/server"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestGetCCurrencies(t *testing.T) {
	addr, cancel := tServer(t)
	defer cancel()
	c, conn := tClient(t, addr)
	defer conn.Close()

	req := &api.GetCCurrenciesRequest{Limit: 3}
	resp, err := c.GetCCurrencies(context.Background(), req)

	require.Nil(t, err)
	require.NotNil(t, resp)
}

func tServer(t *testing.T) (string, func()) {
	logger := log.With(log.NewJSONLogger(os.Stdout), "caller", log.DefaultCaller)

	s, err := server.NewGRPCServer(logger)
	require.Nil(t, err)
	l, err := net.Listen("tcp", "localhost:0")
	require.Nil(t, err)

	go s.Serve(l)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()

	return l.Addr().String(), cancel
}

func tClient(t *testing.T, addr string) (api.ApiClient, *grpc.ClientConn) {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	require.Nil(t, err)

	return api.NewApiClient(c), c
}
