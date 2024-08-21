package main

import (
	"flag"
	"fmt"
	"http/weatherApp/geo"
	"http/weatherApp/weather"
)

func main() {
	cityFromConfig := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	var city string

	flag.Parse()

	geoData, err := geo.GetMyLocation(*cityFromConfig)
	if *cityFromConfig != "" {
		city = *cityFromConfig
	} else {
		city = geoData.City
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	weatherNow := weather.GetWeather(*geoData, *format)
	fmt.Println("Weather in ", city, "is", weatherNow)
}
