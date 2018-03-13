package server

import (
	"runtime/debug"

	"github.com/YuriBuerov/grpc-example/api"
	"github.com/YuriBuerov/grpc-example/pkg/scanner"
	"github.com/go-kit/kit/log"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	*scanner.IPScanner
}

func NewGRPCServer(logger log.Logger) (*grpc.Server, error) {
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Log("error", "recovering from panic", "cause", p, "trace", string(debug.Stack()))
			return status.Errorf(codes.Internal, "%s", p)
		}),
	}

	ipScanner := scanner.NewIPScanner(logger)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	api.RegisterApiServer(s, &server{
		IPScanner: ipScanner,
	})

	return s, nil
}
