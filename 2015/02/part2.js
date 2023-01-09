var input = document.getElementsByTagName('pre')[0].innerHTML;
var lines = input.split('\n');
var ribbonLengths = lines.map(dimText => {
  let parts = dimText.split('x');
  let l = parseInt(parts[0]);
  let w = parseInt(parts[1]);
  let h = parseInt(parts[2]);

  let bow = l * w * h;

  let vals = [l, w, h].sort((a, b) => a - b);

  let sides = vals[0] * 2 + vals[1] * 2;

  return bow + sides;
});
ribbonLengths = ribbonLengths.filter(p => !isNaN(p));
console.log("ribbonLengths", ribbonLengths);
var totalRibbonLength = ribbonLengths.reduce((a, b) => a + BigInt(b), 0n);
console.log("Part 2:", totalRibbonLength);
