input = document.getElementsByTagName('pre')[0].innerText;

rows = input.split('\n').filter(l => l !== '');
nrows = rows.length;
ncols = rows[0].length;

ops = {
  "+": (a, b) => a + b,
  "*": (a, b) => a * b,
}

var strNum = '';
var nums = [];
var op = '';
var problem = [];
for (let col = ncols - 1; col >= 0; --col) {
  let allSpace = true;
  for (let row = 0; row < nrows; ++row) {
    let c = rows[row][col];
    if (c !== ' ') {
      allSpace = false;
    }
    if (0 <= c && c <= 9) {
      strNum += c;
    } else if (c == '*' || c == '+') {
      op = c;
    }
  }
  if (allSpace) {
    problem.push([ops[op], nums]);
    op = '';
    nums = [];
  } else {
    nums.push(parseInt(strNum));
    strNum = '';
  }
}
problem.push([ops[op], nums]);

problem
  .map(([op, nums]) => nums.reduce((a, n) => op(a, n)))
  .reduce((a, n) => a + n);
