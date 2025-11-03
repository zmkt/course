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

	func outputFunction (numberOne float64, numberTwo float64) {
		fmt.Println(numberOne, numberTwo)
	}

	func main() {

		inputFunction()

		fmt.Println("USD to EUR:", USDToEUR)
		fmt.Println("USD to RUB:", USDToRUB)
		fmt.Printf("EUR to RUB: %.2f\n", EURToRUB)
	}