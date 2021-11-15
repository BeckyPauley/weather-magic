package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type temperature struct {
	Temp  float32 `json:"temp"`
	Feels float32 `json:"feels_like"`
}

type weather struct {
	//Id          string `json:"id"`
	Overview    string `json:"main"`
	Description string `json:"description"`
}

type response struct {
	Weather     []weather   `json:"weather"`
	Temperature temperature `json:"main"`
	Visibility  int         `json:"visibility"`
}

func main() {
	fmt.Println("Welcome to Weather Magic!\n")

	var cityVar string
	var weatherVar string

	flag.StringVar(&cityVar, "city", "", "the city to get the weather forecast for")
	flag.StringVar(&weatherVar, "info", "", "the weather forecast info you'd like to know")
	flag.Parse()

	fmt.Printf("Getting the current weather in %s\n\n", strings.Title(cityVar))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	weatherKey := os.Getenv("WEATHER_KEY")

	weatherUrl := "https://api.openweathermap.org/data/2.5/weather?q=" + cityVar + "&appid=" + weatherKey

	result, err := getForecast(weatherUrl)

	fmt.Printf("The %s in %s is:\n", weatherVar, strings.Title(cityVar))
	if weatherVar == "weather" {
		fmt.Printf("Overview: %+v\nDescription: %+v\n", result.Weather[0].Overview, result.Weather[0].Description)
	} else if weatherVar == "temperature" {
		fmt.Printf("Temperature: %+v\nFeels-like: %+ v\n", result.Temperature.Temp, result.Temperature.Feels)
	} else if weatherVar == "visibility" {
		fmt.Printf("Visibility: %+v\n", result.Visibility)
	}
	return
}

func getForecast(weatherUrl string) (response, error) {
	req, err := http.NewRequest(http.MethodGet, weatherUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("UserAgent", "weather-magic")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("HTTP status: %s\n", res.Status)

	w := response{}

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

	return w, err

}
