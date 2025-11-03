package main

import (
	"fmt"
)

const (
	USDToEUR = 0.91
	USDToRUB = 75.23
	EURToRUB = USDToRUB / USDToEUR
)

	func inputFunction () float64 {
		var value float64;
		fmt.Println("Введите значение: ")

		fmt.Scan(&value)

		return value
	}

	func outputFunction (number float64, currencyOne float64, currencyTwo float64) float64{
		return number * currencyOne / currencyTwo
	}

	func main() {

		inputFunction()
		result := outputFunction(2, USDToEUR, USDToRUB)
		
		fmt.Println(result)

		fmt.Println("USD to EUR:", USDToEUR)
		fmt.Println("USD to RUB:", USDToRUB)
		fmt.Printf("EUR to RUB: %.2f\n", EURToRUB)
	}