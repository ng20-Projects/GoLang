package main

import "testing"

func TestEncode(t *testing.T) {
	for _, cases := range testCases {
		if cases.expected == RunLengthEncode(cases.input) {
			t.Logf("Pass for %s", cases.input)
		} else {
			// t.Fatalf("Fail: %s, wanted: %s, got: %s", cases.description, cases.input, cases.expected)
			t.Fatalf("Fail: %s, wanted: %s, got: %s", cases.description, cases.expected, cases.input)
		}
	}
}

func TestDigit(t *testing.T) {
	if isDigit("5") {
		t.Logf("Successfully Passed")
	} else {
		t.Fatalf("Failed: Digit, wanted: true, get false")
	}
}
