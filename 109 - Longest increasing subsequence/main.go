// Describe an algorithm to compute the longest increasing
// subsequence of an array of numbers in O(n log n) time.

// My solution:
// I'm not sure I'm correct, but I believe O(n) is faster than
// O(n log n), at least for arrays with more than 3 elements.
// You can see the plot of the two functions in here:
// https://www.wolframalpha.com/input/?i=y%3Dx*logx%2C+y%3Dx
//
// So I wrote this O(n) algorithm which just needs n steps to
// calculate the longest increasing subsequence.

package main

import "fmt"

func getLongestSequence(numbers []int) (maxStart int, maxEnd int) {
	maxLength, currentLength := 1, 1
	currentStart, maxStart := 0, 0
	index := 1

	for index <= len(numbers)-1 {
		if numbers[index] >= numbers[index-1] {
			currentLength++

			if currentLength > maxLength {
				maxStart = currentStart
				maxLength = currentLength
			}
		} else {
			currentStart = index
			currentLength = 1
		}

		index++
	}

	maxEnd = maxStart + maxLength

	return
}

func main() {
	numbers := []int{1, 1, 1, 2, 1, 5, 7, 9, 13, 3, 3, 3}
	maxStart, maxEnd := getLongestSequence(numbers)

	fmt.Printf("Longest sequence starts at %v with %v digits: %v", maxStart, maxEnd-maxStart, numbers[maxStart:maxEnd])
}
