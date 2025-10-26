package main

import (
	"fmt"
	"strings"
)

func main() {
	rates := map[string]float64{
		"eur_usd": 1 / 0.85, // 1 EUR = 1.176 USD
		"eur_rub": 80.5 / 0.85,
		"usd_eur": 0.85,
		"usd_rub": 80.5,
		"rub_usd": 1 / 80.5,
		"rub_eur": 1 / (80.5 / 0.85),
	}

	originalCurrency, quantity, finalCurrency := userInput()
	finalCurrencyQuantity := convert(originalCurrency, quantity, finalCurrency, rates)

	fmt.Printf("При обмене %.1f %s вы получите %.1f %s\n",
		quantity, strings.ToUpper(originalCurrency),
		finalCurrencyQuantity, strings.ToUpper(finalCurrency))
}

func convert(originalCurrency string, quantity float64, finalCurrency string, rates map[string]float64) float64 {
	key := originalCurrency + "_" + finalCurrency
	rate, exists := rates[key]
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
