shiftBatteries = ([r, l, ...w], bnew) => {
  if (l === undefined) return bnew >= r ? [bnew] : [r];
  return bnew >= r ? [bnew, ...shiftBatteries([l, ...w], r)] : [r, l, ...w];
}

solve = input => input
  .split('\n')
  .filter(l => l !== '')
  .map(bank => bank
    .split('')
    .toReversed()
    .slice(12)
    .reduce((bs, bat) => shiftBatteries(bs, bat),
      bank.split('').slice(-12)))
  .map(ds => parseInt(ds.join('')))
  .reduce((a, b) => a + b);

input = document.getElementsByTagName('pre')[0].innerText;

solve(input);
