package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/files"
)

func main() {
	files.ReadFile()
	files.WriteFile("Привет! Я файл", "file.txt")
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Неправильный url или login")
		return
	}

	myAccount.GeneratePassword(12)

	myAccount.OutputPassword()
}

func promtData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
