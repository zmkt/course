package main

import (
	"fmt"
	"strings"
)

const (
	USD = 0.91
	RUB = 75.23
	EUR = 1.09
)

var currencyValues = map[string]float64{
	"USDT": USD,
	"RUB":  RUB,
	"EUR":  EUR,
}

func sourceCurrency() string {

	var sourceCurrency string = ""

	for {
		if sourceCurrency == "" {
			fmt.Println("Введите исходную валюту (USDT/RUB/EUR): ")
			fmt.Scan(&sourceCurrency)

			if sourceCurrency == "USDT" || sourceCurrency == "usdt" {
				sourceCurrency = "USDT"
			} else if sourceCurrency == "RUB" || sourceCurrency == "rub" {
				sourceCurrency = "RUB"
			} else if sourceCurrency == "EUR" || sourceCurrency == "eur" {
				sourceCurrency = "EUR"
			} else {
				sourceCurrency = ""
				continue
			}
		}

		return sourceCurrency
	}

}

func numberBills() int {

	var numberBills int = 0

	for {
		if numberBills == 0 {
			fmt.Println("Введите количество купюр: ")
			_, err := fmt.Scan(&numberBills)

			if err != nil {
				numberBills = 0
				continue
			}

			if numberBills > 0 {
				return numberBills
			} else {
				numberBills = 0
				continue
			}
		}

		return numberBills
	}
}

func availableTargetCurrencies(excludeCurrency string) string {
	currencies := []string{"USDT", "RUB", "EUR"}
	var filtered []string
	for _, c := range currencies {
		if c != excludeCurrency {
			filtered = append(filtered, c)
		}
	}
	return strings.Join(filtered, "/")
}

func targetCurrency(excludeCurrency string) string {
	var targetCurrency string = ""

	for {
		fmt.Printf("Введите целевую валюту (%s): ", availableTargetCurrencies(excludeCurrency))
		fmt.Scan(&targetCurrency)

		targetCurrency = strings.ToUpper(targetCurrency)

		if targetCurrency != excludeCurrency && (targetCurrency == "USDT" || targetCurrency == "RUB" || targetCurrency == "EUR") {
			return targetCurrency
		} else {
			targetCurrency = ""
			continue
		}
	}
}

func outputFunction(number int, currencyOne float64, currencyTwo float64) float64 {
	return float64(number) * currencyOne / currencyTwo
}

func inputFunction() {
	sourceCurrencyResult := sourceCurrency()
	numberBills := numberBills()
	targetCurrency := targetCurrency(sourceCurrencyResult)

	sourceValue := currencyValues[sourceCurrencyResult]
	targetValue := currencyValues[targetCurrency]

	result := outputFunction(numberBills, sourceValue, targetValue)

	fmt.Printf("Результат: %.2f\n", result)
}

func main() {

	inputFunction()
}
