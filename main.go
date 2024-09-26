package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the address to listen on for HTTP requests")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run() 
}
