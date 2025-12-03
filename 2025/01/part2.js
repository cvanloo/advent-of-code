Number.prototype.mod = function (n) {
  "use strict";
  return ((this % n) + n) % n;
};

document.getElementsByTagName('pre')[0].innerText.split('\n').filter(line => line !== '').reduce((acc, line) => {
  let ops = {
    "L": (a, b) => {
      let m = (a - b).mod(100);
      let z = Math.floor(b / 100);
      z += (m > a ? 1 : 0);
      return [m, z];
    },
    "R": (a, b) => {
      let m = (a + b).mod(100);
      let z = Math.floor(b / 100);
      z += (m < a ? 1 : 0);
      return [m, z];
    }
  };
  let dial = acc[0];
  let zeroes = acc[1];
  [dial, newZeroes] = ops[line[0]](dial, parseInt(line.substring(1)));
  return [dial, zeroes + newZeroes];
}, [50, 0]);
