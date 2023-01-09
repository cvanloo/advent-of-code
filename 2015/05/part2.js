var input = document.getElementsByTagName('pre')[0].innerText;
var lines = input.split('\n');

function hasPair(text) {
  let pairs = new Map();

  for (let i = 1; i < text.length; ++i) {
    let pair = text[i - 1] + text[i];
    let idx = pairs.get(pair);

    if (idx !== undefined) {
      if (idx < (i - 1) || idx > i) {
        return true;
      }
    }
    pairs.set(pair, i);
  }
  return false;
}

function hasRepetition(text) {
  for (let i = 1; i < text.length - 1; ++i) {
    let prev = text[i - 1];
    let next = text[i + 1];
    if (prev == next) {
      return true;
    }
  }

  return false;
}

function isNice(text) {
  return hasPair(text) && hasRepetition(text);
};

console.log("count before", lines.length);
lines = lines.filter(isNice);
console.log("count after", lines.length); // 53
