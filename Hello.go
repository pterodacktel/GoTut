package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	} else {
		return a / b, nil
	}
}
func main() {
	var s, m int
	fmt.Scan(&s, &m)
	d, err := divide(s, m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
}
