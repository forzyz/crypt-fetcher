package main

import (
	// "context"
	"flag"
	// "fmt"
	// "log"

	// "github.com/forzyz/crypt-fetcher/client"
)

func main() {
	// client test code (run this after running the service)
	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	// return
	listenAddr := flag.String("listenaddr", ":3000", "the address to listen on for HTTP requests")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
		
	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run() 
}
