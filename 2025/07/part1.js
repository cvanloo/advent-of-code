input = document.getElementsByTagName('pre')[0].innerText;

rows = input.split('\n').filter(l => l !== '');

rows
  .map(r => r.split('').map((c, i) => c == 'S' || c == '^' ? i : -1).filter(n => n > -1)).filter(ns => ns.length > 0)
  .reduce(([nsplits, state], splitters) => {
    if (state.length === 0) state.push(splitters[0]);
    else {
      let newState = state.reduce((newBeams, beam) => {
        if (splitters.indexOf(beam) > -1) {
          nsplits++;
          newBeams.add(beam - 1);
          newBeams.add(beam + 1);
        } else {
          newBeams.add(beam);
        }
        return newBeams;
      }, new Set([]));
      state = Array.from(newState);
    }
    return [nsplits, state];
  }, [0, []])
