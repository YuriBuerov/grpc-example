package main

import (
	"flag"
	"log"
	"os"
	"fmt"

	"github.com/YuriBuerov/grpc-example/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Minimalistic example of GRPC client
func main() {
	port := flag.String("port", "50051", "server port")
	limit := flag.Int("limit", 10, "limit")
	flag.Parse()
	if port == nil {
		log.Fatal("port is required")
		os.Exit(3)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", *port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to estabilish grpc connection")
		os.Exit(1)
	}

	client := api.NewApiClient(conn)
	defer conn.Close()

	req := &api.GetCCurrenciesRequest{Limit: uint32(*limit)}
	resp, err := client.GetCCurrencies(context.Background(), req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("Response---> %v\n", resp.Currencies)
}
