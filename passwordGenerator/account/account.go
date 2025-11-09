package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	if login == "" {
		return nil, errors.New("Blank login field")
	}
	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}

func (acc *Account) OutputPassword() {
	fmt.Print(acc.Login, " ", acc.Password, " ", acc.Url)

}

func (acc *Account) generatePassword(n int) {
	passwd := make([]rune, n)
	for i := range passwd {
		passwd[i] = runes[rand.Intn(len(runes))]
	}
	acc.Password = string(passwd)
}

func (acc *Account) ToBites() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}
