input = document.getElementsByTagName('pre')[0].innerText;

[freshRanges, ingredients] = input.split('\n\n').map(i => i.split('\n'))

ingredients
    .filter(l => l !== '')
    .map(ing => parseInt(ing))
    .map(ing => freshRanges
        .map(r => r
            .split('-')
            .map(i => parseInt(i)))
        .map(([l, u]) => x => l <= x && x <= u)
        .reduce((cond, f) => x => cond(x) || f(x), x => false)(ing))
    .filter(c => c)
    .length;
