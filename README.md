# weather-magic
### Want to check the weather? Do it straight from the comfort of your terminal with this handy CLI tool. 
</b>

Gives the current weather for you chosen city using the Open Weather API: https://openweathermap.org/current

## Installation
N/A- For future versions.

## Prerequisites
- An account for Open Weather API.
- .env file for api key

## Usage

```
go build
./weather-magic --city=[cityname] --info=[info]
```
e.g.: `./weather-magic --city=london --info=temperature`

Returns the current weather for your chosen location, e.g:
```
The weather in London is:
Overview: Clouds
Description: overcast clouds
```




