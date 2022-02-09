package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func do() {
	var sum int
	defer wait.Done()
	for j := 0; j < 10000; j++ {
		for i := 0; i < 9000000; i++ {
			sum += i
		}
		sum += j
	}
}

func main() {
	wait.Add(2)

	go do()
	go do()

	wait.Wait()
	fmt.Println("Done")
}
