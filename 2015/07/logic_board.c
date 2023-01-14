#include "logic_board.h"
#include <bits/stdint-uintn.h>

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

typedef struct node_stack {
    node *n;
    struct node_stack *prev;
} node_stack;

static void push(node_stack **top, node *n) {
    node_stack *new = (node_stack *)malloc(sizeof(node_stack));
    assert(new != NULL);

    new->n = n;
    new->prev = *top;
    *top = new;
}

static void pop(node_stack **top) {
    node_stack *curr = *top;
    node_stack *prev = curr->prev;
    free(curr);
    curr = NULL;
    *top = prev;
}

uint16_t logic_board_evaluate_stack_friendly(node *circuit) {
    node_stack *stack = NULL;
    push(&stack, circuit);

    while (stack != NULL) {
        node *c = stack->n;
        switch (c->type) {
        case WIRE_POINT: {
            if (c->type_value.wire.input->is_value_set) {
                c->value = c->type_value.wire.input->value;
                c->is_value_set = true;
                pop(&stack);
            } else {
                push(&stack, c->type_value.wire.input);
            }
            break;
        }
        case BINARY_OPERATION: {
            if (c->type_value.bin_op.lhs->is_value_set) {
                if (c->type_value.bin_op.rhs->is_value_set) {
                    uint16_t lhs = c->type_value.bin_op.lhs->value;
                    uint16_t rhs = c->type_value.bin_op.rhs->value;
                    uint16_t res;
                    switch (c->type_value.bin_op.type) {
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
                    c->value = res;
                    c->is_value_set = true;
                    pop(&stack);
                } else {
                    push(&stack, c->type_value.bin_op.rhs);
                }
            } else {
                push(&stack, c->type_value.bin_op.lhs);
            }
            break;
        }
        case UNARY_OPERATION: {
            if (c->type_value.un_op.input->is_value_set) {
                c->value = ~c->type_value.un_op.input->value;
                c->is_value_set = true;
                pop(&stack);
            } else {
                push(&stack, c->type_value.un_op.input);
            }
            break;
        }
        case INPUT_SOURCE: {
            // Root point
            // uint16_t value = c->value;
            c->is_value_set = true;
            pop(&stack);
            // c->value = value;
            // c->is_value_set = true;
            // pop(&stack);
            break;
        }
        default:
            assert(false && "programmer forgot to add branch");
        }
    }

    return circuit->value;
}
