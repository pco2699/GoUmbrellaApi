package main

import (
	"os"
	"umbrellaApi/Server"
	"umbrellaApi/WeatherApiClient"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	wac := WeatherApiClient.NewWeatherApiClient(apiKey)
	server := &Server.Server{
		wac,
	}
	server.Start()
}
