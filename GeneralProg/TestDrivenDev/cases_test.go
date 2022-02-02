package main

var testCases = []struct {
	description string
	input       int
	expected    int
}{
	{
		"Factorial of -ve",
		-1,
		0,
	},
	{
		"Factorial of 0",
		0,
		1,
	},
	{
		"Factorial of 5",
		5,
		120,
	},
	{
		"Factorial of 10",
		10,
		3628800,
	},
	{
		"Factorial of 15",
		15,
		1307674368000,
	},
}
