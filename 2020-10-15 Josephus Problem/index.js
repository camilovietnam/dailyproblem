/**
 * There are N prisoners standing in a circle, waiting to be executed. The executions are carried out
 * starting with the kth person, and removing every successive kth person going clockwise until there
 * is no one left.
 * Given N and k, write an algorithm to determine where a prisoner should stand in order to be the last
 * survivor.
 * For example, if N = 5 and k = 2, the order of executions would be [2, 4, 1, 5, 3], so you should
 * return 3.
 *
 * Bonus: Find an O(log N) solution if k = 2.
 */

const log = console.log;

class Prisoners {
  constructor (n, k) {
    this.prisoners = [...Array(n).keys()];
    this.step = k;
  }

  kill(prisoner) {
    const killed = this.prisoners[prisoner];
    this.prisoners.splice(prisoner, 1);
    return killed;
  }

  next(prisoner, step = 1) {
      return (prisoner + step) % this.prisoners.length; // -1 because the array is reduced in 1 after killing
  }

  display() {
    log (`Prisoner list: ${this.prisoners}`);
  }
}

const p = new Prisoners(5, 2);
p.display()

let killIndex = 0;
const killed = [];

while (p.prisoners.length > 1) {
  killIndex = p.next(killIndex-1, p.step)
  killed.push(p.kill(killIndex))
}

log(`The killed list is: ${killed}`);
log(`Survivor is ${p.prisoners[0]}`)