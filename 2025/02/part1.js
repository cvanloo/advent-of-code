input = document.getElementsByTagName('pre')[0].innerText;

input
  .split(',')
  .map(range => range.split('-'))
  .flatMap(([lower, upper]) => {
    let res = [];
    for (let n = parseInt(lower); n <= parseInt(upper); ++n) {
      let ns = n.toString();
      if (ns.slice(ns.length / 2) === ns.slice(0, ns.length / 2))
        res.push(n);
    }
    return res;
  }).reduce((a, b) => a + b);
