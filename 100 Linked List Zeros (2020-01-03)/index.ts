// Good morning! Here's your coding interview problem for today.

// This problem was asked by Amazon.

// Given a linked list, remove all consecutive nodes that sum to 
// zero. Print out the remaining nodes.

// For example, suppose you are given the input 
// 3 -> 4 -> -7 -> 5 -> -6 -> 6. In this case, you should first 
// remove 3 -> 4 -> -7, then -6 -> 6, leaving only 5.

// My Solution
// I will basically check for every element in the array of numbers,
// if it is the starting point of a zero-sum sequence of digits. If 
// it is, then I will remove all the elements that add up to zero. 
const log: any = console.log;
let numbers : number[]  = [3, 4, -7, 5, -6, 6];

let i:number = 0;
while (i < numbers.length-1) {

	let sum = numbers[i];
	let j = i + 1;

	// Keep adding up until we exhaust the array or the sum is zero
	while (sum !== 0 && j < numbers.length) {
		sum += numbers[j];
		++j;
	}

	// if sum is zero, we need to remove the sequence
	if (sum === 0) {
		numbers.splice(i, j-i);
	} else {
		++i;
	}
}


log (`The final array is`, numbers);