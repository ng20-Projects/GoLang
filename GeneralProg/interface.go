package main

import "fmt"

/*
Interface is an abstract datatype
Consists of only method signature; We can even have a empty interface too

type interfaceName interface{
	method-1
	method-2
	.......
}
*/

type SalaryCalculator interface {
	calculateSalary() int
}

type PrintSalary interface {
	Print()
}

type Salary interface {
	SalaryCalculator
	PrintSalary
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

func (pj PermanentJob) Print() {
	fmt.Println(pj.basicPay)
}

func (pj PermanentJob) calculateSalary() int {
	return pj.basicPay
}

func (pj PermanentJob) getJobName() string {
	return "Permanent Job"
}

func (i Interger) print() {
	fmt.Println(i)
}
func (i Interger) calculateSalary() int {
	return int(i)
}

func totalIncome(salaryCalc []SalaryCalculator) {
	total := 0
	for _, v := range salaryCalc {
		total += v.calculateSalary()
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

	j := Interger(1009)
	j.print()

	//Code-2
	pj := PermanentJob{
		basicPay: 10,
	}
	cj := PermanentJob{
		basicPay: 20,
	}
	fj := PermanentJob{
		basicPay: 30,
	}

	salaryCalc := []SalaryCalculator{pj, cj, fj}
	totalIncome(salaryCalc)
}
