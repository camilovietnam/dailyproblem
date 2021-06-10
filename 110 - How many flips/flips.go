// Good morning! Here's your coding interview problem for today.

// This problem was asked by LinkedIn.

// You are given a string consisting of the letters x and y, such as xyxxxyxyy.
// In addition, you have an operation called flip, which changes a single x to
// y or vice versa.

// Determine how many times you would need to apply this operation to ensure
// that all x's come before all y's. In the preceding example, it suffices to
// flip the second and sixth characters, so you should return 2.

package main

import "fmt"

func main() {
	sequence := []byte{'x', 'y', 'y', 'x', 'x', 'y', 'x', 'x', 'x', 'x', 'y'}

	fmt.Printf("The original sequence: %v\n", string(sequence))
	steps, flipped := swap(sequence)
	fmt.Printf("To arrive to the sorted sequence %v it took %d steps.\n", string(flipped), steps)

}

func swap(sequence []byte) (int, []byte) {
	var flipped = sequence
	var steps int

	for i := len(flipped) - 1; i > 0; i-- {
		if flipped[i] == 'y' {
			continue
		} else {
			nextY := i - 1

			// find where the next Y is
			for nextY >= 0 && flipped[nextY] == 'x' {
				nextY--
			}

			// if we exhausted all possibilities
			if nextY == -1 {
				i = 0
				continue
			}

			// need to flip i and nextY
			flipped[i] = flip(flipped, i)
			flipped[nextY] = flip(flipped, nextY)
			steps += 2
		}
	}

	return steps, flipped
}

func flip(sequence []byte, position int) byte {
	if sequence[position] == 'x' {
		return byte('y')
	} else {
		return byte('x')
	}
}
