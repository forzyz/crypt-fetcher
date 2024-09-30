package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// API URL and headers for CoinMarketCap
const CoinMarketCapURL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return FetchPriceFromAPI(ctx, ticker)
}

// CoinMarketCap response structures
type CoinMarketCapResponse struct {
	Data map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

func FetchPriceFromAPI(ctx context.Context, ticker string) (float64, error) {
    err := godotenv.Load()
    if err != nil {
        return 0, fmt.Errorf("error loading .env file: %v", err)
    }

	CoinMarketCapAPIKey := os.Getenv("COINMARKETCAP_API_KEY")
	if CoinMarketCapAPIKey == "" {
		return 0, fmt.Errorf("COINMARKETCAP_API_KEY environment variable not set")
	}

	client := &http.Client{Timeout: 10 * time.Second}

	// Create a new request with the appropriate query parameters
	req, err := http.NewRequestWithContext(ctx, "GET", CoinMarketCapURL, nil)
	if err != nil {
		return 0, err
	}

	// Add query parameters and headers to the request
	query := req.URL.Query()
	query.Add("symbol", ticker)
	query.Add("convert", "USD") // Convert to USD price
	req.URL.RawQuery = query.Encode()


	req.Header.Add("X-CMC_PRO_API_KEY", CoinMarketCapAPIKey)
	req.Header.Add("Accepts", "application/json")

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error fetching price: status %s", resp.Status)
	}

	// Parse the response JSON
	var cmcResponse CoinMarketCapResponse
	if err := json.NewDecoder(resp.Body).Decode(&cmcResponse); err != nil {
		return 0, err
	}

	// Extract the price from the response
	price := cmcResponse.Data[ticker].Quote["USD"].Price

	if price <= 0 {
		return 0, fmt.Errorf("invalid price: %f", price)
	}

	return price, nil
}
