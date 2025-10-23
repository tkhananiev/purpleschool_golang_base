package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	//"strconv"
	"strings"
)

func main() {
	//fmt.Println(userInputOperations())
	//fmt.Println(userInputNums())
	userInputNums()
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

func userInputNums() []int {
	var inputString string
	var inputStrToInt []int
	pattern := `^\s*\d+(\s*,\s*\d+)*\s*$`
	re := regexp.MustCompile(pattern)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите числа, разделённые запятыми (например: 1,2,3): ")

	// Читаем строку ввода
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if re.MatchString(input) {
		fmt.Println("✅ Ввод корректный!")
		inputString += input
	} else {
		fmt.Println("❌ Ошибка: ожидается формат чисел, разделённых запятыми (например: 1,2,3)")
	}
	parts := strings.Split(inputString, ",")
	for _, part := range parts {
		itemInt, _ := strconv.Atoi(part)
		inputStrToInt = append(inputStrToInt, itemInt)

	}
	return inputStrToInt
}
