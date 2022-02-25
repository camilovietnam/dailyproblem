package brackets

import "testing"

var tests = []struct{
	Name string
	Input string
	Expected bool
} {
	{
		Name: "empty string",
		Input: "",
		Expected: true,
	},
	{
		Name: "One parenthesis",
		Input: "(",
		Expected: false,
	},
	{
		Name: "One pair of round brackets",
		Input: "()",
		Expected: true,
	},
	{
		Name: "One round, one square, one curly",
		Input: "([]){}",
		Expected: true,
	},
	{
		Name: "Multiple rounds, one square",
		Input: "(()(())())[]()",
		Expected: true,
	},
	{
		Name: "Multiple rounds, one square",
		Input: "([])[]({})",
		Expected: true,
	},
	{
		Name: "Example number 1",
		Input: "([)]",
		Expected: false,
	},
	{
		Name: "Example number 2",
		Input: "((()",
		Expected: false,
	}, 
	{
		Name: "A very long string",
		Input: "([[{}]][{{([])}}]){{[]}}[]({})({})",
		Expected: true,
	}, 
	{
		Name: "Missing close square bracket",
		Input: "(([",
		Expected: false,
	}, 
	{
		Name: "Missing close round bracket",
		Input: "(([])",
		Expected: false,
	}, 
	{
		Name: "last character breaks it",
		Input: "{}[()][",
		Expected: false,
	}, 
}

func Test_BalancedTests(t *testing.T) {
	for _, test := range tests {
		t.Log("Running test...", test.Name)
		got := isBalanced(test.Input)

		if got != test.Expected {
			t.Fatalf("Failed test: %s", test.Name)
		}
	}
}