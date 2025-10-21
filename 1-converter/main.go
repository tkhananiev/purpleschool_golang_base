package main

import (
	"fmt"
	"strings"
)

const (
	usdToEur = 0.85
	usdToRub = 80.5
	eurToUsd = 1 / usdToEur
	rubToUsd = 1 / usdToRub
	eurToRub = usdToRub / usdToEur
	rubToEur = 1 / eurToRub
)

func main() {

	originalCurrency, quantity, finalCurrency := userInput()
	finalCurrencyQuantity := convert(originalCurrency, quantity, finalCurrency)
	fmt.Printf("При обмене %.1f %s вы получите %.1f %s", quantity, originalCurrency, finalCurrencyQuantity, finalCurrency)
}

func convert(originalCurrency string, quantity float64, finalCurrency string) float64 {
	var finalCurrencyQuantity float64
	switch {
	case originalCurrency == "eur" && finalCurrency == "usd":
		finalCurrencyQuantity = eurToUsd * quantity
	case originalCurrency == "eur" && finalCurrency == "rub":
		finalCurrencyQuantity = eurToRub * quantity
	case originalCurrency == "usd" && finalCurrency == "eur":
		finalCurrencyQuantity = usdToEur * quantity
	case originalCurrency == "usd" && finalCurrency == "rub":
		finalCurrencyQuantity = usdToRub * quantity
	case originalCurrency == "rub" && finalCurrency == "usd":
		finalCurrencyQuantity = rubToUsd * quantity
	case originalCurrency == "rub" && finalCurrency == "eur":
		finalCurrencyQuantity = rubToEur * quantity
	}
	return finalCurrencyQuantity
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
