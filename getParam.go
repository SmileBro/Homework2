package main

import (
	"fmt"
)

func getParam() (param string, err error) {
	_, err = fmt.Scanln(&param)
	return
}
