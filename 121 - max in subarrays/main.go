// Good morning! Here's your coding interview problem for today.
//
// This problem was asked by Google.
//
// Given an array of integers and a number k, where 1 <= k <= length of the array,
// compute the maximum values of each subarray of length k.
//
// For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8],
// since:
//
// 10 = max(10, 5, 2)
// 7 = max(5, 2, 7)
// 8 = max(2, 7, 8)
// 8 = max(7, 8, 7)
// Do this in O(n) time and O(k) space. You can modify the input array in-place and you
// do not need to store the results. You can simply print them out as you compute them.
package main

import (
	"fmt"
)

func main() {
	numberGroup1 := []int{10, 5, 2, 7, 8, 7}
	result := maxInGroups(numberGroup1, 3)
	fmt.Println("The max values: ", result)

	numbers2 := []int{7, 1, 3, 9, 1, 2, 8, 5, 15, 1, 3}
	result = maxInGroups(numbers2, 3)
	fmt.Println("The max values: ", result)
}

func maxInGroups(numbers []int, groupLength int) []int {
	var result []int

	// we loop over the N elements of the array
	for currentIndex, currentNumber := range numbers {
		// we loop over previous K elements
		for i := 1; i < groupLength; i++ {
			positionOfMax := currentIndex - i
			if positionOfMax < 0 {
				break
			}
			numbers[positionOfMax] = findMax(numbers[positionOfMax], currentNumber)
		}

		if currentIndex >= groupLength-1 {
			result = append(result, numbers[currentIndex-groupLength+1])
		}
	} // the final order is then O(k*N) but given that K is a constant, the order
	// is reduced to O(n)

	return result
}

func findMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
