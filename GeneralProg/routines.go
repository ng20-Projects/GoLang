package main

import (
	"fmt"
	"sync"
)

func main() {
	//Variable Init
	var wait sync.WaitGroup
	var counter int

	//Input Counter
	fmt.Print("Enter the counter: ")
	fmt.Scanln(&counter)

	//Wait Prog
	wait.Add(counter)
	defer wait.Wait()

	//Routine Creation
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
}

func hello(wait *sync.WaitGroup, counter int) {
	fmt.Println(counter)
	wait.Done()
}
