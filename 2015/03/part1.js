function createKey(x, y) {
  return `${x}@${y}`;
}

var input = document.getElementsByTagName('pre')[0].innerText;
var directions = input.split('');

var houses = new Map();
var [x, y] = [0, 0];

houses.set(createKey(x, y), 1); // starting position
directions.forEach(direction => {
  switch (direction) {
    case "^":
      y++;
      break;
    case ">":
      x++;
      break;
    case "v":
      y--;
      break;
    case "<":
      x--;
      break;
  }
  let tmp = houses.get(createKey(x, y)) ?? 0;
  houses.set(createKey(x, y), ++tmp);
});
console.log("Part 1:", houses.size); // 2081
