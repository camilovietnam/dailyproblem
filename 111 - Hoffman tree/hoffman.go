// Problem #528

// Huffman coding is a method of encoding characters based on their
// frequency. Each letter is assigned a variable-length binary string,
// such as 0101 or 111110, where shorter lengths correspond to more
// common letters. To accomplish this, a binary tree is built such that
// the path from the root to any leaf uniquely maps to a character. When
// traversing the path, descending to a left child corresponds to a 0 in
// the prefix, while descending right corresponds to 1.

// Here is an example tree (note that only the leaf nodes have letters):

//         *
//       /   \
//     *       *
//    / \     / \
//   *   a   t   *
//  /             \
// c               s

// With this encoding, cats would be represented as 0000110111.

// Given a dictionary of character frequencies, build a Huffman tree, and
// use it to determine a mapping between characters and their encoded binary
// strings.

package main

import "fmt"

type leafType struct {
	left  *leafType
	right *leafType
	value byte
}

type dictionaryType map[byte]string

func main() {
	// I will manually create a dictionary, then build a tree from it, then
	// build a new dictionary from the tree.

	myDict := buildExampleDict()
	root := buildTreeFromDict(myDict)

	// Let's test the tree
	testCatsTree(root)

	myNewDict := buildDictFromTree(root)
	testDictionary(myNewDict)
}

func buildExampleDict() dictionaryType {
	var myDict dictionaryType = dictionaryType{}
	myDict['c'] = "000"
	myDict['a'] = "01"
	myDict['t'] = "10"
	myDict['s'] = "111"

	return myDict
}

func buildTreeFromDict(myDict dictionaryType) *leafType {
	var root *leafType = &leafType{}

	for char, word := range myDict {
		addToTree(char, word, root)
	}

	return root
}

func addToTree(char byte, word string, root *leafType) {
	var current *leafType = root

	for _, char := range word {
		initCurrentLeaf(current)

		if char == '0' {
			current = current.left
		}

		if char == '1' {
			current = current.right
		}
	}

	current.value = char
}

func initCurrentLeaf(leaf *leafType) {
	if leaf.left == nil {
		leaf.left = &leafType{}
	}

	if leaf.right == nil {
		leaf.right = &leafType{}
	}
}

func buildDictFromTree(root *leafType) dictionaryType {
	var dict dictionaryType = dictionaryType{}

	processNode("", root, dict)

	return dict
}

func processNode(prefix string, node *leafType, dict dictionaryType) {
	// if i find a node without children, then I take the value
	if node.left == nil && node.right == nil {
		dict[node.value] = prefix
	}

	// if the node has left children
	if node.left != nil {
		processNode(prefix+"0", node.left, dict)
	}

	// if the node has right children
	if node.right != nil {
		processNode(prefix+"1", node.right, dict)
	}
}

func testCatsTree(root *leafType) {
	fmt.Println(">> Testing the build tree: \n")
	fmt.Printf("  This should be C: %c\n", root.left.left.left.value)
	fmt.Printf("  This should be A: %c\n", root.left.right.value)
	fmt.Printf("  This should be T: %c\n", root.right.left.value)
	fmt.Printf("  This should be S: %c\n", root.right.right.right.value)
}

func testDictionary(myDict dictionaryType) {
	fmt.Println("\n>> Testing the dictionary: \n")
	fmt.Println("  The value for c is: ", myDict['c'])
	fmt.Println("  The value for a is: ", myDict['a'])
	fmt.Println("  The value for t is: ", myDict['t'])
	fmt.Println("  The value for s is: ", myDict['s'])
}
