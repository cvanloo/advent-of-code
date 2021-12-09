package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(input string) map[int8]int64 {
	fish := map[int8]int64{}
	input = strings.Trim(input, "\n")
	parts := strings.Split(input, ",")

	for _, v := range parts {
		i, err := strconv.Atoi(v)
		panicif(err)
		fish[int8(i)]++
	}

	return fish
}

//const simDays = 18
//const simDays = 80
const simDays = 256

func main() {
	bytes, err := os.ReadFile("input.txt")
	panicif(err)

	fish := parseInput(string(bytes))

	for i := 0; i < simDays; i++ {
		for j := -1; j < 8; j++ {
			fish[int8(j)] = fish[int8(j+1)]
		}
		fish[8] = fish[-1]  // baby fish
		fish[6] += fish[-1] // parent fish

		//fmt.Printf("after %d days: %v\n", i, fish)
	}

	var count int64
	for i := 0; i < 9; i++ {
		count += fish[int8(i)]
	}

	fmt.Printf("There are %d laternfish after %d days\n", count, simDays)
}
