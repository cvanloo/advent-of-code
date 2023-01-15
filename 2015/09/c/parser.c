#include "lexer.h"
#include <assert.h>
#include <stdint.h>
#include <stdlib.h>

#define INITIAL_CAPACITY 500

typedef struct leg {
    struct node *target;
    uint32_t distance;
} leg;

typedef struct node {
    const char *name;
    uint32_t distance;
    struct leg **neighbours;
} node;

node *aoc_parser_parse(token *tokens);
node *aoc_parser_parse(token *tokens) {
    size_t current = 0;
    node *nodes = malloc(sizeof(node) * INITIAL_CAPACITY);

    while (true) {
        if (tokens[current].type == END_OF_FILE) {
            break;
        }

        token city_start = tokens[current++];
        ++current;
        token city_end = tokens[current++];
        ++current;
        uint32_t distance = atoi(tokens[current++].lexeme);

        node start_node;
        start_node.name = city_start.lexeme;

        node end_node;
        end_node.name = city_end.lexeme;

        leg leg_to;
        leg_to.distance = distance;
        leg_to.target = &end_node; // FIXME: shouldn't we store a reference to the pointer?
        start_node.neighbours = malloc(sizeof(leg));
        start_node.neighbours = leg_to;
    }

    return nodes;
}
