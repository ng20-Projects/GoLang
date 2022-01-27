package main

import (
	"errors"
	"fmt"
)

func divide(num, den int) (int, error) {
	if den != 0 {
		return num / den, nil
	} else {
		return 0, errors.New("Denominator cannot be zero")
	}
}

func main() {
	var num, den int
	fmt.Print("Numerator: ")
	fmt.Scanln(&num)
	fmt.Print("Denominator: ")
	fmt.Scanln(&den)
	res, err := divide(num, den)
	if err == nil {
		fmt.Println("Result is", res)
	} else {
		fmt.Println("[!]Warning:", err)
	}
}
