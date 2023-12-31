#include "lexer.h"
#include "logic_board.h"
#include "parser.h"
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DEBUG 0

int main(int argc, char **argv) {
    printf("Hello, World!\n");

    if (argc < 3) {
        fprintf(stderr, "not enough arguments\n");
        return EXIT_FAILURE;
    }

    const char *filename = argv[1];
    FILE *fd = fopen(filename, "r");
    if (fd == NULL) {
        fprintf(stderr, "failed to open file\n");
        return EXIT_FAILURE;
    }

    fseek(fd, 0, SEEK_END);
    const size_t input_size = ftell(fd);
    fseek(fd, 0, SEEK_SET);

    char *input_text = malloc(input_size + 1);
    fread(input_text, input_size, 1, fd);
    input_text[input_size] = 0;
    fclose(fd);

    //  1. Lex
    lexer aoc_lexer = lexer_create(input_text, input_size, 100);
    lexer_tokenize(&aoc_lexer);
    printf("lexer parsed %lu tokens\n", aoc_lexer.token_count);

#if DEBUG == 1
    for (size_t i = 0; i < aoc_lexer.token_count; ++i) {
        lexer_token t = aoc_lexer.tokens[i];
        printf("%s(%s)\n", token_type_str(t.type), t.lexeme);
    }
#endif

    //  2. Parse
    parser aoc_parser = {
        .tokens = aoc_lexer.tokens,
        .token_count = aoc_lexer.token_count,
        .position = 0,
        .logic_board = malloc(sizeof(node) * 5000),
        .logic_capacity = 5000,
        .logic_size = 0,
    };
    parse_logic_board(&aoc_parser);
    printf("Logic Board Nodes: %lu\n", aoc_parser.logic_size);

#if DEBUG == 1
    logic_board_print(aoc_parser.logic_board, aoc_parser.logic_size);
#endif

    // Find circuit node wire_point with cli args provided name.
    const char *node_name = argv[2];
    node *point_of_interest = NULL;
    node *point_of_interest_b = NULL;

    for (size_t i = 0; i < aoc_parser.logic_size; ++i) {
        const node circuit = aoc_parser.logic_board[i];
        if (circuit.type == WIRE_POINT) {
            const char *name = circuit.type_value.wire.name;
            const size_t name_len = circuit.type_value.wire.name_len;

            if (point_of_interest == NULL && 1 == name_len &&
                strncmp(node_name, name, name_len) == 0) {
                point_of_interest = &aoc_parser.logic_board[i];
            }
            if (point_of_interest_b == NULL && 1 == name_len &&
                strncmp("b", name, name_len) == 0) {
                point_of_interest_b = &aoc_parser.logic_board[i];
            }
        }
    }
    assert(point_of_interest != NULL && "wire point not found");
    assert(point_of_interest_b != NULL && "wire point 'b' not found");

    // Do some re-wiring
    // Result from part 1: 16076
    //   make this be the input into wire b
    point_of_interest_b->value = 16076;
    point_of_interest_b->is_value_set = true;

    //  3. Emulate
    // uint16_t result = logic_board_evaluate(point_of_interest);
    uint16_t result = logic_board_evaluate_stack_friendly(point_of_interest);
    printf("Value of %s: %hu\n", node_name, result);

    // Cleanup
    lexer_destroy(&aoc_lexer);
    free(input_text);
    free(aoc_parser.logic_board);

    return EXIT_SUCCESS;
}
