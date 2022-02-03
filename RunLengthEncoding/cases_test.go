package main

var testCases = []struct {
	input       string
	expected    string
	description string
}{
	{
		"55332",
		"error: digits not allowed",
		"Digits",
	},
	{
		"  hee he",
		"2 h2e he",
		"White-Space",
	},
	{
		"!!@@@#",
		"2!3@#",
		"Special Symbols",
	},
	{
		"\"\"\"\"",
		"4\"",
		"Special Symbols: Quotation",
	},
	{
		"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB",
		"12WB12W3B24WB",
		"Mixed[UpperCase]: repeating and non-repeating letters",
	},
	{
		"WWWBBBCCDDBB",
		"3W3B2C2D2B",
		"[UpperCase]: repeating letters",
	},
	{
		"aabcccdeeee",
		"2ab3cd4e",
		"Mixed[LowerCase]: repeating and non-repeating letters",
	},
	{
		"aabbcccddeee",
		"2a2b3c2d3e",
		"[LowerCase]: repeating letters",
	},
	{
		"AaabBBccDeeee",
		"A2ab2B2cD4e",
		"Mixed[BothCase]: repeating and non-repeating letters",
	},
	{
		"AAA bB  CccDE",
		"3A bB2 C2cDE",
		"Mixed[BothCase + WhiteSpaces]: repeating and non-repeating letters",
	},
	{
		"ABCDEFGHIJKLMNOP",
		"ABCDEFGHIJKLMNOP",
		"Non-repeating letters",
	},
	{
		"A",
		"A",
		"Single Character",
	},
	{
		"",
		"",
		"Empty",
	},
}
