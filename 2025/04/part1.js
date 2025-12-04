input = document.getElementsByTagName('pre')[0].innerText;

nbrs = (x, y) => [
  [x-1, y-1], [x, y-1], [x+1, y-1],
  [x-1, y  ],/*[x, y]*/ [x+1, y  ],
  [x-1, y+1], [x, y+1], [x+1, y+1]
];

grid = input.split('\n').map(r => r.split(''));

adj = 0;
for (let x = 0; x < grid[0].length; ++x) {
  for (let y = 0; y < grid.length; ++y) {
    if (grid[x][y] !== '@') continue;
    let ns = nbrs(x, y).map(([x, y]) => (gx = grid[x], gx !== undefined ? gx[y] : gx));
    let r = ns.filter(n => n === '@').length < 4 ? 1 : 0;
    adj += r;
  }
}

adj;
