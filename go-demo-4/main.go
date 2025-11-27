package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/files"
	"go-demo-4/output"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {

	vault := account.NewVault(files.NewJsonDB("data.json"))

	fmt.Println("__Менеджер паролей__")
Menu:
	for {
		variant := promptData([]string{"1. Создать аккаунт", "2. Найти аккаунт", "3. Удалить аккаунт", "4. Выход", "Выберите вариант"})
		menuFunction := menu[variant]
		if menuFunction == nil {
			break Menu
		}
		menuFunction(vault)
	}

}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска: "})
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
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
