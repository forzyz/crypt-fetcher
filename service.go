package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC":     50000,
	"ETH":     2000,
	"FORTHIS": 1000,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond)

	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker %s is not supported", ticker)
	}

	return price, nil
}
