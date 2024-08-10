package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/jtatman/Cryptify/cryptify_app/authentication"
	"github.com/jtatman/Cryptify/cryptify_app/internal/trading"
	"github.com/jtatman/Cryptify/cryptify_app/internal/wallet"
	"github.com/jtatman/Cryptify/cryptify_app/web/chart"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read Binance API and secret keys from environment variables
	binanceAPIKey := getEnv("BINANCE_API_KEY", "")
	binanceAPISecret := getEnv("BINANCE_API_SECRET", "")

	// Create instances of authentication service, trading engine, and user wallet
	authService := authentication.NewAuthService()
	tradingEngine := trading.NewTradingEngine()
	userWallet := wallet.NewWallet()

	// Create instances of BinanceClient and BinanceProvider
	binanceClient := &chart.BinanceClient{
		APIKey:    binanceAPIKey,
		APISecret: binanceAPISecret,
	}

	binanceProvider := &chart.BinanceProvider{
		BinanceClient: binanceClient,
	}

	// Create instances of MockProvider and BinanceProvider
	mockProvider := &chart.MockProvider{}

	// Use BinanceProvider for real data or MockProvider for mock data
	// chartDataProvider := binanceProvider
	chartDataProvider := mockProvider

	// Specify the time range for fetching chart data
	startTime := time.Now().Add(-24 * time.Hour) // Example: 24 hours ago
	endTime := time.Now()

	// Fetch chart data for a specific cryptocurrency
	chartData, err := chartDataProvider.GetChartData("BTC", startTime, endTime)
	if err != nil {
		fmt.Printf("Error fetching chart data: %v\n", err)
		return
	}

	fmt.Printf("Chart data:\n%s\n", chartData)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
