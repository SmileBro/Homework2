package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (line string, err error) {
	var result float64
	f1 := func(c rune) bool { return c == 43 }

	//f2 := func(c rune) bool { return !unicode.IsNumber(c) }
	s1 := strings.FieldsFunc(input, f1)
	//s2 := strings.FieldsFunc(input, f2)
	//fmt.Println(s1)
	//fmt.Println(s2)

	for _, v := range s1 {
		out, err1 := ExpMinus(v)
		if err1 != nil {
			err = err1
			return
		}
		result = result + out
	}
	line = fmt.Sprintf("%v", result)
	return

}
func ExpMinus(input string) (result float64, err error) {

	f1 := func(c rune) bool { return c == 45 }
	//f2 := func(c rune) bool { return !unicode.IsNumber(c) }
	s1 := strings.FieldsFunc(input, f1)
	//s2 := strings.FieldsFunc(input, f2)
	//fmt.Println(s1)
	//fmt.Println(s2)
	var min bool
	if input[0] == 45 {
		min = true
	}
	flag := true
	for _, v := range s1 {
		if flag {
			flag = false
			temp, err1 := ExpMult(v)
			if err1 != nil {
				err = err1
				return
			}
			if !min {
				result = temp
			} else {
				result = -temp
			}

		} else {
			temp, err1 := ExpMult(v)
			if err1 != nil {
				err = err1
				return
			}
			result = result - temp
		}
	}
	return
}
func ExpMult(input string) (result float64, err error) {
	f1 := func(c rune) bool { return c == 42 }
	//f2 := func(c rune) bool { return !unicode.IsNumber(c) }
	s1 := strings.FieldsFunc(input, f1)
	//s2 := strings.FieldsFunc(input, f2)
	//fmt.Println(s1)
	//fmt.Println(s2)
	result = 1.0
	for _, v := range s1 {
		out, err1 := ExpDiv(v)
		if err1 != nil {
			err = err1
			return
		}
		result = result * out
	}
	return
}
func ExpDiv(input string) (result float64, err error) {
	f1 := func(c rune) bool { return c == 47 }
	//f2 := func(c rune) bool { return !unicode.IsNumber(c) }
	s1 := strings.FieldsFunc(input, f1)
	//s2 := strings.FieldsFunc(input, f2)
	//fmt.Println(s1)
	//fmt.Println(s2)

	flag := true
	for _, v := range s1 {
		if flag {
			flag = false
			temp, err1 := strconv.ParseFloat(s1[0], 64)
			if err1 != nil {
				err = err1
				return
			}
			result = temp
		} else {
			temp, err1 := strconv.ParseFloat(v, 64)
			if err1 != nil {
				err = err1
				return
			}
			result = result / temp
		}
	}

	return
}
