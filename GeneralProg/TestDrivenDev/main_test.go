package main

import "testing"

func TestFactorial(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.input); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
