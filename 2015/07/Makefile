CC := gcc
CFLAGS := -I. -Wall -Wextra -pedantic -Wshadow -Wpointer-arith -Wcast-qual\
		  -Wstrict-prototypes -Wmissing-prototypes -Wno-missing-braces\
		  -Wno-missing-field-initializers -Wformat=2 -Wswitch-default\
		  -Wswitch-enum -Wcast-align -Wbad-function-cast -Wstrict-overflow=5\
		  -Winline -Wunreachable-code -Wlogical-op -Wfloat-equal\
		  -Wstrict-aliasing=2 -Wredundant-decls -Wold-style-definition\
		  -Werror -ggdb3 -O0 -fno-omit-frame-pointer -ffloat-store\
		  -fno-common -fstrict-aliasing -lm
RELEASEFLAGS := -I. -pedantic -O3 -DNDEBUG -lm
CFILES := main.c logic_board.c lexer.c parser.c

all: $(CFILES)
	$(CC) -std=c99 -o part1 $^ $(CFLAGS)

release: $(CFILES)
	$(CC) -std=c99 -o part1 $^ $(RELEASEFLAGS)

.PHONY: clean
clean:
	rm part1

.PHONY: debug
debug: all test.txt
	gdb -tui --args part1 test.txt d

.PHONY: lsp
lsp:
	bear -- make
