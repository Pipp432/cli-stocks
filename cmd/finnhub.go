package cmd

import (
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

func getFinnhubClient() finnhub.DefaultApiService {

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("API_KEY"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	return *finnhubClient
}
