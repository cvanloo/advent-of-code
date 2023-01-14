#include "logic_board.h"

uint16_t logic_board_evaluate(node *circuit) {
    switch (circuit->type) {
    case WIRE_POINT: {
        circuit->value = logic_board_evaluate(circuit->type_value.wire.input);
        break;
    }
    case BINARY_OPERATION: {
        uint16_t lhs = logic_board_evaluate(circuit->type_value.bin_op.lhs);
        uint16_t rhs = logic_board_evaluate(circuit->type_value.bin_op.rhs);
        uint16_t res;
        switch (circuit->type_value.bin_op.type) {
        case AND: {
            res = lhs & rhs;
            break;
        }
        case OR: {
            res = lhs | rhs;
            break;
        }
        case LSHIFT: {
            res = lhs << rhs;
            break;
        }
        case RSHIFT: {
            res = lhs >> rhs;
            break;
        }
        default:
            assert(false && "programmer forgot to add branch");
        }
        circuit->value = res;
        break;
    }
    case UNARY_OPERATION: {
        uint16_t in = logic_board_evaluate(circuit->type_value.un_op.input);
        circuit->value = ~in;
        break;
    }
    case INPUT_SOURCE: {
        // Root point
        assert(circuit->value != 0);
        break;
    }
    default:
        assert(false && "programmer forgot to add branch");
    }

    return circuit->value;
}
