package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type= %T, Value= %v \n", i, i)
	// var num int
	// var ok bool
	// 	num, ok = i.(int) //Varible 'ok' handle the runtime error
	// 	fmt.Println(num)
	// 	fmt.Println(ok)
	switch i.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Integer")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unsupported Data Type")
	}
}

func main() {
	describe("Hello")
	describe(20)
	describe(true)
	describe(31.71)
}
