package main

import "fmt"

func checkIsogram(str string) bool {
	count := 0
	for i := range str {
		for j := range str {
			if i != j {
				if str[i] == str[j] {
					count++
				}
			} else {
				continue
			}
		}
	}
	if count == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var isogram string
	fmt.Scanln(&isogram)
	res := checkIsogram(isogram)
	if res {
		fmt.Printf("Yes, the string %v is Isogram", isogram)
		fmt.Println("")
	} else {
		fmt.Printf("No, the string %v is Isogram", isogram)
		fmt.Println("")
	}
}
