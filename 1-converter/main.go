package main

import (
	"fmt"
	"strings"
)

const (
	usdToEur = 0.85
	usdToRub = 80.5
	eurToRub = usdToRub / usdToEur
)

func main() {

	originalCurrency, quantity, finalCurrency := userInput()
	fmt.Println(originalCurrency, quantity, finalCurrency)
}

func convert(originalCurrency string, quantity float64, finalCurrency string) float64 {
	return 0
}

func userInput() (string, float64, string) {
	var originalCurrency string
	var finalCurrency string
	var quantity float64
	for {
		fmt.Print("Введите исходную валюту - EUR/USD/RUB: ")
		fmt.Scan(&originalCurrency)
		originalCurrencyLower := strings.ToLower(originalCurrency)
		if originalCurrencyLower != "eur" && originalCurrencyLower != "usd" && originalCurrencyLower != "rub" {
			fmt.Println("Некорректный ввод.")
			continue
		} else {
			originalCurrency = originalCurrencyLower
			break
		}
	}
	for {
		fmt.Print("Введите количество: ")
		fmt.Scan(&quantity)
		quantityType := fmt.Sprintf("%T", quantity)
		if quantityType != "float64" && quantityType != "int" {
			fmt.Println("Некорректный ввод.")
			continue
		} else {
			break
		}
	}
	for {
		switch originalCurrency {
		case "eur":
			fmt.Print("Введите целеву валюту - USD/RUB: ")
		case "usd":
			fmt.Print("Введите целеву валюту - EUR/RUB: ")
		case "rub":
			fmt.Print("Введите целеву валюту - EUR/USD: ")
		}
		fmt.Scan(&finalCurrency)
		finalCurrencyLower := strings.ToLower(finalCurrency)
		if finalCurrencyLower != "eur" && finalCurrencyLower != "usd" && finalCurrencyLower != "rub" {
			fmt.Println("Некорректный ввод.")
			continue
		} else if finalCurrencyLower == originalCurrency {
			fmt.Println("Целевая валюта не может совпадать с исходной.")
			continue

		} else {
			finalCurrency = finalCurrencyLower
			break
		}
	}
	return originalCurrency, quantity, finalCurrency
}
