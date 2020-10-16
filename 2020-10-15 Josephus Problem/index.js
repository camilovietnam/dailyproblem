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

    this.prisoners = Array.from(Array(n+1).keys()).splice(1)
    this.step = k;
  }

  kill(prisoner) {
    const killed = this.prisoners[prisoner];
    this.prisoners.splice(prisoner, 1);
    return killed;
  }

  next(prisoner, step = 1) {
      return (prisoner + step) % this.prisoners.length;
  }

  display() {
    log (`Prisoner list: ${this.prisoners}`);
  }
}

/* Execution starts here */
const args = process.argv.slice(2);
if (args.length < 2) {
  log('??? Please provide arguments N, k. \nExample: node index.js 5 2');
  return;
}
const n = parseInt(args[0]);
const k = parseInt(args[1]);

const p = new Prisoners(n, k);
log(`Total prisoners: ${p.prisoners.length}. Step: ${p.step}`);

let killIndex = 0;
const killed = [];

while (p.prisoners.length > 1) {
  const nextToKill = p.next(killIndex - 1, p.step)
  killIndex = nextToKill
  killed.push(p.kill(killIndex))
}

log(`The killed list is: ${killed}`);
log(`Survivor is ${p.prisoners[0]}`)