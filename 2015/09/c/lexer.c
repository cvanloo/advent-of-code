#include "lexer.h"
#include <ctype.h>

#define INITIAL_CAPACITY 500

static inline char current(aoc_lexer *l) {
    return l->position[l->mark];
}

static inline char next(aoc_lexer *l) {
    return l->position[l->mark + 1];
}

static inline bool is_char(const char c) {
    return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z');
}

static inline bool is_number(const char c) {
    return '0' <= c && c <= '9';
}

static inline void advance(aoc_lexer *l) {
    ++l->mark;
}

static inline void reverse(aoc_lexer *l) {
    --l->mark;
}

static inline void forward(aoc_lexer *l) {
    ++l->position;
}

static inline void fixate(aoc_lexer *l) {
    l->position += l->mark;
    l->mark = 0;
}

static token emit(aoc_lexer *l, token_type type) {
    const size_t lexeme_len = l->mark;
    char *lexeme = malloc(lexeme_len + 1);
    for (size_t i = 0; i < lexeme_len; ++i) {
        lexeme[i] = l->position[i];
    }
    lexeme[lexeme_len] = 0;
    fixate(l);

    token tok = {
        .type = type,
        .lexeme_len = lexeme_len,
        .lexeme = lexeme,
        .lineno = l->lineno,
    };

    return tok;
}

static void add_token(aoc_lexer *l, token tok) {
    const size_t count = l->token_count;
    const size_t capacity = l->token_capacity;

    if (count == capacity) {
        const size_t new_capacity = (capacity * 3) / 2 + 1;
        token *new_tokens = realloc(l->tokens, new_capacity * sizeof(token));
        assert(new_tokens != NULL);
        l->tokens = new_tokens;
    }

    l->tokens[l->token_count++] = tok;
}

static inline bool is_space_not_newline(const char c) {
    return c != '\n' && isspace(c);
}

static void skip_whitespace(aoc_lexer *l) {
    while (is_space_not_newline(current(l))) {
        forward(l);
    }
}

static token next_city(aoc_lexer *l) {
    assert(is_char(current(l)));

    while (is_char(current(l))) {
        advance(l);
    }

    return emit(l, CITY);
}

static token next_keyword_to(aoc_lexer *l) {
    assert(current(l) == 't' && next(l) == 'o');
    advance(l);
    advance(l);
    return emit(l, TO);
}

static token next_equals(aoc_lexer *l) {
    assert(current(l) == '=');
    advance(l);
    return emit(l, EQUAL);
}

static token next_number(aoc_lexer *l) {
    assert(is_number(current(l)));

    while (is_number(current(l))) {
        advance(l);
    }

    return emit(l, DISTANCE);
}

static token next_newline(aoc_lexer *l) {
    assert(current(l) == '\n');
    advance(l);
    token tok = emit(l, NEWLINE);
    ++l->lineno;

    return tok;
}

static token next_eof(aoc_lexer *l) {
    assert(current(l) == 0);
    token tok;
    tok.lineno = l->lineno;
    tok.type = END_OF_FILE;
    tok.lexeme = NULL;
    tok.lexeme_len = 0;
    return tok;
}

aoc_lexer aoc_lexer_create(char *input) {
    aoc_lexer l;
    l.position = input;
    l.mark = 0;
    l.tokens = (token *)malloc(sizeof(token) * INITIAL_CAPACITY);
    l.lineno = 1;
    l.token_count = 0;
    l.token_capacity = INITIAL_CAPACITY;

    return l;
}

void aoc_lexer_destroy(aoc_lexer l) {
    for (size_t i = 0; i < l.token_count; ++i) {
        free(l.tokens[i].lexeme);
    }
    free(l.tokens);
    l.tokens = NULL;
}

void aoc_lexer_tokenize(aoc_lexer *l) {
    while (true) {
        skip_whitespace(l);
        const char c = current(l);
        token tok;

        if (c == 0) {
            break;
        }

        if (c == '\n') {
            tok = next_newline(l);
        } else if (is_char(c)) {
            const char n = next(l);
            if (c == 't' && n == 'o') {
                tok = next_keyword_to(l);
            } else {
                tok = next_city(l);
            }
        } else if (c == '=') {
            tok = next_equals(l);
        } else if (is_number(c)) {
            tok = next_number(l);
        } else {
            assert(false && "invalid token");
            break;
        }

        add_token(l, tok);
    }

    add_token(l, next_eof(l));
}
