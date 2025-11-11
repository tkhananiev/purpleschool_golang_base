package main

import (
	"fmt"
	"passwordGenerator/account"
)

func main() {
	fmt.Println("Менеджер паролей")
Menu:
	for {
		choice := printMenu()
		switch choice {
		case 1:
			createAccount()
		case 2:
			FindAccountByLogin()
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
	var vault = account.NewVault()
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

func FindAccountByLogin() {
}
func deleteAccount() {
}
