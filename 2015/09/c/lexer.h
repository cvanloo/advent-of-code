#ifndef AOC_LEXER_H
#define AOC_LEXER_H

#include <assert.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef enum token_type {
    END_OF_FILE,
    NEWLINE,
    CITY,
    TO,
    EQUAL,
    DISTANCE,
} token_type;

static inline const char *token_type_str(token_type type) {
    static const char *strings[] = {
        "END_OF_FILE", "NEWLINE", "CITY", "TO", "EQUAL", "DISTANCE",
    };

    return strings[type];
}

typedef struct token {
    token_type type;
    char *lexeme;
    size_t lexeme_len;
    size_t lineno;
} token;

typedef struct aoc_lexer {
    char *position;
    size_t mark;
    size_t lineno;
    token *tokens;
    size_t token_count;
    size_t token_capacity;
} aoc_lexer;

aoc_lexer aoc_lexer_create(char *input);
void aoc_lexer_tokenize(aoc_lexer *l);
void aoc_lexer_destroy(aoc_lexer l);

#endif // AOC_LEXER_H
