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
CFILES := main.c lexer.c

all: $(CFILES)
	$(CC) -std=c99 -o a.out $^ $(CFLAGS)

release: $(CFILES)
	$(CC) -std=c99 -o a.out $^ $(RELEASEFLAGS)

.PHONY: clean
clean: a.out
	rm a.out

.PHONY: lsp
lsp:
	bear -- make
