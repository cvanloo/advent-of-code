input = document.getElementsByTagName('pre')[0].innerText;

input
  .split('\n')
  .filter(l => l !== '')
  .map(bank => bank
    .split('')
    .reverse()
    .slice(2)
    .reduce(([r, l], bat) => bat >= r ? [bat, (r > l ? r : l)] : [r, l],
      [bank[bank.length - 2], bank[bank.length - 1]]))
  .map(([r, l]) => parseInt(r + l))
  .reduce((a, b) => a + b);
