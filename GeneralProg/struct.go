package main

import "fmt"

type moboArch struct {
	camera float32
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
}

func (i moboArch) call(number string) {
	fmt.Printf("Calling... [%v]\n", number)

}

func main() {
	iPhone := moboArch{}
	iPhone.camera = 12.4
	iPhone.color = "Black"
	iPhone.call("9162227847")

	fmt.Println("Features of iPhone : ", iPhone)

}
