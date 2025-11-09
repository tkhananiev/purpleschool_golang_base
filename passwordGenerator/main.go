package main

import (
	"fmt"
	"passwordGenerator/account"
	"passwordGenerator/files"
)

func main() {
	createAccount()
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
	file, err := myAccount.ToBites()
	if err != nil {
		fmt.Println("Не удалось сохранить в JSON")
		return
	}
	files.WriteFile(file, "data.json")
	fmt.Println(string(file))
}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}
