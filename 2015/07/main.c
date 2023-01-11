#include "logic_board.h"
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

void cleanup(void) {
  // clean up resources
}

int main(int argc, char **argv) {
  atexit(cleanup);
  printf("Hello, World!\n");
  exit(EXIT_SUCCESS);
}
