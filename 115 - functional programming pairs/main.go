Good morning! Here's your coding interview problem for today.

This problem was asked by Jane Street.

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and
last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4))
returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

Implement car and cdr

Solution:
I didn't understand the functional programming part of this problem, so I have basically
copied the solution from https://galaiko.rocks/posts/dcp/problem-5/

Functional programming for me is a little "backwards". Instead of thinking about
returning variables and operating on variables, you have to return and operate
on functions which in turn operate on something. This layered thinking can
sometimes really throw me off. Perhaps it is a good sign that I need to study a
little bit more about Functional Programming.

Anyway, here's the code.

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