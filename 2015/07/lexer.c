#include "lexer.h"
#include <assert.h>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static lexer_token emit(lexer *l, lexer_token_type type);
static void add_token(lexer *l, lexer_token token);

static inline char lc(lexer *l) {
    return l->position[l->mark];
}

static inline char ln(lexer *l) {
    return l->position[l->mark + 1];
}

static inline char lnn(lexer *l) {
    return l->position[l->mark + 2];
}

static inline void adv(lexer *l) {
    ++l->mark;
}

static inline void fwd(lexer *l) {
    ++l->position;
}

static inline void rev(lexer *l) {
    --l->mark;
}

static inline void fix(lexer *l) {
    l->position += l->mark;
    l->mark = 0;
}

static inline bool is_space(const char c) {
    return c != '\n' && isspace(c);
}

static void skip_whitespace(lexer *l) {
    while (is_space(l->position[0])) {
        fwd(l);
    }
    fix(l);
}

static inline bool is_num_char(const char c) {
    return '0' <= c && c <= '9';
}

static inline bool is_op_char(const char c) {
    return 'A' <= c && c <= 'Z';
}

static inline bool is_var_char(const char c) {
    return 'a' <= c && c <= 'z';
}

static lexer_token next_number(lexer *l) {
    assert(is_num_char(lc(l)));

    while (is_num_char(lc(l))) {
        adv(l);
    }

    lexer_token token = emit(l, VALUE);
    fix(l);
    return token;
}

static lexer_token next_variable(lexer *l) {
    assert(is_var_char(lc(l)));

    while (is_var_char(lc(l))) {
        adv(l);
    }

    lexer_token token = emit(l, VARIABLE);
    fix(l);
    return token;
}

static lexer_token next_operation(lexer *l) {
    static const struct {
        const char *keyword;
        lexer_token_type token_type;
    } keywords[] = {
        {   "AND",    OPERATION_AND},
        {    "OR",     OPERATION_OR},
        {"LSHIFT", OPERATION_LSHIFT},
        {"RSHIFT", OPERATION_RSHIFT},
        {   "NOT",    OPERATION_NOT},
    };

    assert(is_op_char(lc(l)));

    while (is_op_char(lc(l))) {
        adv(l);
    }

    lexer_token token = emit(l, UNQUALIFIED);
    fix(l);

    const size_t keywords_size = sizeof(keywords) / sizeof(keywords[0]);
    const char *word = token.lexeme;
    const size_t word_size = token.lexeme_size;

    for (size_t i = 0; i < keywords_size; ++i) {
        const char *keyword = keywords[i].keyword;
        const size_t keyword_size = strlen(keyword);

        if (keyword_size == word_size &&
            strncmp(word, keyword, word_size) == 0) {
            token.type = keywords[i].token_type;
            break;
        }
    }

    return token;
}

static lexer_token next_connector(lexer *l) {
    assert(lc(l) == '-' && ln(l) == '>');

    adv(l);
    adv(l);

    lexer_token token = emit(l, CONNECTOR);
    fix(l);
    return token;
}

static lexer_token next_newline(lexer *l) {
    assert(lc(l) == '\n');

    // Forward instead of advance, so that emit won't allocate any memory for
    // the lexeme string.
    fwd(l);

    lexer_token token = emit(l, INSTRUCTION_END);
    fix(l);
    token.lexeme = malloc(3);
    token.lexeme[0] = '\\';
    token.lexeme[1] = 'n';
    token.lexeme[2] = 0;
    token.lexeme_size = 2;
    ++l->lineno;
    return token;
}

static lexer_token next_eof(lexer *l) {
    assert(lc(l) == 0);
    fwd(l);
    lexer_token token = emit(l, END_OF_FILE);
    token.lexeme = malloc(4);
    token.lexeme[0] = 'E';
    token.lexeme[1] = 'O';
    token.lexeme[2] = 'F';
    token.lexeme[3] = 0;
    token.lexeme_size = 3;
    return token;
}

lexer lexer_create(char *input, size_t input_size, size_t capacity) {
    lexer_token *tokens = malloc(sizeof(lexer_token) * capacity);
    assert(tokens != NULL);
    lexer l = {
        .position = input,
        .input_size = input_size,
        .mark = 0,
        .lineno = 1,
        .tokens = tokens,
        .token_capacity = capacity,
        .token_count = 0,
        .has_error = false,
    };
    return l;
}

void lexer_destroy(lexer *l) {
    const lexer_token *tokens = l->tokens;
    for (size_t i = 0; i < l->token_count; ++i) {
        free(tokens[i].lexeme);
    }
    free(l->tokens);
    l->tokens = NULL;
    l->token_count = 0;
}

void lexer_tokenize(lexer *l) {
    while (true) {
        skip_whitespace(l);
        const char c = lc(l);
        lexer_token token = {.type = UNQUALIFIED};

        if (c == 0) {
            break;
        }

        if (is_num_char(c)) {
            token = next_number(l);
        } else if (is_var_char(c)) {
            token = next_variable(l);
        } else if (is_op_char(c)) {
            token = next_operation(l);
        } else {
            if (c == '\n') {
                token = next_newline(l);
            } else if (c == '-') {
                const char n = ln(l);
                if (n == '>') {
                    token = next_connector(l);
                }
            }
        }

        add_token(l, token);
    }

    lexer_token token = next_eof(l);
    add_token(l, token);
}

static lexer_token emit(lexer *l, lexer_token_type type) {
    size_t lexeme_size = l->mark;
    char *lexeme = malloc(lexeme_size + 1);
    assert(lexeme != NULL);

    for (size_t i = 0; i < lexeme_size; ++i) {
        lexeme[i] = l->position[i];
    }

    lexeme[lexeme_size] = 0;

    lexer_token token = {
        .type = type,
        .lexeme = lexeme,
        .lexeme_size = lexeme_size,
        .lineno = l->lineno,
    };

    return token;
}

static void add_token(lexer *l, lexer_token token) {
    const size_t size = l->token_count;
    const size_t capacity = l->token_capacity;

    if (size == capacity) {
        const size_t new_capacity = (capacity * 3) / 2 + 1;
        l->token_capacity = new_capacity;
        lexer_token *new_tokens =
            realloc(l->tokens, new_capacity * sizeof(lexer_token));
        assert(new_tokens != NULL);
        l->tokens = new_tokens;
    }

    l->tokens[l->token_count++] = token;
}
