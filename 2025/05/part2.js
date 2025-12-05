input = document.getElementsByTagName('pre')[0].innerText;

[freshRanges, _] = input.split('\n\n').map(i => i.split('\n'));

addRange = (ranges, [l, u]) => {
    for (let [rl, ru] of ranges) {
        if (l >= rl && l <= ru) {
            if (u <= ru && u >= rl) {
                return ranges; // entire range already included
            } else {
                return addRange(ranges, [ru+1, u]);
            }
        } else if (u >= rl && u <= ru) {
            return addRange(ranges, [l, rl-1]);
        } else if (l < rl && u > ru) {
            addRange(ranges, [l, rl-1]);
            return addRange(ranges, [ru+1, u]);
        }
    }
    ranges.push([l, u]);
    return ranges;
}

freshRanges
    .map(r => r
        .split('-')
        .map(i => parseInt(i)))
    .reduce((ranges, range) => addRange(ranges, range), [])
    .map(([l, u]) => u-l+1)
    .reduce((a, b) => a + b);
