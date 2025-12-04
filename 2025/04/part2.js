var input = document.getElementsByTagName('pre')[0].innerText;

nbrs = (x, y) => [
  [x-1, y-1], [x, y-1], [x+1, y-1],
  [x-1, y  ],/*[x, y]*/ [x+1, y  ],
  [x-1, y+1], [x, y+1], [x+1, y+1]
];

grid = input.split('\n').map(r => r.split(''));

removed = 0;
do {
  removable = [];
  for (let x = 0; x < grid[0].length; ++x) {
    for (let y = 0; y < grid.length; ++y) {
      if (grid[x][y] !== '@') continue;
      let ns = nbrs(x, y).map(([x, y]) => (gx = grid[x], gx !== undefined ? gx[y] : gx));
      if (ns.filter(n => n === '@').length < 4)
        removable.push([x, y]);
    }
  }
  removable.forEach(([x, y]) => {
    grid[x][y] = '.';
    removed++;
  });
} while(removable.length > 0);

grid.map(r => r.join('')).join('\n');
removed;
