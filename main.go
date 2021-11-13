package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Weather Magic!\n")

	var cityVar string

	flag.StringVar(&cityVar, "city", "", "the city to get the weather forecast for")
	flag.Parse()

	fmt.Printf("Getting the weather for: %s\n", cityVar)

	// welcome message - done
	// prompt for input (city name) - done
	// check- print input (city name) - done
	// parse input - done
	// set url for open weather api as a variable
	// request to api
	// parse api response
	// return result
}
