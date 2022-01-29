package main

import "fmt"

// type Integer int

type moboArch struct {
	camera float32
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
	proc   newArch
}

type newArch struct {
	coreProcessor  int
	multiProcessor int
}

func (i moboArch) call(number string) {
	fmt.Printf("Calling... [%v]\n", number)
	fmt.Println(i.camera)
}

func main() {
	iPhone := moboArch{}
	iPhone.camera = 12.4
	iPhone.color = "Black"
	iPhone.call("9162227847")
	iPhone.proc.coreProcessor = 10
	iPhone.proc.multiProcessor = 20
	fmt.Println("Features of iPhone : ", iPhone)

	Realme := moboArch{}
	Realme.camera = 15.2
	Realme.color = "Blue"
	Realme.call("9955300897")
	Realme.proc.coreProcessor = 30
	Realme.proc.multiProcessor = 10
	fmt.Println("Features of Realme : ", Realme)
}
