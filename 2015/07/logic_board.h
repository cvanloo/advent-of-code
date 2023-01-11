#ifndef AOC_LOGIC_BOARD_H
#define AOC_LOGIC_BOARD_H

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef enum {
  AND,
  OR,
  RSHIFT,
  LSHIFT,
  NOT,
} operation_type;

typedef enum {
  WIRE_POINT,
  BINARY_OPERATION,
  UNARY_OPERATION,
  INPUT_SOURCE,
} node_type;

typedef struct binary_operation {
  operation_type type;
  struct node *lhs;
  struct node *rhs;
} binary_operation;

typedef struct unary_operation {
  operation_type type;
  struct node *input;
} unary_operation;

typedef struct wire_point {
  const char *name;
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

uint16_t logic_board_evaluate(node *circuit);

#endif // AOC_LOGIC_BOARD_H
