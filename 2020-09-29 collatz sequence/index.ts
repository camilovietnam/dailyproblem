// A Collatz sequence in mathematics can be defined as follows. Starting with any positive integer:

// if n is even, the next number in the sequence is n / 2
// if n is odd, the next number in the sequence is 3n + 1
// It is conjectured that every such sequence eventually reaches the number 1. Test this conjecture.

// Bonus: What input n <= 1000000 gives the longest sequence?

const log: any = console.log;

const collatz: any = (number: number) => {
    let steps:number = 0;
    while (number !=1) {
        number = (number%2 === 0) ? number/2 :3*number +1;
        steps++;
    }

    return [number, steps];
}

log('*** Let\'s Collatz! ***');
let hypothesis:boolean = true;

let i = 2;
const limit = 1000000;
let [end, steps] = collatz(i)
let maxSteps = steps;
let maxStepNumber = i;

for (i++; i < limit; ++i) {
    let [end, steps] = collatz(i)

    if (end != 1) {
        log ('We found a collatz different than 1 for ' + i);
        hypothesis = false;
        break;
    }

    if (steps > maxSteps) {
        maxSteps = steps;
        maxStepNumber = i;
    }
}

log('Our hypothesis was: ' + ((hypothesis)?'true':'false'));
log(`The longest collatz sequence under ${limit} was ${maxSteps} for value ${maxStepNumber}`);