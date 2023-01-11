#include "logic_board.h"

uint16_t logic_board_evaluate(node *circuit) {
  node *current = circuit;

  switch (current->type) {
  case WIRE_POINT: {
    // TODO: Follow input.
    // Then set value.
    break;
  }
  case BINARY_OPERATION: {
    // TODO: Continue following both input signals.
    // Then evaluate op.
    break;
  }
  case UNARY_OPERATION: {
    // TODO: Continue following input singal.
    // Then evaluate op.
    break;
  }
  case INPUT_SOURCE: {
    // TODO: We're at the root point, go back downwards from here.
    // uint16_t val = current->value;
    break;
  }
  default:
    assert(false && "programmer forgot to add branch");
    return 0;
  }

  return circuit->value;
}
