package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Weather Magic!\n")

	var cityVar string

	flag.StringVar(&cityVar, "city", "", "the city to get the weather forecast for")
	flag.Parse()

	fmt.Printf("Getting the weather for: %s\n", cityVar)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	weatherKey := os.Getenv("WEATHER_KEY")

	weatherApi := "https://api.openweathermap.org/data/2.5/weather?q=" + cityVar + "&appid=" + weatherKey

	req, err := http.NewRequest(http.MethodGet, weatherApi, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("UserAgent", "weather-magic")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP status: %s\n", res.Status)

	// welcome message - done
	// prompt for input (city name) - done
	// check- print input (city name) - done
	// parse input - done
	// set url for open weather api as a variable - done
	// request to api - done (200)
	// parse api response
	// return result
}
