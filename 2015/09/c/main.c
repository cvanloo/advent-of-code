#include "lexer.h"
#include <assert.h>
#include <bits/types/FILE.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    if (argc < 2) {
        fprintf(stderr, "missing arguments\n");
        return EXIT_FAILURE;
    }

    const char *filename = argv[1];
    FILE *fd = fopen(filename, "r");
    if (fd == NULL) {
        fprintf(stderr, "failed to open file\n");
        return EXIT_FAILURE;
    }
    fseek(fd, 0, SEEK_END);
    size_t file_size = ftell(fd);
    fseek(fd, 0, SEEK_SET);

    char *input = malloc(file_size + 1);
    size_t ret = fread(input, file_size, 1, fd);
    if (ret != 1) {
        fprintf(stderr, "failed to read file\n");
        return EXIT_FAILURE;
    }
    input[file_size] = 0;

    aoc_lexer lexer = aoc_lexer_create(input);
    aoc_lexer_tokenize(&lexer);
    printf("Lexer tokenized %lu tokens\n", lexer.token_count);

    for (size_t i = 0; i < lexer.token_count; ++i) {
        token tok = lexer.tokens[i];
        printf("%s, %lu, %s\n", token_type_str(tok.type), tok.lineno, tok.lexeme);
    }

    aoc_lexer_destroy(lexer);
    free(input);
    return EXIT_SUCCESS;
}
