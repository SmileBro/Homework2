package main

import (
	"fmt"
	"strings"
)

func Separate(input string) (result string, err error) {

	if strings.Contains(input, "(") {
		b1 := strings.Count(input, "(")
		b2 := strings.Count(input, ")")
		if b1 != b2 || strings.Contains(input, "-+") {
			fmt.Println("Неправильный ввод")
			return
		}

		for i := 0; i < b1; i++ {
			input, err = Calculate(input)
			input = strings.ReplaceAll(input, "--", "+")
		}
	}
	result, err = Solve(input)
	return
}
func Calculate(input string) (result string, err error) {
	var output string
	i := strings.Index(input, "(")
	c := i
	var j int

	for j = c; j < len(input); j++ {
		if input[j] == 40 {
			i = j
		}
		if input[j] == 41 {

			break
		}
	}

	output, err = Solve(input[i+1 : j])
	result = input[0:i] + output + input[j+1:]

	return
}
