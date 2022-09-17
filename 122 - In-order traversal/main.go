package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
	prev  *node
}

func main() {
	tree := buildTree()
	fmt.Println("Orden h: ")
	inOrderTraversal(tree)
	fmt.Println("Orden 1: ")
	inOrderTraversalCheap(tree)
}

func buildTree() *node {
	root := &node{
		value: 1,
	}

	root.left = &node{
		value: 2,
		prev:  root,
	}

	root.right = &node{
		value: 3,
		prev:  root,
	}

	root.left.left = &node{
		value: 4,
		prev:  root.left,
	}

	root.left.right = &node{
		value: 5,
		prev:  root.left,
	}

	root.left.right.left = &node{
		value: 6,
		prev:  root.left.right,
	}
	root.left.right.right = &node{
		prev:  root.left.right,
		value: 7,
	}

	return root
}

func inOrderTraversal(root *node) {
	if root.left != nil {
		inOrderTraversal(root.left)
	}
	fmt.Printf("%d,", root.value)

	if root.right != nil {
		inOrderTraversal(root.right)
	}
}

func inOrderTraversalCheap(root *node) {
	current := root

	for {
		for current.left != nil {
			current = current.left
		}

		fmt.Printf("%d,", current.value)

		for current.prev.left.value != current.value {
			current = current.prev
			if current.value == root.value {
				fmt.Println("  /end")
				return
			}
		}
		current = current.prev
		fmt.Printf("%d,", current.value)
		current = current.right
	}
}
