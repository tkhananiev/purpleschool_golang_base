package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operations := userInputOperations()
	nums := userInputNums()
	fmt.Printf("Результат операции: %.1f", calculate(operations, nums))
}

func userInputOperations() string {
	var userInput string
	for {
		fmt.Print("Введите арифметическую операцию avg/sum/med или exit для выхода из меню: ")
		fmt.Scan(&userInput)
		userInput = strings.ToLower(userInput)
		if userInput != "avg" && userInput != "sum" && userInput != "med" && userInput != "exit" {
			fmt.Println("❌Ошибка: ожидается ввод только следующих значений - avg/sum/med или exit для выхода из меню. ")
			continue
		}
		if userInput == "exit" {
			break
		}
		return userInput
	}
	return ""
}

func userInputNums() []float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите числа, разделённые запятыми (например: 1,2,3): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, ",")
	var result []float64

	for _, part := range parts {
		part = strings.TrimSpace(part)
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			fmt.Println("❌ Ошибка: неверный формат числа →", part)
			return nil
		}
		result = append(result, num)
	}

	fmt.Println("✅ Ввод корректный!")
	return result
}

func calculate(userInputOperations string, userInputNums []float64) float64 {
	var result float64
	switch userInputOperations {
	case "avg":
		var sum float64
		for _, num := range userInputNums {
			sum += num
		}
		result = sum / float64(len(userInputNums))
	case "sum":
		for _, num := range userInputNums {
			result += num
		}
	case "med":
		n := len(userInputNums)
		if n == 0 {
			return 0 // или math.NaN()
		}
		sort.Float64s(userInputNums)

		if n%2 == 1 {
			return userInputNums[n/2] // нечётное — середина
		} else {
			return (userInputNums[n/2-1] + userInputNums[n/2]) / 2 // чётное — среднее двух
		}
	}
	return result
}
