package main

import (
	"fmt"
	"sort"
)

func main() {
	const N int = 6
	var arr [N]int
	var k int
	for i := 0; i < 6; i++ {
		fmt.Printf("arr[%d]: ", i)
		fmt.Scanln(&arr[i])
	}
	fmt.Print("k: ")
	fmt.Scanln(&k)
	arrS := arr[:]
	sort.Ints(arrS)
	fmt.Println(arrS)
	fmt.Println(arr[len(arr)-k])

}
