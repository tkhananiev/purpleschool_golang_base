package main

import (
	"fmt"
	"strings"
)

var rates = map[string]map[string]float64{
	"eur": {
		"eur_usd": 1 / 0.85,
		"eur_rub": 80.5 / 0.85,
	},
	"usd": {
		"usd_eur": 0.85,
		"usd_rub": 80.5,
	},
	"rub": {
		"rub_usd": 1 / 80.5,
		"rub_eur": 1 / (80.5 / 0.85),
	},
}
var ptrRates = &rates

func main() {

	originalCurrency, quantity, finalCurrency := userInput()
	finalCurrencyQuantity := convert(originalCurrency, quantity, finalCurrency, &rates)

	fmt.Printf("При обмене %.1f %s вы получите %.1f %s\n",
		quantity, strings.ToUpper(originalCurrency),
		finalCurrencyQuantity, strings.ToUpper(finalCurrency))
}

func convert(originalCurrency string, quantity float64, finalCurrency string, rates *map[string]map[string]float64) float64 {
	key := originalCurrency + "_" + finalCurrency
	originalCurrencyRate, exist := (*rates)[originalCurrency]
	if !exist {
		fmt.Println("Неизвестная исходная валюта")
		return 0
	}
	rate, exists := originalCurrencyRate[key]
	if !exists {
		fmt.Println("Неизвестная комбинация валют")
		return 0
	}
	return quantity * rate
}

func userInput() (string, float64, string) {
	var originalCurrency, finalCurrency string
	var quantity float64

	for {
		fmt.Print("Введите исходную валюту - EUR/USD/RUB: ")
		fmt.Scan(&originalCurrency)
		originalCurrency = strings.ToLower(originalCurrency)
		if originalCurrency == "eur" || originalCurrency == "usd" || originalCurrency == "rub" {
			break
		}
		fmt.Println("Некорректный ввод.")
	}

	fmt.Print("Введите количество: ")
	fmt.Scan(&quantity)

	for {
		fmt.Print("Введите целевую валюту - EUR/USD/RUB: ")
		fmt.Scan(&finalCurrency)
		finalCurrency = strings.ToLower(finalCurrency)
		if finalCurrency == originalCurrency {
			fmt.Println("Целевая валюта не может совпадать с исходной.")
			continue
		}
		if finalCurrency == "eur" || finalCurrency == "usd" || finalCurrency == "rub" {
			break
		}
		fmt.Println("Некорректный ввод.")
	}

	return originalCurrency, quantity, finalCurrency
}
