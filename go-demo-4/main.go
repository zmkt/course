package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/files"
)

func main() {

	fmt.Println("__Менеджер паролей__")
	for {
		variant := getMenu()

	Menu:
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}

}

func getMenu() int {

	var variant int

	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	fmt.Scan(&variant)
	return variant

}

func findAccount() {

}

func deleteAccount() {

}

func createAccount() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Неправильный url или login")
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}

	files.WriteFile(data, "data.json")
}

func promtData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
