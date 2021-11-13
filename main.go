package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "flags"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Weather Magic!\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the place you'd like to find the weather forecast for:\n")
	city, _ := reader.ReadString('\n')
	fmt.Println(city)

	// weatherKey := os.Getenv("WEATHER_KEY")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	testEnv := os.Getenv("TEST_ENV")
	fmt.Println(testEnv)

	// weatherApi = api.openweathermap.org/data/2.5/weather?q={city name}&appid={API key}

	// welcome message - done
	// prompt for input (city name) - done
	// check- print input (city name) - done
	// parse input
	// set url for open weather api as a variable
	// request to api
	// parse api response
	// return result
}
