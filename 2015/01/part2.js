var input = document.getElementsByTagName('pre')[0].innerHTML.split('');
var floor = 0;

input.forEach(el => {
  if (el === '(') {
    floor++
  } else if (el === ')') {
    floor--
  }
});
console.log("Part 1:", floor) // 280

floor = 0;
var position = 0;
while (floor >= 0) {
  let current = input[position]
  if (current === '(') {
    floor++
  } else if (current === ')') {
    floor--
  }
  position++;
}
console.log("Part 2:", position) // 1797
