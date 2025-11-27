package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func renderNumbers() ([]int, error) {

	var numbers []int

	reader := bufio.NewReader(os.Stdin)
	numbersStr, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	numbersStr = strings.TrimSpace(numbersStr)

	parse := strings.Split(numbersStr, ",")

	for _, value := range parse {
		numberStr := strings.TrimSpace(value)
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("Ошибка2")
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func AVG(numbers *[]int) float64 {
	var length int = len(*numbers)
	var sum int = 0
	for _, value := range *numbers {
		sum += value
	}
	return float64(sum) / float64(length)
}

func SUM(numbers *[]int) int {
	var sum = 0

	for _, value := range *numbers {
		sum += value
	}

	return sum
}

func main() {

	var operation string
	var numbers []int

	fmt.Print("Введите операцию (AVG/SUM/MED): ")

	fmt.Scan(&operation)

	fmt.Print("Введите числа через запятую: ")

	numbers, err := renderNumbers()

	if err != nil {
		fmt.Println("Ошибка2")
		return
	}

	operation = strings.ToLower(operation)

	medFloat64 := func(numbers *[]int) float64 {
		n := len(*numbers)
		if n%2 == 1 {
			return float64((*numbers)[n/2])
		} else {
			return float64((*numbers)[n/2-1]+(*numbers)[n/2]) / 2.0
		}
	}

	operations := map[string]func(*[]int) float64{
		"avg": AVG,
		"sum": func(nums *[]int) float64 { return float64(SUM(nums)) },
		"med": medFloat64,
	}

	if fn, ok := operations[operation]; ok {
		result := fn(&numbers)
		fmt.Println(result)
	} else {
		fmt.Println("Неправильный оператор")
	}
	

}
