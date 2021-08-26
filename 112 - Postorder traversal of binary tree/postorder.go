// Good morning! Here's your coding interview problem for today.

// This problem was asked by Google.

// Given the sequence of keys visited by a postorder traversal of 
// a binary search tree, reconstruct the tree.

// For example, given the sequence 2, 4, 3, 8, 7, 5, you should 
// construct the following tree:

//     5
//    / \
//   3   7
//  / \   \
// 2   4   8

// *** Solution ***

// We start with a tree with one root (the first number). From here, we compare
// the root to the next number in the sequence. 

// If the number is bigger: New number becomes the root, old root becomes left branch 
// If the number is smaller: New number becomes the root, old root becomes the right
// 		branch, left branch of old root becomes left branch of new root

package main 

import "fmt"

type node struct {
	Value int
	LeftBranch *node
	RightBranch *node
}

func main () {
	var sequence = "243875"
	var root *node
	
	root = buildTree(sequence)
	testTree(root)
}

func buildTree(sequence string) *node {
	var root *node

	for _, c := range sequence {
		value := parseInt(c)
		// case 1: there is no root yet
		if root == nil {
			root = &node {
				Value: value,
			}

			continue
		}

		// case 2: new number is bigger
		if value > root.Value {
			root = &node {
				Value: value,
				LeftBranch: root,
			}
		}

		// case 3: new number is smaller
		if value < root.Value {
			root = &node {
				Value: value,
				LeftBranch: root.LeftBranch,
				RightBranch: root,
			}
			root.RightBranch.LeftBranch = nil
		}
	}

	return root
}

func parseInt(char rune) int {
	return int (char - '0')
}

func testTree(root *node) {
	fmt.Println("Root is: ", root.Value)
	fmt.Println("Left branch of root:", root.LeftBranch.Value)
	fmt.Println("Left branch of Left branch:", root.LeftBranch.LeftBranch.Value)
	fmt.Println("Right branch of Left branch", root.LeftBranch.RightBranch.Value)
	fmt.Println("Right branch of root:", root.RightBranch.Value)
	fmt.Println("Right branch of Right branch:", root.RightBranch.RightBranch.Value)	
}