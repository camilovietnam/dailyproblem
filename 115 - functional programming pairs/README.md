# Problem:

```
Good morning! Here's your coding interview problem for today.

This problem was asked by Jane Street.

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

Implement car and cdr
```

## Solution:

I didn't understand the functional programming part of this problem, so I have basically copied the solution from [this address](https://galaiko.rocks/posts/dcp/problem-5/)

Functional programming for me is a little "backwards". Instead of thinking about returning variables and operating on variables, you have to return and operate on functions which in turn operate on something. This layered thinking can sometimes really throw me off. Perhaps it is a good sign that I need to study a little bit more about Functional Programming.

## Running the code
Run the command `npm test -v` to run the tests. The tests use the functions declared in the file `main.go`.
