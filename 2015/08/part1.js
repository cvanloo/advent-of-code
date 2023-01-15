var input = document.getElementsByTagName('pre')[0].innerHTML;

var counts = input.split('\n').map(line => {
    const total = line.length;

    let actual = 0;
    const chars = line.split('');
    for (let i = 0; i < chars.length; ++i) {
        const c = chars[i];
        const n = chars[i + 1];

        if (c == '"') {
            // Don't count it.
        } else if (c == '\\') {
            if (n == '\\') {
                ++actual;
                ++i;
            } else if (n == '"') {
                ++actual;
                ++i;
            } else if (n == 'x') {
                ++actual;
                i += 3;
            }
        } else {
            ++actual;
        }
    }

    console.log(total, actual);
    return [total, actual];
});

var res = counts.reduce((acc, t) => {
    const [total, actual] = t;
    return acc += (total - actual);
}, 0);

console.log(res);
