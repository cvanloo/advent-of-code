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

func parseInput(input string) ([]int, error) {
	input = strings.TrimRight(input, "\n")
	valuesStr := strings.Split(input, ",")

	// set cap to avoid resize on append
	values := make([]int, 0, len(valuesStr))

	for _, v := range valuesStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return values, err
		}
		values = append(values, num)
	}
	return values, nil
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	panicif(err)

	input := string(inputBytes)
	values, err := parseInput(input)
	panicif(err)

	/*count := map[int]int{}
	for _, v := range values {
		count[v]++
	}*/

	max := 0
	min := math.MaxInt

	for _, v := range values {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	differences := map[int]int{}

	for min <= max {
		for _, v := range values {
			differences[min] += int(math.Abs(float64(v - min)))
		}
		min++
	}

	last := math.MaxInt
	smallest := math.MaxInt
	for v, k := range differences {
		if k < last {
			last = k
			smallest = v
		}
	}

	fmt.Printf("optimal alignment is %d at a cost of %d\n", smallest, last)
}
