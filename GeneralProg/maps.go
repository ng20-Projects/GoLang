package main

import "fmt"

func main() {
	studentClass := map[int]string{ //{Key-> RollNumber : Value-> Name}
		1: "Rahul",
		2: "Nikhilesh",
		3: "Prajjawal",
		4: "Saket",
		5: "Hardik",
	}

	studentClass[6] = "Vigneshwar"
	studentClass[7] = "Eesha"
	studentClass[8] = "Suraj"
	fmt.Println(studentClass)
	fmt.Println("")

}
