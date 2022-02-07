package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		The Normal Output should be 1 => 1 | 2 =>2 .. and so on
		But this is not happening because of the Race Condition
	*/
	{
		counter := 0
		for i := 1; i <= 10; i++ {
			go func(i int) {
				fmt.Println(i, "=>", counter)
				counter++
				fmt.Println(i, "=>", counter)
			}(i)
		}
		time.Sleep(time.Second * 3)
	}

	//Seperator
	fmt.Println("###################")
	fmt.Println("###################")
	fmt.Println("###################")
	// END
	//Solution for the above issue => use of WaitGroup
	{
		var mutex sync.Mutex
		counter := 0
		for i := 1; i <= 10; i++ {
			go func(i int) {
				fmt.Println(i, "=>", counter)
				mutex.Lock()
				counter++
				fmt.Println(i, "=>", counter)
				mutex.Unlock()
			}(i)
		}
		// fmt.Println(counter)
		time.Sleep(time.Second * 3)
	}
}
