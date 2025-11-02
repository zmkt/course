package main

import (
	"fmt"
)

const (
	USDToEUR = 0.91
	USDToRUB = 75.23
	EURToRUB = USDToRUB / USDToEUR
)

func main() {
	fmt.Println("USD to EUR:", USDToEUR)
	fmt.Println("USD to RUB:", USDToRUB)
	fmt.Printf("EUR to RUB: %.2f\n", EURToRUB)
}