#ifndef AOC_LOGIC_BOARD_H
#define AOC_LOGIC_BOARD_H

#include <assert.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

typedef enum {
    AND,
    OR,
    RSHIFT,
    LSHIFT,
} binary_operation_type;

typedef enum {
    WIRE_POINT,
    BINARY_OPERATION,
    UNARY_OPERATION,
    INPUT_SOURCE,
} node_type;

typedef struct binary_operation {
    binary_operation_type type;
    struct node *lhs;
    struct node *rhs;
} binary_operation;

typedef struct unary_operation {
    // Type can only be NOT
    struct node *input;
} unary_operation;

typedef struct wire_point {
    const char *name;
    size_t name_len;
    struct node *input;
} wire_point;

typedef struct node {
    node_type type;
    union {
        wire_point wire;
        binary_operation bin_op;
        unary_operation un_op;
    } type_value;
    uint16_t value;
} node;

static inline void logic_board_print(const node *board, const size_t board_len) {
    for (size_t i = 0; i < board_len; ++i) {
        node n = board[i];

        switch (n.type) {
        case INPUT_SOURCE: {
            printf("INPUT_SOURCE(%hu)", n.value);
            break;
        }
        case WIRE_POINT: {
            printf("WIRE_POINT(%s)", n.type_value.wire.name);
            break;
        }
        case BINARY_OPERATION: {
            switch (n.type_value.bin_op.type) {
            case AND: {
                printf("BINARY_OPERATION(AND)");
                break;
            }
            case OR: {
                printf("BINARY_OPERATION(OR)");
                break;
            }
            case RSHIFT: {
                printf("BINARY_OPERATION(RSHIFT)");
                break;
            }
            case LSHIFT: {
                printf("BINARY_OPERATION(LSHIFT)");
                break;
            }
            default:
                assert(false && "unreachable");
            }
            break;
        }
        case UNARY_OPERATION: {
            printf("UNARY_OPERATION(NOT)");
            break;
        }
        default:
            assert(false && "unreachable");
        }

        printf("\n");
    }
}

uint16_t logic_board_evaluate(node *circuit);

#endif // AOC_LOGIC_BOARD_H
