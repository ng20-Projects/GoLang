package main

import (
	"fmt"
)

type Today struct {
	date  int
	month string
	year  int
	day   string
}

func displayDetailsOfToday(td Today) (int, string, int, string) {
	fmt.Print("Date: ")
	fmt.Scanln(&td.date)
	fmt.Print("Month: ")
	fmt.Scanln(&td.month)
	fmt.Print("Year: ")
	fmt.Scanln(&td.year)
	fmt.Print("Day: ")
	fmt.Scanln(&td.day)
	// fmt.Printf("%d %s %d %s", td.date, td.month, td.year, td.day)
	return int(td.date), string(td.month), int(td.year), string(td.day)
}

func return2values() (int, string) {
	return 20, "Nikhilesh"
}

func funcForArr(arr [3]int) {
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(arr[i])
	// }
	for _, v := range arr {
		fmt.Println(v)
	}
	// fmt.Println(arr[1:3])
	// fmt.Println(arr[1:2])
	// fmt.Println(arr[1:1])
	// fmt.Println(arr[0:1])
}

func main() {
	// v1, v2 := return2values()
	// fmt.Printf("V1 is %d and V2 is %s\n", v1, v2)

	// const c1 = 20
	// const c2 string = "Nikhil"
	// v1 := 10
	// v2 := "esh"
	// fmt.Println(c1 + v1)
	// fmt.Println(c2 + v2)

	// var arr [10][10]int
	// val := 0
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		arr[i][j] = val
	// 		val++
	// 	}
	// }
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Printf("Value of arr[%d][%d]: ", i, j)
	// 		fmt.Println(arr[i][j])
	// 	}
	// }

	// arr2 := [3]int{10, 20, 30}
	// funcForArr(arr2)
	// fmt.Printf("Length: %d | Capacity %d\n", len(arr2), cap(arr2))

	// myslice := []int{11, 23, 38}
	// fmt.Println(myslice)
	// fmt.Printf("Length: %d | Capacity %d\n", len(myslice), cap(myslice))

	// myslice2 := [][]int{{1, 10}, {2, 20}, {3, 30}}
	// fmt.Println(myslice2)
	// fmt.Printf("Length: %d | Capacity %d\n", len(myslice), cap(myslice))

	// var val = arr2[1:3]
	// fmt.Printf("Type -> %T and Value -> %v]\n", val, val)

	// var i int = 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 	var count int = 0
	// GO1:
	// 	if count < 5 {
	// 		fmt.Println("Achieved !!")
	// 	} else {
	// 		goto GO2
	// 	}
	// 	count++
	// 	goto GO1
	// GO2:
	// 	fmt.Println("End")

	// var td Today                                        //Line 5
	// date, month, year, day := displayDetailsOfToday(td) //Line 12
	// fmt.Printf("%d %s %d %s \n", date, month, year, day)

	// var u rune
	// u = 'a'
	// fmt.Println(u)
	// fmt.Println(reflect.TypeOf(u))

	{
		var arr [10]int
		arr[0] = 1
		arr[1] = 2
		arr[2] = 3
		for v := range arr {
			arr[v] = v + 1
		}
		// arr[0] = 20
		func(arrF [10]int) {
			fmt.Printf("Value of arr: %v, Type of arr: %T\n", arrF, arrF)
		}(arr)

		// var slice []int
		// slice = append(slice, 1)
		// slice = append(slice, 2)
		// for i := 3; i < 11; i++ {
		// 	slice = append(slice, i)
		// }
		// func(sliceF []int) {
		// 	fmt.Printf("Value of arr: %v, Type of arr: %T\n", sliceF, sliceF)
		// }(slice)

		// // var digits map[int]int //map is declared nil, hence nothing can be added to it later on
		// var digits = make(map[int]int)
		// digits[0] = 1
		// digits[1] = 2
		// for i := 0; i < 10; i++ {
		// 	digits[i] = i + 1
		// }
		// func(digitsF map[int]int) {
		// 	fmt.Println(digitsF)
		// }(digits)
	}

	{
		// var input interface{}

		// input = 10
		// fmt.Println(input)
		// input = 10.02
		// fmt.Println(input)
		// input = true
		// fmt.Println(input)
		// input = "Data"
		// fmt.Println(input)

		// fmt.Scanln(&input)
		// fmt.Println(input)
	}

}
