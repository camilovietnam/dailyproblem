const timestamps = [
	{ timestamp: 1598083514, count: 2, type: "enter"},// 2
	{ timestamp: 1598083934, count: 3, type: "enter"},// 5
	{ timestamp: 1598084174, count: 1, type: "enter"},// 6
	{ timestamp: 1598084234, count: 1, type: "exit"}, // 5
	{ timestamp: 1598084474, count: 1, type: "exit"}, // 4
	{ timestamp: 1598084834, count: 5, type: "enter"},// 9 <- max
	{ timestamp: 1598084954, count: 2, type: "exit"}, // 7
	{ timestamp: 1598085314, count: 3, type: "exit"}, // 4
	{ timestamp: 1598085674, count: 4, type: "exit"}, // 0
];

const more_timestamps = [
	{ timestamp: 1598092874, count: 2, type: "enter"}, // 2
	{ timestamp: 1598092994, count: 3, type: "enter"}, // 5
	{ timestamp: 1598093114, count: 4, type: "enter"}, // 9
	{ timestamp: 1598093174, count: 1, type: "exit"},  // 8
	{ timestamp: 1598093294, count: 2, type: "enter"}, // 10
	{ timestamp: 1598093474, count: 1, type: "exit"},  // 9
	{ timestamp: 1598093774, count: 3, type: "enter"}, // 12
	{ timestamp: 1598093834, count: 1, type: "exit"},  // 11
	{ timestamp: 1598093902, count: 3, type: "enter"}, // 14
	{ timestamp: 1598094121, count: 5, type: "enter"}, // 19 <- max
	{ timestamp: 1598094262, count: 2, type: "exit"},  // 17
	{ timestamp: 1598094982, count: 3, type: "exit"},  // 14
	{ timestamp: 1598095342, count: 4, type: "exit"},  // 10
	{ timestamp: 1598095522, count: 1, type: "enter"}, // 11
	{ timestamp: 1598095762, count: 4, type: "exit"},  // 7
	{ timestamp: 1598095942, count: 3, type: "exit"},  // 4
	{ timestamp: 1598096122, count: 4, type: "exit"},  // 0
];


function busiestPeriod (timestamps) {
	let max = count = 0;
	let max_start = max_end = null;
	let in_max = false;

	timestamps.forEach(timestamp => {
		count = (timestamp.type ==="enter") 
			? count + timestamp.count
			: count - timestamp.count;
		
		if (count > max) {
			max = count;
			max_start = timestamp.timestamp; 
			in_max = true;

			return;
		} 

		if (in_max) {
			in_max = false;
			max_end = timestamp.timestamp;
		}
	});

	return [max_start, max_end, max];
}

/*
*
*	INDEX.JS
*	It will calculate the busiest period and print on screen
*
*/

const busy_first = busiestPeriod (timestamps);
console.log(`First busiest period: ${busy_first[0]} ~ ${busy_first[1]}, ${busy_first[2]} people.`)

const busy_second = busiestPeriod(more_timestamps);
console.log(`Second busiest period: ${busy_second[0]} ~ ${busy_second[1]}, ${busy_second[2]} people.`);