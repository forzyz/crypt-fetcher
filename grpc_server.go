package main

import (
	"context"
	"math/rand"
	"net"

	"github.com/forzyz/crypt-fetcher/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcherServer(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)
	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)
	
	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcherServer(svc PriceFetcher) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (f *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	reqId := rand.Intn(1000000)
	ctx = context.WithValue(ctx, "requestID", reqId)
	price, err := f.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}

	return resp, nil
}
