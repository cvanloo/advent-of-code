#include "parser.h"
#include "lexer.h"
#include "logic_board.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static inline lexer_token pc(parser *p) {
    return p->tokens[p->position];
}

static inline lexer_token pn(parser *p) {
    return p->tokens[p->position + 1];
}

static inline lexer_token pnn(parser *p) {
    return p->tokens[p->position + 2];
}

static inline void adv(parser *p) {
    ++p->position;
}

static node *add_node(parser *p, node n) {
    const size_t count = p->logic_size;
    const size_t capacity = p->logic_capacity;

    if (count == capacity) {
        const size_t new_capacity = capacity * 3 / 2 + 1;
        node *new_logic_board = realloc(p->logic_board, new_capacity);
        assert(new_logic_board != NULL);
        p->logic_board = new_logic_board;
        p->logic_capacity = new_capacity;
    }

    size_t index = p->logic_size;
    ++p->logic_size;
    p->logic_board[index] = n;
    return &p->logic_board[index];
}

static node *next_variable(parser *p) {
    assert(pc(p).type == VARIABLE);

    // Check if variable already exists, otherwise create it.
    const char *search = pc(p).lexeme;
    const size_t search_len = pc(p).lexeme_size;
    for (size_t i = 0; i < p->logic_size; ++i) {
        const node circuit = p->logic_board[i];
        if (circuit.type == WIRE_POINT) {
            const char *name = circuit.type_value.wire.name;
            const size_t name_len = circuit.type_value.wire.name_len;

            if (search_len == name_len &&
                strncmp(search, name, search_len) == 0) {
                return &p->logic_board[i];
            }
        }
    }

    node circuit;
    circuit.type = WIRE_POINT;
    circuit.value = 0;
    circuit.is_value_set = false;
    circuit.type_value.wire.name = search;
    circuit.type_value.wire.name_len = search_len;

    return add_node(p, circuit);
}

static void next_wire_point(parser *p, node *input) {
    assert(pc(p).type == CONNECTOR);
    assert(pn(p).type == VARIABLE);
    adv(p);

    node *circuit = next_variable(p);
    adv(p);

    assert(input != NULL);
    circuit->type_value.wire.input = input;
}

static void next_input_source(parser *p) {
    assert(pc(p).type == VALUE);

    uint16_t value = atoi(pc(p).lexeme);
    adv(p);

    node circuit;
    circuit.type = INPUT_SOURCE;
    circuit.value = value;
    circuit.is_value_set = true;
    next_wire_point(p, add_node(p, circuit));
}

static node *next_value_or_variable(parser *p) {
    node *n;
    switch (pc(p).type) {
    case VALUE: {
        uint16_t value = atoi(pc(p).lexeme);
        node circuit;
        circuit.type = INPUT_SOURCE;
        circuit.value = value;
        circuit.is_value_set = true;
        n = add_node(p, circuit);
        break;
    }
    case VARIABLE: {
        n = next_variable(p);
        break;
    }
    case OPERATION_AND:
    case OPERATION_LSHIFT:
    case OPERATION_RSHIFT:
    case OPERATION_OR:
    case OPERATION_NOT:
    case UNQUALIFIED:
    case END_OF_FILE:
    case INSTRUCTION_END:
    case CONNECTOR:
    default:
        assert(false && "not an operand");
    }
    return n;
}

static void next_binary_operation(parser *p) {
    node *wire_point_left = next_value_or_variable(p);
    adv(p);

    node operation;
    operation.type = BINARY_OPERATION;
    operation.value = 0;
    operation.is_value_set = false;
    switch (pc(p).type) {
    case OPERATION_AND: {
        operation.type_value.bin_op.type = AND;
        break;
    }
    case OPERATION_LSHIFT: {
        operation.type_value.bin_op.type = LSHIFT;
        break;
    }
    case OPERATION_RSHIFT: {
        operation.type_value.bin_op.type = RSHIFT;
        break;
    }
    case OPERATION_OR: {
        operation.type_value.bin_op.type = OR;
        break;
    }
    case UNQUALIFIED:
    case END_OF_FILE:
    case INSTRUCTION_END:
    case VALUE:
    case VARIABLE:
    case OPERATION_NOT:
    case CONNECTOR:
    default:
        assert(false && "invalid operation");
    }
    adv(p);

    node *wire_point_right = next_value_or_variable(p);
    adv(p);

    operation.type_value.bin_op.lhs = wire_point_left;
    operation.type_value.bin_op.rhs = wire_point_right;

    next_wire_point(p, add_node(p, operation));
}

static void next_unary_operation(parser *p) {
    assert(pc(p).type == OPERATION_NOT);

    node operation;
    operation.type = UNARY_OPERATION;
    operation.value = 0;
    operation.is_value_set = false;
    adv(p);

    assert(pc(p).type == VARIABLE);
    node *input = next_variable(p);
    adv(p);

    assert(input != NULL);
    operation.type_value.un_op.input = input;

    next_wire_point(p, add_node(p, operation));
}

static void next_variable_wire_point(parser *p) {
    assert(pc(p).type == VARIABLE);
    assert(pn(p).type == CONNECTOR);

    node *input_var = next_variable(p);
    adv(p);

    next_wire_point(p, input_var);
}

void parse_logic_board(parser *p) {
    while (true) {
        lexer_token c = pc(p);
        // printf("LINE: %lu\n", c.lineno);

        if (c.type == END_OF_FILE) {
            break;
        }

        switch (c.type) {
        case VALUE: {
            lexer_token n = pn(p);
            if (n.type == CONNECTOR) {
                next_input_source(p);
            } else {
                next_binary_operation(p);
            }
            break;
        }
        case VARIABLE: {
            lexer_token n = pn(p);
            switch (n.type) {
            case OPERATION_AND: {
            case OPERATION_OR:
            case OPERATION_LSHIFT:
            case OPERATION_RSHIFT:
                next_binary_operation(p);
                break;
            }
            case CONNECTOR: {
                next_variable_wire_point(p);
                break;
            }
            case UNQUALIFIED:
            case END_OF_FILE:
            case INSTRUCTION_END:
            case VALUE:
            case VARIABLE:
            case OPERATION_NOT:
            default:
                fprintf(stderr,
                        "Error on line %lu: VARIABLE must be followed by an "
                        "OPERATION, found %s\n",
                        c.lineno, token_type_str(n.type));
            }
            break;
        }
        case OPERATION_NOT: {
            next_unary_operation(p);
            break;
        }
        case UNQUALIFIED:
        case END_OF_FILE:
        case INSTRUCTION_END:
        case OPERATION_AND:
        case OPERATION_RSHIFT:
        case OPERATION_LSHIFT:
        case OPERATION_OR:
        case CONNECTOR:
        default:
            fprintf(
                stderr,
                "Error on line %lu: An instruction must start with a VALUE or "
                "VARIABLE, found %s\n",
                c.lineno, token_type_str(c.type));
        }

        assert(pc(p).type == INSTRUCTION_END);
        adv(p);
    }
}
