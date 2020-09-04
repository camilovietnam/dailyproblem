const log: any = console.log;

// Declaration of types
interface Point {
  x: number,
  y: number,
}

interface Dimension {
  width: number,
  height: number,
}

interface Rectangle {
  top_left: Point,
  dimensions: Dimension
}

interface Range {
  start: number,
  end: number
}

const findOverlapX: any = (a: Rectangle, b: Rectangle) => {
  // if not overlap, return 0;
  if ((!existsInRange(a.top_left.x, {
      start: b.top_left.x,
      end: b.top_left.x + b.dimensions.width
    }))
    && (!existsInRange(b.top_left.x, {
      start: a.top_left.x,
      end: a.top_left.x + a.dimensions.width
    }))
  ) {
    return 0;
  }

  // if overlap, return the total of squares
  const minOverlapX = findMinOverlapX(a, b);
  const maxOverlapX = findMaxOverlapX(a, b);

  return maxOverlapX - minOverlapX;
}

const findOverlapY: any = (a: Rectangle, b: Rectangle)  => {
  // if not overlap, return 0;
  if ((!existsInRange(a.top_left.y, {
      start: b.top_left.y - b.dimensions.height,
      end: b.top_left.y,
    }))
    && (!existsInRange(b.top_left.y, {
      start: a.top_left.y - a.dimensions.height,
      end: a.top_left.y
    }))
  ) {
    return 0;
  }

  // if overlap, return the total of squares
  const minOverlapY = findMinOverlapY(a, b);
  const maxOverlapY = findMaxOverlapY(a, b);

  return maxOverlapY - minOverlapY;
}

// helper method to find if a value exists inside a range
const existsInRange: any = (value: number, range: Range) => {
  return (value >= range.start) && (value <= range.end);
}

const findMinOverlapX: any = (a: Rectangle, b:Rectangle) => {
  // this method assumes overlap exists
  if (a.top_left.x >= b.top_left.x) {
    return a.top_left.x
  }

  return b.top_left.x;
}

const findMaxOverlapX: any = (a: Rectangle, b: Rectangle) => {
  // this method assumes overlap exists
  if (a.top_left.x + a.dimensions.width >= b.top_left.x + b.dimensions.width) {
    return b.top_left.x + b.dimensions.width;
  }

  return a.top_left.x + a.dimensions.width;
}

const findMinOverlapY: any = (a: Rectangle, b: Rectangle) => {
  // this method assumes overlap exists
  if (a.top_left.y - a.dimensions.height >= b.top_left.y - b.dimensions.height) {
    return a.top_left.y - a.dimensions.height;
  }
  return b.top_left.y - b.dimensions.height;
}

const findMaxOverlapY: any = (a: Rectangle, b:Rectangle) => {
  // this method assumes overlap exists
  if (a.top_left.y >= b.top_left.y) {
    return b.top_left.y;
  }

  return a.top_left.y;
}

const recA : Rectangle = {
  top_left: {
    x: 1,
    y: 4
  },
  dimensions : {
    width: 3,
    height: 3
  }
}

const recB : Rectangle = {
  top_left: {
    x: 0,
    y: 5
  },
  dimensions : {
    width: 4,
    height: 3
  }
}

const overlapX = findOverlapX(recA, recB);
const overlapY = findOverlapY(recA, recB);

log(`The overlap covers ${overlapX*overlapY} squares`);