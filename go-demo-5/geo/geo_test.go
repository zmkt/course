package geo_test

import (
	"go-demo-5/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - Выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результата с expected
	if err != nil {
		t.Error("Ошибка получения города")
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}

}

func TestGetMyLocationNoCity(t *testing.T) {
	test := "Londonasd"
	_, err := geo.GetMyLocation(test)

	if err != geo.ErrorNoCity{
		t.Errorf("Ожидалось %v, получено %v", geo.ErrorNoCity, err)
	}
}
