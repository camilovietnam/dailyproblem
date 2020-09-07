const log:any = console.log;

// Given an array of elements, return the length of the longest subarray where 
// all its elements are distinct.

// For example, given the array [5, 1, 3, 5, 2, 3, 4, 1], return 5 as the longest
// subarray of distinct elements is [5, 2, 3, 4, 1].


function countDistinct (data:number[]) {
	let distinct:number[] = [];

	data.map((item:number) => {
		distinct[item] = 1;
	});

	return distinct.filter(i => !!i).length;
}

const data:number[] = [5, 1, 3, 5, 2, 3, 4, 1];
const distinct:number = countDistinct(data);

log(`The total of unique elements: ${distinct}`);

