package main

import (
	"fmt"
)

func caesarCipherShiftRune(r rune) rune {
	return (((r - 'A') + 1) % 26) + 'A'
}

func main() {
	var input string
	fmt.Scanln(&input)
	var output string

	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRune(r)
		}

		output += string(r)
	}

	fmt.Println(output) // UIJT JT B UFTU
}
