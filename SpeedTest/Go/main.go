package main

import (
	"fmt"
)

func main() {
	// time.Time()
	var sum int
	for j := 0; j < 10000; j++ {
		for i := 0; i < 9000000; i++ {
			sum += i
		}
		sum += j
	}
	for j := 0; j < 10000; j++ {
		for i := 0; i < 9000000; i++ {
			sum += i
		}
		sum += j
	}
	fmt.Println("Done")
	// fmt.Println("Time taken: ")
}
