package main

import (
	"net"
	"os"

	"github.com/YuriBuerov/grpc-example/pkg/server"
	"github.com/go-kit/kit/log"
)

// We don't have config in this example, so port defined as constant
const (
	port = ":50051"
)

func main() {
	// Initialize logger
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller, "version", "1.0")

	// Recover from panic
	defer func() {
		if err := recover(); err != nil {
			logger.Log("error", "recover from panic", "cause", err)
			os.Exit(1)
		}
	}()

	// Initialize GRPC server
	grpcServer, err := server.NewGRPCServer(logger)
	if err != nil {
		logger.Log("error", "failed to init grpc server", "cause", err)
		os.Exit(1)
	}
	defer grpcServer.GracefulStop()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Log("error", "failed to listen", "cause", err)
		os.Exit(1)
	}

	logger.Log("event", "GRPC server started", "addr", port)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log("error", "GRPC server", "cause", err)
		os.Exit(1)
	}
}
