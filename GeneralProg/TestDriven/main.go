package main

func Factorial(input int) int {
	var fact int = input
	if input < 0 {
		if input == 0 {
			return 1
		} else {
			return 0
		}
	} else if input == 1 {
		return input
	} else {
		fact = fact * Factorial(input-1)
		return fact
	}
}

func main() {
	// var input int
	// fmt.Scanln(&input)
	// fmt.Println(Factorial(input))
}
