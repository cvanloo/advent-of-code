package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}

// parseInput returns a map having a value as the key and the number of
// occurence of the value as the value.
func parseInput(input string) (map[int]int, error) {
	input = strings.TrimRight(input, "\n")
	valuesStr := strings.Split(input, ",")

	values := map[int]int{}

	for _, v := range valuesStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return values, err
		}
		values[num]++
	}
	return values, nil
}

// countFuel calculates the fuel required to align all values at a certain
// value.
func countFuel(values map[int]int, align int) int {
	result := 0
	for k, v := range values {
		distance := int(math.Abs(float64(k - align)))
		distance = (distance * (distance + 1)) / 2
		result += distance * v
	}
	return result
}

const max = 2000

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	panicif(err)

	input := string(inputBytes)
	values, err := parseInput(input)
	panicif(err)

	last := math.MaxInt
	smallestAlignment := 0
	for align := 0; align < max; align++ {
		fuel := countFuel(values, align)
		if fuel < last {
			smallestAlignment = align
			last = fuel
		}
	}

	fmt.Printf("optimal alignment is %d at a cost of %d\n", smallestAlignment,
		last)
}
