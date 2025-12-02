package weather_test

import (
	"fmt"
	"go-demo-5/geo"
	"go-demo-5/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Italia"
	geoData := geo.GeoData{
		City: expected,
	}

	format := 2

	res, err := weather.GetWeather(geoData, format)

	if err != nil {
		t.Error("Ошибка при получении данных")
	}

	if strings.Contains(res, expected) {
		fmt.Printf("Ожидаю %s, получаю %s", expected, string(res))
		t.Errorf("Ожидалось %v, получено %v", expected, string(res))
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 125},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			test := "Italiaqwe"
			geoData := geo.GeoData{
				City: test,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != nil {
				t.Errorf("Ожидалось %v, получено %v", weather.Error200, err)
			}
		})
	}
}
