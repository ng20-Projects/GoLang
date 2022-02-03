package main

import (
	"fmt"
	"strconv"
)

var digit []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func isDigit(input string) bool {
	for _, val := range digit {
		if val == input {
			return true
		}
	}
	return false
}

func RunLengthEncode(input string) string {
	var encodedString string = ""
	var numberOfOccurence int = 0

	//For Empty String
	if input == "" {
		return ""
	}

	charToCompare := input[0] //To start we assign the first character
	for i := 0; i < len(input); i++ {
		if input[i] == charToCompare {
			numberOfOccurence += 1
		} else if isDigit(string(input[i])) { //For Digits Checking in the input string
			return "error: digits not allowed"
		} else {
			if numberOfOccurence == 1 {
				encodedString += string(charToCompare)
			} else {
				encodedString += strconv.Itoa(numberOfOccurence) + string(charToCompare)
			}
			charToCompare = input[i]
			numberOfOccurence = 1
		}
	}
	if numberOfOccurence == 1 {
		encodedString += string(charToCompare)
	} else {
		encodedString += strconv.Itoa(numberOfOccurence) + string(charToCompare)
	}

	return encodedString
}

func main() {
	var input string
	fmt.Print("Enter the String: ")
	fmt.Scanln(&input)
	en := RunLengthEncode(input)
	fmt.Printf("RLE -> [%s]\n", en)
}
