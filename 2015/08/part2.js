var input = document.getElementsByTagName('pre')[0].innerHTML;
var lines = input.split('\n').filter(line => line != "");

var counts = lines.map(line => {
    const actual = line.length;

    const chars = line.split('');
    let encoded = 2;
    for (let i = 0; i < chars.length; ++i) {
        const c = chars[i];
        if (c == '"') {
            encoded += 2;
        } else if (c == '\\') {
            encoded += 2;
        } else {
            ++encoded;
        }
    }

    console.log(actual, encoded);
    return [actual, encoded];
});

var res = counts.reduce((acc, t) => {
    const [actual, encoded] = t;
    return acc += (encoded - actual);
}, 0);

console.log(res);
