package main

import (
	"fmt"
	"github.com/fatih/color"
	"passwordGenerator/account"
)

func main() {
	var vault = account.NewVault()
	fmt.Println("Менеджер паролей")

Menu:
	for {
		choice := printMenu()
		switch choice {
		case 1:
			createAccount(vault)
		case 2:
			FindAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu

		}
	}
}

func createAccount(vault *account.Vault) {
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

func FindAccount(vault *account.Vault) {
	url := promtData("Введите url: ")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунт не найден")
		return
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promtData("Введите URL аккаунта для удаления: ")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("✅ Аккаунт успешно удалён")
	} else {
		color.Red("❌ Аккаунт не найден")
	}
}
