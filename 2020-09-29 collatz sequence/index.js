// A Collatz sequence in mathematics can be defined as follows. Starting with any positive integer:
// if n is even, the next number in the sequence is n / 2
// if n is odd, the next number in the sequence is 3n + 1
// It is conjectured that every such sequence eventually reaches the number 1. Test this conjecture.
// Bonus: What input n <= 1000000 gives the longest sequence?
var log = console.log;
var collatz = function (number) {
    var steps = 0;
    while (number != 1) {
        number = (number % 2 === 0) ? number / 2 : 3 * number + 1;
        steps++;
    }
    return [number, steps];
};
log('*** Let\'s Collatz! ***');
var hypothesis = true;
var i = 2;
var limit = 1000000;
var _a = collatz(i), end = _a[0], steps = _a[1];
var maxSteps = steps;
var maxStepNumber = i;
for (i++; i < limit; ++i) {
    var _b = collatz(i), end_1 = _b[0], steps_1 = _b[1];
    if (end_1 != 1) {
        log('We found a collatz different than 1 for ' + i);
        hypothesis = false;
        break;
    }
    if (steps_1 > maxSteps) {
        maxSteps = steps_1;
        maxStepNumber = i;
    }
}
log('Our hypothesis was: ' + ((hypothesis) ? 'true' : 'false'));
log("The longest collatz sequence under " + limit + " was " + maxSteps + " for value " + maxStepNumber);
