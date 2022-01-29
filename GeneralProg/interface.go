package main

import "fmt"

/*
Interface is an abstract datatype
Consists of only method signature; We can even have a empty interface too

type interfaceName interface{
	method-1
	method-2
	.......
	a = input("")
	type a -> "any"
}
*/

type SalaryCalculator interface {
	calculateSalary() int
	getJobName() string
}

type PrintSalary interface {
	Print()
}

type Salary interface {
	SalaryCalculator //-> calculateSalary() + getJobName()
	PrintSalary      //-> Print()
}

// Struct Datatypes -----------
type PermanentJob struct {
	basicPay int
}

type ContractJob struct {
	basicPay int
}

type FreelanceJob struct {
	basicPay int
}

type Interger int

// End -----------

func (pj PermanentJob) calculateSalary() int {
	return pj.basicPay
}

func (pj PermanentJob) getJobName() string {
	return "Permanent Job"
}

func (pj PermanentJob) Print() {
	fmt.Println(pj.basicPay)
}

func (cj ContractJob) calculateSalary() int {
	return cj.basicPay
}

func (cj ContractJob) getJobName() string {
	return "Contract Job"
}

func (cj ContractJob) Print() {
	fmt.Println(cj.basicPay)
}

func (fj FreelanceJob) calculateSalary() int {
	return fj.basicPay
}

func (fj FreelanceJob) getJobName() string {
	return "Freelance Job"
}

func (fj FreelanceJob) Print() {
	fmt.Println(fj.basicPay)
}

func (i Interger) print() { //func(i int) print(){}
	fmt.Println(i)
}
func (i Interger) calculateSalary() int {
	return int(i)
}

func totalIncome(salaryCalc []SalaryCalculator) {
	total := 0
	for _, v := range salaryCalc {
		total += v.calculateSalary()
		fmt.Println(v.getJobName())
	}
	fmt.Println(total)
}

func main() {
	// var salaryCalc SalaryCalculator
	// This variable can hold the data of multiple datatypes like type int(), string() ...
	// But there is some condition that a Interface have to follow
	// Condition is that the datatype should implement all the methods of the interface

	// Code-1
	// pj := PermanentJob{
	// 	basicPay: 20,
	// }
	// salaryCalc := pj
	// total := salaryCalc.calculateSalary()
	// fmt.Println(total)

	// j := Interger(1009) // type Interger int -> Var(value)
	// j.print()

	//Code-2
	pj := PermanentJob{
		basicPay: 10,
	}
	cj := ContractJob{
		basicPay: 20,
	}
	fj := FreelanceJob{
		basicPay: 30,
	}

	// salaryCalc := []SalaryCalculator{pj, cj, fj}
	// totalIncome(salaryCalc)

	salary := []Salary{pj, cj, fj}
	fmt.Println(salary)
	salary[0].Print()
	salary[1].Print()
	salary[2].Print()

	fmt.Println(salary)
	for i := range salary {
		fmt.Println(i)
		// fmt.Println (salary[i].Print())
	}
	fmt.Sprintln("Nikhil")
}
