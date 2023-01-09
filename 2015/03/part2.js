function createKey(coords) {
  return `${coords.x}@${coords.y}`;
}

var input = document.getElementsByTagName('pre')[0].innerText;
var directions = input.split('');

var houses = new Map();
var santa = {
  x: 0,
  y: 0
}
var robo = {
  x: 0,
  y: 0
}
var current = santa;

houses.set(createKey(current), 1); // starting position
directions.forEach(direction => {
  switch (direction) {
    case "^":
      current.y++;
      break;
    case ">":
      current.x++;
      break;
    case "v":
      current.y--;
      break;
    case "<":
      current.x--;
      break;
  }
  let tmp = houses.get(createKey(current)) ?? 0;
  houses.set(createKey(current), ++tmp);
  current = current === santa ? robo : santa;
});
console.log("Part 2:", houses.size); // 2341
