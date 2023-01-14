#ifndef AOC_PARSER_H
#define AOC_PARSER_H

#include "lexer.h"
#include "logic_board.h"

typedef struct parser {
    lexer_token *tokens;
    size_t token_count;
    size_t position;
    node *logic_board;
    size_t logic_size;
    size_t logic_capacity;
} parser;

void parse_logic_board(parser *p);

#endif // AOC_PARSER_H
