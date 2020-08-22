function distance(word1, word2, context) {
	let match = context.match(`(${word1}.*${word2})`);

	while (match) {
		context = match[0];
		match = context.match(`(${word1}.*${word2})`);		
	}
	
	const count = context.trim().split(' ').length

	return count -2;
}


const word1 = "shortest"
const word2 = "between"
const context = "the shortest distance between two words is not between many"
const wordDistance = distance(word1, word2, context)

console.log(`Distance between words: ${wordDistance}`)
