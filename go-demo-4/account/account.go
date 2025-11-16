package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	Account
	cratedAt  time.Time
	updatedAt time.Time
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc)
}

func (acc *Account) GeneratePassword(n int) {
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

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("Неправильный login")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		fmt.Println("Ошибка")
		return nil, errors.New("Неправильный url")
	}

	newAcc := &AccountWithTimeStamp{
		cratedAt:  time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}

	if password == "" {
		newAcc.GeneratePassword(123)
	}

	return newAcc, nil

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
