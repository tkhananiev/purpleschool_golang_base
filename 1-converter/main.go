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
