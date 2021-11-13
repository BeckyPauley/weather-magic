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
./weather-magic --city=[cityname]
```
e.g.: `./weather-magic --city=london`

Returns overview and detail of the weather for your chosen location, e.g:
```
The weather in london is:
Overview: Clouds
Description: overcast clouds
```




