package main

import (
	"fmt"
)

type stringMap = map[string]string

func inputArea() string {
	fmt.Println("Выберите действие:")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")

	var input string

	fmt.Scan(&input)

	return input
}

func one(result stringMap) {
	fmt.Println("----------------------")
	fmt.Println("")
	for key, item := range result {
		fmt.Println(key, ":", item)
	}
	fmt.Println("")
	fmt.Println("----------------------")
}

func two(result stringMap) {

	var name string
	var value string

	fmt.Print("Введите название: ")

	fmt.Scan(&name)

	fmt.Print("Введите значение: ")

	fmt.Scan(&value)

	result[name] = value

	fmt.Println("----------------------")
	fmt.Println("")
	fmt.Println("Значение добавлено")
	fmt.Println("")
	fmt.Println("----------------------")
}

func three(result stringMap) {

	fmt.Print("Введите название: ")

	var input string

	fmt.Scan(&input)

	delete(result, input)

	fmt.Println("----------------------")
	fmt.Println("")
	fmt.Println("Значение удалено")
	fmt.Println("")
	fmt.Println("----------------------")
}

func main() {

	result := map[string]string{
		"one":  "one",
		"twoo": "twoo",
	}
Menu:
	for {
		input := inputArea()
		switch input {
		case "1":
			one(result)
		case "2":
			two(result)
		case "3":
			three(result)
		case "4":
			break Menu
		}
	}

}
