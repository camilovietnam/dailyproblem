// Good morning! Here's your coding interview problem for today.
// This problem was asked by Etsy.
// Given a sorted array, convert it into a height-balanced binary search tree.


/*
	Solution:
	Since the array is ordered, we can pretty much ensure that the element in the
	middle will always be the highest node in a branch. So for example, for [1,2,3] we
	know that the "2" in the middle will be the root, and the remaining numbers on the
	left and right will become the left and right branches. So, we can think of an
	array as this:
	[branch], [root], [branch]

	And one branch could be converted into a tree depending on its length:
	Length 1 (as in [1]): the branch becomes a node with "nil" Left and Right branches
	Length 2 (as in [1,2]: the branch becomes a node with one branch. Since the numbers
		are ordered, we can take the second element as root (biggest one) and the other as
		left branch (smaller than the root)
	Any length above 2 we will again consider as fitting the pattern:
	[branch], [root], [branch]
	And we will recursively process the branches until we hit one of the base cases
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	myArray := []int{1,2,3}
	tree := arrayToTree(myArray)

	fmt.Println("Root:", tree.Value)
	fmt.Println("Left:", tree.Left.Value)
	fmt.Println("Right:", tree.Right.Value)
}

func arrayToTree(sortedArray []int) *node {
	// Base cases
	switch len(sortedArray) {
	case 1:
		return &node {
			Value: sortedArray[0],
		}
	case 2:
		return &node{
			Value: sortedArray[1],
			Left: &node {
				Value: sortedArray[0],
			},
		}
	}

	middlePoint := middlePoint(len(sortedArray))

	root := &node{
		Left: arrayToTree(sortedArray[0:middlePoint]),
		Right: arrayToTree(sortedArray[middlePoint+1:]),
		Value: sortedArray[middlePoint],
	}

	return root
}

func middlePoint(length int) int {
	return int(math.Round(float64(length/2)))
}

