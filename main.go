package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/forzyz/crypt-fetcher/client"
	"github.com/forzyz/crypt-fetcher/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "JSON API transport listen address")
		grpcAddr = flag.String("grpc", ":4000", "GRPC API transport listen address")
		svc      = NewLoggingService(NewMetricService(&priceFetcher{}))
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {		
		for {
			time.Sleep(3 * time.Second)
			// use whatever crypto ticker you want. Have fun with it!
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "ETH"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
