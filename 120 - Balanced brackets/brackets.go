package brackets

// This problem was asked by Facebook.
//
// Given a string of round, curly, and square open and closing
// brackets, return whether the brackets are balanced
// (well-formed).
//
// For example, given the string "([])[]({})", you should
// return true.
//
// Given the string "([)]" or "((()", you should return false.
//

func isBalanced(seq string) bool {
	var expectedClose string

	for _, c := range seq {
		if isOpen(c) {
			expectedClose = opposite(c) + expectedClose
		} else {

			if c != rune(expectedClose[0]) || len(expectedClose) == 0 {
				return false
			} else {
				expectedClose = expectedClose[1:]
			}
		}
	}

	return len(expectedClose) == 0 
}

func isOpen(c rune) bool {
	return c == '(' || c == '[' || c == '{'
}

func opposite (c rune) string {
	switch c {
	case '(':
		return ")"
	case '[':
		return "]"
	case '{':
		return "}"
	}

	return ""
}