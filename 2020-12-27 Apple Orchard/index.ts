// Good morning! Here's your coding interview problem for today.

// This problem was asked by Google.

// A girl is walking along an apple orchard with a bag in each hand. She likes to pick apples 
// from each tree as she goes along, but is meticulous about not putting different kinds of 
// apples in the same bag.

// Given an input describing the types of apples she will pass on her path, in order, determine
// the length of the longest portion of her path that consists of just two types of apple trees.

// For example, given the input [2, 1, 2, 3, 3, 1, 3, 5], the longest portion will involve types
// 1 and 3, with a length of four.

const _ = console.log;

let path = [2, 1, 2, 3, 3, 1, 1, 3, 5];
let lastStack: number[] = [];
let currentStack: number[] = [];

// resolve first 2 numbers
if (path[0] === path[1]) {
    currentStack = [path[0], path[0]];
} else {
    lastStack = [path[0]];
    currentStack = [path[1]];
}

let currentMax: number = 2;
let totalMax: number = currentMax;

for (let i: number=2; i < path.length; ++i) {
    // if the number changes back to lastStack
    if (path[i] === lastStack[0]) {
        lastStack = currentStack;
        currentStack = [path[i]];
        currentMax++;
        // _(`Read [${path[i]}]: Added to a new Stack, Local Max: ${currentMax}. Stacks: `, lastStack, currentStack)
        continue;
    }

    // if the number continues in currentStack
    if (path[i] === currentStack[0]) {
        currentMax++;
        currentStack.push(path[i]);
        // _(`Read [${path[i]}]: Added to current Stack, Local Max: ${currentMax}. Stacks:`, lastStack, currentStack)
        continue;
    }

    // if we change to a new number
        //  update Max
    if (currentMax > totalMax) {
        totalMax = currentMax;
    }

    // change Stacks
    lastStack = currentStack;
    currentStack = [path[i]];
    currentMax = 1 + lastStack.length;
    
    // _(`Read [${path[i]}]: Added to a new Stack, Local Max: ${currentMax}. Stacks:`, lastStack, currentStack)
}

_(`The longest path is: (${totalMax} trees)`)