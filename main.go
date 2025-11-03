package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult (imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %v", imt)
	fmt.Println(result)
}

func calculateIMT (userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput () (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return userKg, userHeight
}