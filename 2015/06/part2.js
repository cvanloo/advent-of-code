var input = document.getElementsByTagName('pre')[0].innerText.split('\n');
var lights = new Array(1000).fill(0).map(() => new Array(1000).fill(0));

var instructions = input.map(instruction => {
  let parts = instruction.split(' ');
  let action = null;
  let start = {};
  let end = {};

  if (parts[0] === 'turn') {
    action = parts[1];
    let s = parts[2].split(',');
    start = {
      x: parseInt(s[0]),
      y: parseInt(s[1])
    }
    let e = parts[4].split(',');
    end = {
      x: parseInt(e[0]),
      y: parseInt(e[1])
    }
  } else if (parts[0] === 'toggle') {
    action = 'toggle';
    let s = parts[1].split(',');
    start = {
      x: parseInt(s[0]),
      y: parseInt(s[1])
    }
    let e = parts[3].split(',');
    end = {
      x: parseInt(e[0]),
      y: parseInt(e[1])
    }
  } else {
    console.error('invalid instruction', parts[0]);
    return null;
  }

  return {
    action: action,
    start: start,
    end: end
  };
}).filter(i => i !== null);

instructions.forEach(i => {
  for (let x = i.start.x; x <= i.end.x; ++x) {
    for (let y = i.start.y; y <= i.end.y; ++y) {
      switch (i.action) {
        case 'on':
          lights[x][y] += 1;
          break;
        case 'off':
          if (lights[x][y] > 0) {
            lights[x][y] -= 1;
          }
          break;
        case 'toggle':
          lights[x][y] += 2;
          break;
        default:
          console.error('invalid action', action);
          break;
      }
    }
  }
});

lights.flat().reduce((a, l) => a + l); // 14110788
