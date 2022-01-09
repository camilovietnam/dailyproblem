package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

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

// We need Parent to be able to traverse the tree upwards
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

// Will build this tree:
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

// Will return a flat list we can iterate over
func flattenTree(n *node) []*node {
	if n.Left == nil && n.Right == nil {
		return nil
	}

	var list = []*node{n}

	if n.Left != nil {
		list = append(list, n.Left)
		list = append(list, flattenTree(n.Left)...)
	}

	if n.Right != nil {
		list = append(list, n.Right)
		list = append(list, flattenTree(n.Right)...)
	}

	return list
}

func getCost(a, b *node) int {
	// First the cases where one or both nodes have no parents
	
	// - A is the root
	if a.Parent == nil && b.Parent != nil {
		return b.Value + getCost(a, b.Parent)
	}

	// - B is the root
	if b.Parent == nil && a.Parent != nil {
		return a.Value + getCost(a.Parent, b)
	}

	// - Both nodes are root
	if a.Parent == nil && b.Parent == nil {
		return a.Value
	}

	// Second, cases where both nodes have a parent

	// - B is the parent of A
	if a.Parent.Value == b.Value || b.Parent.Value == a.Value{
		return a.Value + b.Value
	}

	// - A is the parent of B
	if a.Parent.Value == b.Parent.Value {
		return a.Value + b.Value + a.Parent.Value
	}

	// - Both nodes have different parents
	return a.Value + b.Value + getCost(a.Parent, b.Parent)
}

func main() {
	var (
		root = buildTree()
		children = flattenTree(root)
		costs = make (map[int]map[int]int, len(children))
		maxCost, maxOrigin, maxEnd int
	)
	
	for _, nodeA := range children {
		costs[nodeA.Value] = make(map[int]int, len(children))

		for _, nodeB := range children {
			// ignore cost of going from one node to the node itself
			if nodeA.Value == nodeB.Value {
				continue
			}

			var cost int = getCost(nodeA, nodeB)

			if cost > maxCost {
				maxCost = cost
				maxOrigin = nodeA.Value
				maxEnd = nodeB.Value
			}
		}
	}

	pp.Println(fmt.Sprintf("Max cost: from node %d to node %d: (Cost: %d)", maxOrigin, maxEnd, maxCost))
}
