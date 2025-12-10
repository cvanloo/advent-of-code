input = document.getElementsByTagName('pre')[0].innerText;

rows = input.split('\n').filter(l => l !== '');

rows
  .map(r => r.split('').map((c, i) => c == 'S' || c == '^' ? i : -1).filter(n => n > -1)).filter(ns => ns.length > 0)
  .reduce((state, splitters) => {
    if (state.length === 0) {
      state[splitters[0]] = 1;
    } else {
      newState = [];
      for (let i = 0; i < state.length; ++i) {
        if (state[i] > 0) {
          if (splitters.indexOf(i) > -1) {
            newState[i - 1] = (newState[i - 1] ?? 0) + state[i];
            newState[i + 1] = (newState[i + 1] ?? 0) + state[i];
          } else {
            newState[i] = (newState[i] ?? 0) + state[i];
          }
        }
      }
      state = newState;
    }
    return state;
  }, []).reduce((a, b) => a + b);
