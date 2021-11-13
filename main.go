package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

type name struct {
	Name string `json:"name"`
}

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

	w := name{}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("unable to parse value: %q, error: %s",
			string(body), err.Error())
	}

	fmt.Println(w)

	// welcome message - done
	// prompt for input (city name) - done
	// check- print input (city name) - done
	// parse input - done
	// set url for open weather api as a variable - done
	// request to api - done (200)
	// parse api response - basic done
	// return result
}
