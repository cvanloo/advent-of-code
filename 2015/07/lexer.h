#ifndef AOC_LEXER_H
#define AOC_LEXER_H

#include <ctype.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef enum lexer_token_type {
    UNQUALIFIED,
    END_OF_FILE,
    INSTRUCTION_END,
    VALUE,
    VARIABLE,
    OPERATION,
    CONNECTOR,
} lexer_token_type;

static inline const char *token_type_str(lexer_token_type type) {
    static const char *strings[] = {
        "UNQUALIFIED",
        "END_OF_FILE",
        "INSTRUCTION_END",
        "VALUE",
        "VARIABLE",
        "OPERATION",
        "CONNECTOR",
    };

    return strings[type];
}

typedef struct lexer_token {
    lexer_token_type type;
    const char *lexeme;
    uint64_t lineno;
} lexer_token;

typedef struct lexer {
    lexer_token *tokens;
    size_t token_count;
    size_t token_capacity;
    const char *input;
    const size_t input_size;
    char *position;
    size_t mark;
    uint64_t lineno;
    bool has_error;
} lexer;

lexer lexer_create(char *input, size_t input_size, size_t capacity);
void lexer_destroy(lexer *l);
void lexer_tokenize(lexer *l);

#endif // AOC_LEXER_H
