package main

import (
	"fmt"
	"github.com/fatih/color"
	"passwordGenerator/account"
)

var vault = account.NewVault()

func main() {
	fmt.Println("Менеджер паролей")
Menu:
	for {
		choice := printMenu()
		switch choice {
		case 1:
			createAccount()
		case 2:
			FindAccount()
		case 3:
			deleteAccount()
		case 4:
			break Menu

		}
	}
}

func createAccount() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url:  ")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid login or url")
		return
	}
	vault.AddAccount(*myAccount)
}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

func printMenu() int {
	var choice int
	fmt.Println("Введите пункт меню: ")
	fmt.Println("1. Создать аккаунт\n2. Найти аккаунт\n3. Удалить аккаунт\n4. Выйти")
	fmt.Scanln(&choice)
	return choice
}

func FindAccount() {
	url := promtData("Введите url: ")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		//fmt.Println("Аккаунт не найден")
		color.Red("Аккаунт не найден")
		return
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount() {
}
