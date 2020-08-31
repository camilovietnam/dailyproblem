const log = console.log;

function findPalindromes(string, length) {
    let start = 0;
    let end = length - 1;
    let palindromes = [];

    while (end < string.length) {
        const candidate = string.substr(start, length);
        const reverse = candidate.split('').reverse().join('');
        if (candidate === reverse) {
            palindromes.push(candidate);
        }
        end = start + length;
        start = start + 1;
    }

    return palindromes;
}

function findFewerPalindromes(queryString) {
    let palindromesFound = [];
    let candidateLength = queryString.length;

    while (candidateLength > 0) {
        const newPalindromes = findPalindromes(queryString, candidateLength);

        newPalindromes.forEach(palindrome => {
            if (queryString.includes(palindrome)) {
                queryString = queryString.replace(palindrome, '');
                palindromesFound.push(palindrome);
            }
        })

        if (candidateLength > queryString.length) {
            candidateLength = queryString.length;
        } else {
            candidateLength--;
        }
    }
    return palindromesFound;
}

const queryString = 'racecarannakayak';
const fewerPalindromes = findFewerPalindromes(queryString);


log(`Here are the fewer palindromes:`, fewerPalindromes);