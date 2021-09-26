package main

import "fmt"

func main () {
	fmt.Println("Ran correctly")
}

type pairOperation func (int, int) int
type funcOperation func(pairOperation) int

func cons(a, b int) funcOperation {
	pair := func (operate pairOperation) int {
		return operate(a,b)
	}
	return pair
}

func sum (runOnPair funcOperation) int {
	addition := func(a, b int) int {
		return a+b
	}
	return runOnPair(addition)
}

func car(runOnPair funcOperation) int {
	selectLeft := func(a, b int) int {
		return a
	}
	return runOnPair(selectLeft)
}

func cdr(runOnPair funcOperation) int {
	selectRight := func (a, b int) int {
		return b
	}
	return runOnPair(selectRight)
}