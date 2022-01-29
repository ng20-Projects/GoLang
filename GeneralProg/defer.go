package main

import "fmt"

func f() {
	fmt.Println("defer f()")
}

func x() {
	fmt.Println("defer x()")
}

func main() {
	fmt.Println(1)
	fmt.Println(2)
	defer f()
	fmt.Println(3)
	fmt.Println(4)
	defer x()
	fmt.Println(5)
}
