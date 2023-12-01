// gcc -Wall -Wextra -Wswitch-enum -g main.c
// valgrind --leak-check=full --show-leak-kinds=all --track-origins=yes ./a.out input.txt
#include <stdlib.h>
#include <stdio.h>
#include <assert.h>
#include <errno.h>
#include <string.h>
#include <stdio.h>
#include <ctype.h>

#if 1
#define PART_2
#endif

#define DEFER(code) do { \
	exit_code = (code); \
	goto defer; \
} while (0);

struct Slice {
	const char *start;
	int length;
};

struct Arena {
	size_t capacity;
	size_t length;
	char bytes[];
};

struct Arena *
arena_new(size_t capacity);

void *
arena_alloc(struct Arena *a, size_t size);

int
read_file(const char *path, char **out_bytes, size_t *out_len);

void
split_lines(struct Arena *a, const char *text, size_t len, struct Slice **out_lines, int *out_line_count);

int
parse_int(struct Slice text, int *digit);

int numbers[1024];
int number_count = 0;

int
main(int argc, char **argv) {
	assert(argc >= 2 && "Provide a file name");

	const char *input_file_name = argv[1];
	char *content = NULL;
	size_t content_length = 0;
	assert(read_file(input_file_name, &content, &content_length));

	struct Arena *arena_lines = arena_new(1000 * sizeof(struct Slice));
	struct Slice *lines;
	int line_count = 0;
	split_lines(arena_lines, content, content_length, &lines, &line_count);

#if 0
	for (int i = 0; i < line_count; ++i) {
		struct Slice line = lines[i];
		printf("%d: %.*s\n", i, line.length, line.start);
	}
#endif

	struct Slice line, substr;
	int digit_in_text_len, first_digit, last_digit;
	for (int line_idx = 0; line_idx < line_count; ++line_idx) {
		first_digit = last_digit = -1;
		line = lines[line_idx];
		for (int char_idx = 0; char_idx < line.length; ++char_idx) { // find first number
			substr.start = line.start + char_idx;
			substr.length = line.length - char_idx;
			digit_in_text_len = parse_int(substr, &first_digit);
			if (digit_in_text_len > 0) {
				break;
			}
		}
		for (int char_idx = line.length-1; char_idx >= 0; --char_idx) { // find last number
			substr.start = line.start + char_idx;
			substr.length = line.length - char_idx;
			digit_in_text_len = parse_int(substr, &last_digit);
			if (digit_in_text_len > 0) {
				break;
			}
		}
		assert(first_digit > -1);
		if (last_digit == -1) last_digit = first_digit;
		numbers[number_count++] = first_digit * 10 + last_digit;
		//printf("%d: %d\n", line_idx, numbers[number_count-1]);
	}

	int result = 0;
	int *number = &numbers[0];
	for (int *end = &number[number_count]; number < end; ++number) {
		//printf("%d, ", *number);
		result += *number;
	}
	printf("Result: %d\n", result);
	// Part 1: 55386
	// Part 2: 54824

	free(content);
	free(arena_lines);
	return 0;
}

int
parse_int(struct Slice text, int *digit) {
	static const struct {
		const char *name;
		int number;
	} number_map[] = {
		// '<,'>Tab/'\d'/l1r0
		{"1",     1},
		{"2",     2},
		{"3",     3},
		{"4",     4},
		{"5",     5},
		{"6",     6},
		{"7",     7},
		{"8",     8},
		{"9",     9},
#ifdef PART_2
		{"one",   1},
		{"two",   2},
		{"three", 3},
		{"four",  4},
		{"five",  5},
		{"six",   6},
		{"seven", 7},
		{"eight", 8},
		{"nine",  9},
#endif
	};
	static const int number_map_count = sizeof(number_map) / sizeof(number_map[0]);

	for (int i = 0; i < number_map_count; ++i) {
		const char *name = number_map[i].name;
		const size_t name_len = strlen(name);
		if (text.length >= (int) name_len) {
			if (strncmp(text.start, name, name_len) == 0) {
				if (digit) *digit = number_map[i].number;
				return name_len;
			}
		}
	}
	return 0;
}

int
read_file(const char *path, char **bytes, size_t *len) {
	int exit_code = 0;

	FILE *fd = fopen(path, "rb");
	if (!fd) {
		printf("%s\n", strerror(errno));
		DEFER(0);
	}

	fseek(fd, 0, SEEK_END);
	long size = ftell(fd);
	if (size <= 0) DEFER(1);
	fseek(fd, 0, SEEK_SET);

	char *buf = (char *) malloc(size);
	if (!buf) DEFER(0);

	size_t bs_read = fread(buf, 1, size, fd);
	if (bytes) *bytes = buf;
	if (len) *len = bs_read;
	DEFER(1);

defer:
	fclose(fd);
	return exit_code;
}

void
split_lines(struct Arena *a, const char *text, size_t len, struct Slice **out_lines, int *out_line_count) {
	struct Slice *first = NULL;
	size_t last_line_end = 0;
	int slice_count = 0;
	for (size_t i = 0; i < len; ++i) {
		if (text[i] == '\n') {
			int length = i - last_line_end;
			if (length > 0) {
				struct Slice *slice = arena_alloc(a, sizeof(struct Slice));
				if (!first) first = slice;
				slice->start = text + last_line_end;
				slice->length = length;
				++slice_count;
				// gdb: p *slice->start@slice->length
			}
			last_line_end = i + 1;
		}
	}
	if (out_lines) *out_lines = first;
	if (out_line_count) *out_line_count = slice_count;
}

struct Arena *
arena_new(size_t capacity) {
	struct Arena *arena;
	arena = malloc(sizeof(*arena) + sizeof(char) * capacity);
	arena->capacity = capacity;
	arena->length = 0;
	return arena;
}

void *
arena_alloc(struct Arena *a, size_t size) {
	assert(a->length + size <= a->capacity);
	void *ptr = (void *) a->bytes + a->length;
	a->length += size;
	return ptr;
}
