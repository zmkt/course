package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/files"
	"go-demo-4/output"

	"github.com/fatih/color"
)

func main() {

	vault := account.NewVault(files.NewJsonDB("data.json"))

	fmt.Println("__Менеджер паролей__")
	for {
		variant := promptData([]string{"1. Создать аккаунт", "2. Найти аккаунт", "3. Удалить аккаунт", "4. Выход", "Выберите вариант"})

	Menu:
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска: "})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления: "})
	isDetected := vault.DeleteAccountByUrl(url)

	if isDetected {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин: "})
	password := promptData([]string{"Введите пароль: "})
	url := promptData([]string{"Введите url: "})

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Неправильный url или login")
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
