package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("Recover",r)
		}
	}()

	fmt.Println("__Калькулятор индекса массы тела__")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)

		if err != nil {
			panic("Не заданы параметры для рассчета")
		}

		outputResult(IMT)

		repeatRes := repeatCalculation()
		if !repeatRes {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %v", imt)
	fmt.Println(result)
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	const IMTPower = 2

	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}

	IMT := userKg / math.Pow(userHeight/100, IMTPower)

	switch {
	case IMT < 16:
		fmt.Println("У вас недостаток веса")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыток веса")
	default:
		fmt.Println("У вас степень ожирения")
	}

	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return userKg, userHeight
}

func repeatCalculation() bool {
	var userChoise string
	fmt.Print("Хотите ввести еще один вариан? (y/n): ")
	fmt.Scan(&userChoise)

	if userChoise == "n" || userChoise == "y" {
		if userChoise == "n" {
			return false
		} else if userChoise == "y" {
			return true
		}
	} else {
		fmt.Print("Введите корректный варинт? (y/n): ")
		fmt.Scan(&userChoise)
		if userChoise == "n" {
			return false
		} else if userChoise == "y" {
			return true
		}
	}
	return false
}
