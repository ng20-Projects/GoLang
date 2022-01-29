package main

import "fmt"

func divide(num, den int) int {
	return num / den
}

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("[*]Recovered")
	}
}

func main() {
	var num, den int
	fmt.Print("Numerator: ")
	fmt.Scanln(&num)
	fmt.Print("Denominator: ")
	fmt.Scanln(&den)

	/*
		// Approach-1
		defer func() {
			fmt.Println(recover())
		}()
	*/

	// Approach-2
	defer recoverFunc()
	if den == 0 {
		panic("Denominator is zero")
	}
	fmt.Println(divide(num, den))
}
