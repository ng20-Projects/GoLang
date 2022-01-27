package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		continue
	}
	elapsed := time.Since(start) // This will check how long has passed since time.Now() was executed
	fmt.Println(elapsed)

}
