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
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{"1. Создать аккаунт", "2. Найти аккаунт по URL", "3. Найти аккаунт по логину", "4. Удалить аккаунт", "5. Выход", "Выберите вариант"}

func main() {

	vault := account.NewVault(files.NewJsonDB("data.json"))

	fmt.Println("__Менеджер паролей__")
Menu:
	for {
		variant := promptData(menuVariants...)
		menuFunction := menu[variant]
		if menuFunction == nil {
			break Menu
		}
		menuFunction(vault)
	}

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска: ")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска: ")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления: ")
	isDetected := vault.DeleteAccountByUrl(url)

	if isDetected {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Неправильный url или login")
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData(prompt ...string) string {
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
