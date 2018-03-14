package server

import (
	"net/http"
	"runtime/debug"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/coinmarketcap"
	"github.com/go-kit/kit/log"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	*coinmarketcap.CTicker
}

func NewGRPCServer(logger log.Logger) (*grpc.Server, error) {
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Log("error", "recovering from panic", "cause", p, "trace", string(debug.Stack()))
			return status.Errorf(codes.Internal, "%s", p)
		}),
	}

	cTicker := coinmarketcap.NewCTicker(logger, http.DefaultClient)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	api.RegisterApiServer(s, &server{
		CTicker: cTicker,
	})

	return s, nil
}
