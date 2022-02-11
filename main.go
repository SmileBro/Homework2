package main

import "fmt"

func main() {
	fmt.Println("Введите строку для рассчета:")
	input, err := getParam()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Неправильный ввод")
		return
	} else {
		fmt.Print(input)
		fmt.Print("=")
	}
	result, err := Solve(input)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Print(result)
	}
}
