Number.prototype.mod = function(n) {
  "use strict";
  return ((this % n) + n) % n;
};

document.getElementsByTagName('pre')[0].innerText
  .split('\n')
  .filter(line => line !== '')
  .reduce((acc, line) => {
    let ops = {
      "L": (a, b) => (a - b).mod(100),
      "R": (a, b) => (a + b).mod(100)
    };
    let dial = acc[acc.length - 1];
    acc.push(ops[line[0]](dial, parseInt(line.substring(1))));
    return acc;
  }, [50])
  .filter(d => d === 0)
  .length;
