input = document.getElementsByTagName('pre')[0].innerText;

input.split(',').map(range => range.split('-')).flatMap(([lower, upper]) => {
  let res = [];
  for (let n = parseInt(lower); n <= parseInt(upper); ++n) {
    let ns = n.toString();
    for (let i = 1; i <= ns.length / 2; ++i) {
      let nsp = ns.substring(0, i).repeat(ns.length / i);
      if (nsp === ns) {
        res.push(n);
        break;
      }
    }
  }
  return res;
}).reduce((a, b) => a + b);
