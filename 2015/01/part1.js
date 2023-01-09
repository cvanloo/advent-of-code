// Run from input site:
// https://adventofcode.com/2015/day/1/input
var input = document.getElementsByTagName('pre')[0].innerHTML;
var floor = 0;

input.split('').forEach(el => {
  if (el === '(') {
    floor++
  } else if (el === ')') {
    floor--
  }
});
console.log(floor) // 280