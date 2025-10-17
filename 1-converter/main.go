package main

import (
	"fmt"
)

const (
	usdToEur = 0.85
	usdToRub = 80.5
)

func main() {
	fmt.Println(eurToRub(50))

}

func eurToRub(usd float64) float64 {
	return usdToRub / usdToEur
}
