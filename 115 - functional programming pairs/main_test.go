package main

import (
	"testing"
)

type pairType struct {
	a int
	b int
}

type testType struct {
	input pairType
	expected int
	name string
}

func Test_Basic(t *testing.T) {
	tests := []testType{
		{input: pairType{3,4}, expected: 7, name: "Test_Cons_1",},
		{input: pairType{1,2}, expected: 3, name: "Test_Cons_2",},
		{input: pairType{1,0}, expected: 1, name: "Test_Cons_3",},
		{input: pairType{100,222}, expected: 322, name: "Test_Cons_4",},
	}

	for _, test := range tests {
		result := sum(cons(test.input.a, test.input.b))
		if result != test.expected {
			t.Fatalf(
				"Failed test %s: Got %d, expected %d",
				test.name,
				result,
				test.expected,
			)
		}
	}
}

func Test_Car(t *testing.T) {
	tests := []testType{
		{input: pairType{3,4}, expected: 3, name: "Test_Car_1",},
		{input: pairType{1,2}, expected: 1, name: "Test_Car_2",},
		{input: pairType{0,100}, expected: 0, name: "Test_Car_3",},
		{input: pairType{9,1}, expected: 9, name: "Test_Car_4",},
	}

	for _, test := range tests {
		result := car(cons(test.input.a, test.input.b))
		if result != test.expected {
			t.Fatalf(
				"Failed test %s: Got %d, expected %d",
				test.name,
				result,
				test.expected,
			)
		}
	}
}


func Test_Cdr(t *testing.T) {
	tests := []testType{
		{input: pairType{3,4}, expected: 4, name: "Test_Cdr_1",},
		{input: pairType{1,2}, expected: 2, name: "Test_Cdr_2",},
		{input: pairType{0,100}, expected: 100, name: "Test_Cdr_3",},
		{input: pairType{9,1}, expected: 1, name: "Test_Cdr_4",},
	}

	for _, test := range tests {
		result := cdr(cons(test.input.a, test.input.b))
		if result != test.expected {
			t.Fatalf(
				"Failed test %s: Got %d, expected %d",
				test.name,
				result,
				test.expected,
			)
		}
	}
}
