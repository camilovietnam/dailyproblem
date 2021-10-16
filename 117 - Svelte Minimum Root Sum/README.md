# Sorted Squares
Problem solved on october 16, 2021
## Problem description:

```
This question was asked by Apple.
Given a binary tree, find a minimum path sum from root to a leaf.
For example, the minimum path in this tree is [10, 5, 1, -1], which has sum 15.

  10
 /  \
5    5
 \     \
   2    1
       /
     -1
```

Ok, so for this one I decided to build a small script to actually allow the person to introduce the values of the different leaves in the tree, which took more time than writing the solution to the actual problem, because it wasn't necessarily trivial. First, I wrote code to allow the person to add tree levels as html input elements, and then I had to read these inputs and actually convert the table of inputs into a binary tree. 

Once the tree is built, I needed two methods, one to convert each one of the possible paths into an array of numbers (for example, the tree in the question would become [[10,5,2], [10,5,1,-1]]) and then, we basically loop through these arrays and sum the values, updating a `minSum` value if we find the path with the minimum cost. 

## How to run the code in this folder:
- Install dependencies with `npm install`
- Run the web server with `npm run dev`. The app can be visited at [http://localhost:5000](http://localhost:5000)

## Screenshot
![Page Screenshot](https://raw.githubusercontent.com/camilovietnam/dailyproblem/master/117%20-%20Svelte%20Minimum%20Root%20Sum/screenshot.png)
