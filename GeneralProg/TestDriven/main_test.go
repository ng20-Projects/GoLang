package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	for _, cases := range testCases {
		if Factorial(cases.input) == cases.expected {
			t.Logf("Pass : factorial(%d)", cases.input)
		} else {
			t.Fatalf("Fail : %s, wanted %d, got %d", cases.description, cases.expected, Factorial(cases.input))
		}
	}
}

func TestZeroFactorial(t *testing.T) {
	if Factorial(0) == 1 {
		t.Logf("Pass")
		return
	} else {
		t.Fatalf("Fail: %s, wanted %d, got %d", "Factorial of Zero", 0, Factorial(0))
	}
}

func TestFiveFactorial(t *testing.T) {
	if Factorial(5) == 120 {
		t.Logf("Pass")
		return
	} else {
		t.Fatalf("Fail: %s, wanted %d, got %d", "Factorial of Five", 120, Factorial(0))
	}
}
func TestTenFactorial(t *testing.T) {
	if Factorial(10) == 3628800 {
		t.Logf("Pass")
		return
	} else {
		t.Fatalf("Fail: %s, wanted %d, got %d", "Factorial of Ten", 3628800, Factorial(0))
	}
}
func TestFifteenFactorial(t *testing.T) {
	if Factorial(15) == 1307674368000 {
		t.Logf("Pass")
		return
	} else {
		t.Fatalf("Fail: %s, wanted %d, got %d", "Factorial of Fifteen", 1307674368000, Factorial(0))
	}
}
