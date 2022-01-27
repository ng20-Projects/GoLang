package main

import "fmt"

func RotForLowerCase(char rune, key int) rune {
	return (((char - 'a') + rune(key)) % 26) + 'a'
}

func RotForUpperCase(char rune, key int) rune {
	return (((char - 'A') + rune(key)) % 26) + 'A'
}

func main() {
	var plainText, cipherText string
	var key int
	fmt.Print("Enter the Plain Text: ")
	fmt.Scanln(&plainText)
	fmt.Print("Rot Key: ")
	fmt.Scanln(&key)
	for _, i := range plainText {
		if i >= 'a' && i <= 'z' {
			i = RotForLowerCase(i, key)
		} else if i >= 'A' && i <= 'Z' {
			i = RotForUpperCase(i, key)
		}
		cipherText += string(i)
	}
	fmt.Printf("Cipher Text of %s is %s\n", plainText, cipherText)
}
