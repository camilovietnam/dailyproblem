package main

import "github.com/k0kubun/pp"

/********************************************************************
	Given a binary tree of integers, find the maximum path sum
	between two nodes. The path must go through at least one node,
	and does not need to go through the root

	My solution:
	- First, I will build a list of all the nodes. Then, I will loop
	twice over the list:
		for _, nodeI := range list {
			for _, nodeJ := range list {
				// calculate cost from nodeI to nodeJ
				// save this cost in an array cost[nodeI][nodeJ]
			}
		}
	- After having the array of costs, we just have to traverse this map
		and find the maximum cost, and then note the I,J indices.

********************************************************************/

type node struct {
	Parent *node
	Left   *node
	Right  *node
	Value  int
}

func newChild(value int, parent *node) *node {
	return &node{
		Parent: parent,
		Left:   nil,
		Right:  nil,
		Value:  value,
	}
}

// This is the tree we will build with this method:
//     1
//  2    3
// 6 5  9 8
func buildTree() *node {
	root := &node{
		Value: 1,
	}

	root.Left = newChild(2, root)
	root.Right = newChild(3, root)

	root.Left.Left = newChild(6, root.Left)
	root.Left.Right = newChild(5, root.Left)

	root.Right.Left = newChild(9, root.Right)
	root.Right.Right = newChild(8, root.Right)

	return root
}

func listChildren(n *node) []*node {
	if n.Left == nil && n.Right == nil {
		return nil
	}

	var list = []*node{n}

	if n.Left != nil {
		list = append(list, n.Left)
		list = append(list, listChildren(n.Left)...)
	}

	if n.Right != nil {
		list = append(list, n.Right)
		list = append(list, listChildren(n.Right)...)
	}

	return list
}

func main() {
	root := buildTree()
	children := listChildren(root)

	pp.Println("The list is: ", children)
}
