package main

import (
	"fmt"
)

const (
	usdToEur = 0.85
	usdToRub = 80.5
	eurToRub = usdToRub / usdToEur
)

func main() {
	fmt.Println(eurToRub)
}

func userInput(originalСurrency, intermediateCurrency, finalCurrency string) (string, string, string) {
	return originalСurrency, intermediateCurrency, finalCurrency
}

func convert(quantity int, originalСurrency string, finalCurrency string) float64 {
	return 0
}
