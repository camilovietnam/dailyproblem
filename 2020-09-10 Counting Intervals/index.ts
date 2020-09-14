// Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
// Intervals can "touch", such as [0, 1] and [1, 2], but they won't be considered overlapping.
// For example, given the intervals (7, 9), (2, 4), (5, 8), return 1 as the last interval can be removed and the first two won't overlap.
// The intervals are not necessarily sorted in any order.

const log : any = console.log;

interface Interval {
  start: number,
  end: number,
  overlaps: number
}

function sortByLeft(intervals: Interval[]) {
  let copy:Interval[] = intervals.slice(0);

  for (let i:number = 0; i < copy.length; ++i) {
    for (let j:number = i+1; j < copy.length; ++j) {
      if (copy[j].start < copy[i].start) {
        [copy[i], copy[j]] = [copy[j], copy[i]]
      }
    }
  }
  return copy
}

function sortByRight(intervals: Interval[]) {
  let copy:Interval[] = intervals.slice(0)

  for (let i:number = 0; i < copy.length; ++i) {
    for (let j:number = i+1; j < copy.length; ++j) {
      if (copy[j].end < copy[i].end) {
        [copy[i], copy[j]] = [copy[j], copy[i]]
      }
    }
  }
  return copy;
}

function countToTheLeft(interval: Interval, intervals: Interval[]) {
  let sorted:Interval[] = sortByRight(intervals)
  let i:number = 0;

  while (sorted[i].end <= interval.start) {
    ++i
  }

  return i;
}

function countToTheRight(interval: Interval, intervals: Interval[]) {
  let sorted:Interval[] = sortByLeft(intervals)
  let i:number = 0;

  while (i < sorted.length && sorted[i].start < interval.end) {
    ++i
  }

  if (i === sorted.length) {
    return 0;
  }

  return sorted.length - i;
}

function countOverlaps (intervals: Interval[]) {
  for (const i in intervals) {
    intervals[i].overlaps = intervals.length
      - countToTheLeft(intervals[i], intervals)
      - countToTheRight(intervals[i], intervals)
      - 1;
  }
}

function sortByOverlaps(intervals: Interval[]) {
  countOverlaps(intervals)

  for (let i:number = 0; i < intervals.length; ++i) {
    for (let j:number = 0; j < intervals.length; ++j) {
      if (intervals[i].overlaps > intervals[j].overlaps) {
        [intervals[i], intervals[j]] = [intervals[j], intervals[i]]
      }
    }
  }
}

// todo
function hasOverlaps(intervals: Interval []) {
  const sorted:Interval[] = sortByLeft(intervals);

  let overlaps:number = 0;
  let index:number = 0;

  while (overlaps === 0 && index < (sorted.length - 1)) {
    if (sorted[index].end > sorted[index+1].start) {
      overlaps++;
    }

    index++;
  }

  return overlaps > 0;
}

/**
 *       Main Code
 */

const intervals:Interval[] = [
  { start: 2, end: 4, overlaps: null},
  { start: 0, end: 2, overlaps: null},
  { start: 4, end: 7, overlaps: null},
  { start: 1, end: 5, overlaps: null},
  { start: 4, end: 6, overlaps: null}
];

sortByOverlaps(intervals);

let removals = 0;

while (hasOverlaps(intervals)) {
  intervals.shift();
  removals++;
}

log(`The minimum number of removals to eliminate overlapping is ${removals} intervals.`);

// For a second group of intervals

const intervals2:Interval[] = [
  {start: 7, end: 9 , overlaps: null},
  {start: 2, end: 4, overlaps: null},
  {start: 5, end: 8, overlaps: null}
];

sortByOverlaps(intervals2);

removals = 0;

while (hasOverlaps(intervals2)) {
  intervals2.shift();
  removals++;
}

log(`The minimum number of removals to eliminate overlapping is ${removals} intervals.`);