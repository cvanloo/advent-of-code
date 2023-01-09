var input = document.getElementsByTagName('pre')[0].innerHTML;
var lines = input.split('\n');
var paperSizes = lines.map(dimText => {
  let parts = dimText.split('x');
  let l = parseInt(parts[0]);
  let w = parseInt(parts[1]);
  let h = parseInt(parts[2]);

  let s = (2 * l * w) + (2 * w * h) + (2 * h * l);
  
  let vals = [l, w, h].sort((a, b) => a - b);
  
  let extra = vals[0] * vals[1];
  
  return extra + s;
});
paperSizes = paperSizes.filter(p => !isNaN(p));
console.log(paperSizes);
var totalPaperNeded = paperSizes.reduce((a, b) => a + BigInt(b), 0n);
console.log("Part 1:", totalPaperNeded);
