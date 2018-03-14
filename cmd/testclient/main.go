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

func main() {
	port := flag.String("port", "50051", "server port")
	domainName := flag.String("domain-name", "stackoverflow.com", "domain name")
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

	req := &api.ScanIPRequest{DomainName: *domainName}
	resp, err := client.ScanIPAddr(context.Background(), req)

	log.Printf("Response---> %v\n", resp.IpAddresses)
}
