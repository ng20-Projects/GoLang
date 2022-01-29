/*
Create an implementation of the rotational cipher, also sometimes called the Caesar cipher.

The Caesar cipher is a simple shift cipher that relies on transposing all the letters in the alphabet using an integer key between 0 and 26. Using a key of 0 or 26 will always yield the same output due to modular arithmetic. The letter is shifted for as many values as the value of the key.

The general notation for rotational ciphers is ROT + <key>. The most commonly used rotational cipher is ROT13.

A ROT13 on the Latin alphabet would be as follows:

Plain:  abcdefghijklmnopqrstuvwxyz //shiftkey =27
Cipher: nopqrstuvwxyzabcdefghijklm  //27-26=1

It is stronger than the Atbash cipher because it has 27 possible keys, and 25 usable keys.

Ciphertext is written out in the same formatting as the input including spaces and punctuation.
Examples

    ROT5 omg gives trl
    ROT0 c gives c
    ROT26 Cool gives Cool
    ROT13 The quick brown fox jumps over the lazy dog. gives Gur dhvpx oebja sbk whzcf bire gur ynml qbt.
    ROT13 Gur dhvpx oebja sbk whzcf bire gur ynml qbt. gives The quick brown fox jumps over the lazy dog.

*/

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
