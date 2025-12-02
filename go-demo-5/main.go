package main

import (
	"flag"
	"fmt"
	"go-demo-5/geo"
	"go-demo-5/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 16, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("geoData-->", geoData)

	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("weatherData-->", weatherData)

}
