package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func log(err error) {
	fmt.Fprintf(os.Stderr, "encountered error: %v\n", err)
}

var increased, last int = -1, 0

func handle(depth int) {
	if last == 0 || last < depth {
		increased++
	}
	last = depth
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log(err)
		} else {
			handle(value)
		}
	}

	check(scanner.Err())

	fmt.Printf("Depth increased %d times.\n", increased)
}
