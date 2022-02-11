package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (result float64, err error) {

	//f1 := func(c rune) bool { return !(c == 45) && !(c == 43) }
	f1 := func(c rune) bool { return !(c == 43) }
	//f2 := func(c rune) bool { return (c == 45) || (c == 43) }
	f2 := func(c rune) bool { return c == 43 }
	s1 := strings.FieldsFunc(input, f1)
	s2 := strings.FieldsFunc(input, f2)
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := make([]string, len(s2))
	for i, v := range s2 {
		temp, err1 := Exp(v)
		if err1 != nil {
			err = err1
			return
		}
		s3[i] = temp

	}
	fmt.Println(s3)
	flag := true
	if len(s1) > 0 {
		for i, v := range s1 {
			if flag {
				flag = false
				switch v {
				case "-":
					{
						t1, err1 := strconv.ParseFloat(s3[i], 64)
						if err1 != nil {
							err = err1
							return
						}
						t2, err1 := strconv.ParseFloat(s3[i+1], 64)
						if err1 != nil {
							err = err1
							return
						}
						result = t1 - t2
					}
				case "+":
					{
						t1, err1 := strconv.ParseFloat(s3[i], 64)
						if err1 != nil {
							err = err1
							return
						}
						t2, err1 := strconv.ParseFloat(s3[i+1], 64)
						if err1 != nil {
							err = err1
							return
						}
						result = t1 + t2
					}

				}
			} else {
				if i != len(s1) {
					switch v {
					case "-":
						{
							t1, err1 := strconv.ParseFloat(s3[i+1], 64)
							if err1 != nil {
								err = err1
								return
							}

							result = result - t1
						}
					case "+":
						{
							t1, err1 := strconv.ParseFloat(s3[i+1], 64)
							if err1 != nil {
								err = err1
								return
							}

							result = result + t1
						}

					}
				}
			}
		}
	} else {
		t2, err1 := strconv.ParseFloat(s3[0], 64)
		if err1 != nil {
			err = err1
			return
		}
		result = t2
	}
	return

}
func Exp(s string) (line string, err error) {
	var result float64

	if strings.Contains(s, "-") {
		buf := strings.Split(s, "-")
		fmt.Println(buf)
		if buf[0] == "" {
			result1, err1 := strconv.ParseFloat(buf[1], 64)
			if err1 != nil {
				err = err1
				return
			}
			result = result1
		} else {
			result1, err1 := strconv.ParseFloat(buf[0], 64)
			if err1 != nil {
				err = err1
				return
			}
			result = result1
		}

		for i, v := range buf {
			if i != 0 {
				temp, err1 := Expresion(v)
				if err1 != nil {
					err = err1
					return
				}
				result = result - temp
			}
		}

	} else {
		if strings.Contains(s, "*") {
			result = 1.0
			temp := strings.Split(s, "*")
			for _, v := range temp {
				buf, err1 := Exp(v)
				if err1 != nil {
					err = err1
					return
				}
				t1, err1 := strconv.ParseFloat(buf, 64)
				if err1 != nil {
					err = err1
					return
				}
				result = result * t1
			}
		} else {
			if strings.Contains(s, "/") {
				result = 1.0
				flag := true
				temp := strings.Split(s, "/")
				for _, v := range temp {
					buf, err1 := Exp(v)
					if err1 != nil {
						err = err1
						return
					}
					t1, err1 := strconv.ParseFloat(buf, 64)
					if err1 != nil {
						err = err1
						return
					}
					if flag {
						flag = false
						result = t1
					} else {
						result = result / t1
					}
				}
			} else {
				_, err1 := strconv.ParseFloat(s, 64)
				if err1 != nil {
					err = err1
					fmt.Println("Неправильное выражение")
					return
				}
				line = s
				return
			}
		}
	}

	line = fmt.Sprintf("%v", result)
	return
}
func Expresion(s string) (result float64, err error) {
	if strings.Contains(s, "*") {
		result = 1.0
		temp := strings.Split(s, "*")
		for _, v := range temp {
			buf, err1 := Exp(v)
			if err1 != nil {
				err = err1
				return
			}
			t1, err1 := strconv.ParseFloat(buf, 64)
			if err1 != nil {
				err = err1
				return
			}
			result = result * t1
		}
	} else {
		if strings.Contains(s, "/") {
			result = 1.0
			flag := true
			temp := strings.Split(s, "/")
			for _, v := range temp {
				buf, err1 := Exp(v)
				if err1 != nil {
					err = err1
					return
				}
				t1, err1 := strconv.ParseFloat(buf, 64)
				if err1 != nil {
					err = err1
					return
				}
				if flag {
					flag = false
					result = t1
				} else {
					result = result / t1
				}
			}
		} else {
			result1, err1 := strconv.ParseFloat(s, 64)
			if err1 != nil {
				err = err1
				fmt.Println("Неправильное выражение")
				return
			}
			result = result1
			return

		}
	}
	return
}
