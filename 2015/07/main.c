#include "logic_board.h"
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
  printf("Hello, World!\n");

  if (argc < 2) {
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

  // TODO:
  //  1. Lex
  //  2. Parse
  //  3. Emulate

  return EXIT_SUCCESS;
}
