/**
 * Solution 1: This solution views a a number as a range of digits, with
 * the left margin being 2, and the right margin being the number itself.
 * It then finds the middle point, and asks whether the square root is the
 * middle itself, whether it exists in the first half of the range, to the
 * left of the middle point, (left margin ~ middle) or whether it exists in
 * the second half of the range (middle ~ right margin).
 * In the two latter cases, it begins the process of searching again, using
 * the new halved range.
 *
 * Eventually the middle point will either match the root (verified by
 * multiplying the middle against itself), or match one of the margins, left
 * or right. In any case, we return the value of middle.
 * This means the algorithm will not find rational roots, only integers.
 * Example:
 * square(81) = 9
 * square(82) = 9
 * square(99) = 0
 * square(10) = 100
 */

const log = console.log;

const square = number => {
    let left = 2
    let middle = number/2
    let right = number

    while (middle*middle != number && middle != right && middle != left) {
        if (middle*middle > number) {
            right = middle
        } else {
            left = middle
        }

        middle = parseInt((right + left)/2)
    }

    return middle;
}

let args = process.argv.slice(2);
const number = parseInt(args[0])
const squareRoot = square(number)

log(`The square root of ${number} is ${squareRoot}`)