package main

import "fmt"

var strand1, strand2 string

func hammingDistance(s1, s2 string) int {
	count := 0
	for i := 0; i < len(s1); i++ {
		if !(s1[i] == s2[i]) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Scanln(&strand1) //Example - GAGCCTACTAACGGGAT
	fmt.Scanln(&strand2) //Example - CATCGTAATGACGGCCT
	if len(strand1) == len(strand2) {
		fmt.Println("Hamming Distance is ", hammingDistance(strand1, strand2))
	} else {
		fmt.Println("[!]Warning: 'Length of the DNA strands given are not same'")
	}
}
