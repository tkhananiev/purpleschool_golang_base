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

func userInput(originalCurrency, intermediateCurrency, finalCurrency string) (string, string, string) {
	return originalCurrency, intermediateCurrency, finalCurrency
}

func convert(quantity float64, originalCurrency string, finalCurrency string) float64 {
	return 0
}
