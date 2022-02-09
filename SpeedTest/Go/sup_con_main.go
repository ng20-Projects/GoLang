package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func do() {
	var sum int
	var internalWait sync.WaitGroup
	defer wait.Done()
	defer internalWait.Wait()
	go func() {
		for j := 0; j < 10000; j++ {
			for i := 0; i < 9000000; i++ {
				sum += i
			}
			sum += j
		}
		internalWait.Done()
	}()
}

func main() {
	wait.Add(2)

	go do()
	go do()

	wait.Wait()
	fmt.Println("Done")
}
