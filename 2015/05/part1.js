var input = document.getElementsByTagName('pre')[0].innerText;
var lines = input.split('\n');

const isNice = (text) => {
  const vowels = ['a', 'e', 'i', 'o', 'u'];
  const forbidden = ["ab", "cd", "pq", "xy"];
  
  let hasVowels = text.split('').filter(r => vowels.includes(r));
  if (hasVowels.length < 3) {
    return false;
  }
  
  let repetitionFound = false;
  for (let i = 1; i < text.length; ++i) {
    if (text[i-1] == text[i]) {
      repetitionFound = true;
    }
    let word = text[i-1] + text[i];
    if (forbidden.includes(word)) {
			return false;
    }
  }
  if (!repetitionFound) {
    return false;
  }
  
  return true;
};

console.log("count before", lines.length);
lines = lines.filter(isNice);
console.log("count after", lines.length); // 258
