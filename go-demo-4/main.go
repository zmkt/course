package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	account
	cratedAt  time.Time
	updatedAt time.Time
}

func (acc account) outputPassword() {
	fmt.Println(acc)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// func newAccount(login, password, urlString string) (*account, error) {

// 	if login == "" {
// 		return nil, errors.New("Неправильный login")
// 	}

// 	_, err := url.ParseRequestURI(urlString)

// 	if err != nil {
// 		fmt.Println("Ошибка")
// 		return nil, errors.New("Неправильный url")
// 	}

// 	newAcc := &account{
// 		url:      urlString,
// 		login:    login,
// 		password: password,
// 	}

// 	if password == "" {
// 		newAcc.generatePassword(123)
// 	}

// 	return newAcc, nil

// }

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("Неправильный login")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		fmt.Println("Ошибка")
		return nil, errors.New("Неправильный url")
	}

	newAcc := &accountWithTimeStamp{
		cratedAt:  time.Now(),
		updatedAt: time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}

	if password == "" {
		newAcc.generatePassword(123)
	}

	return newAcc, nil

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func main() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url: ")

	myAccount, err := newAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Неправильный url или login")
		return
	}

	myAccount.generatePassword(12)

	myAccount.outputPassword()

}

func promtData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
