package main

import (
	"fmt"
	"strconv"
	"time"
)

const input = "3113322113"
//const input = "211"

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func partition(digits []int) (groupedDigits [][]int) {
	lastDigit := digits[0]
	currentGroup := []int{}
	for _, digit := range digits {
		if digit != lastDigit {
			groupedDigits = append(groupedDigits, currentGroup)
			currentGroup = []int{}
		}
		currentGroup = append(currentGroup, digit)
		lastDigit = digit
	}
	groupedDigits = append(groupedDigits, currentGroup) // don't forget last group
	return groupedDigits
}

func next(groupedDigits [][]int) []int {
	nextDigits := make([]int, 0, 2*len(groupedDigits))
	for _, group := range groupedDigits {
		nextDigits = append(nextDigits, len(group))
		nextDigits = append(nextDigits, group[0])
	}
	return nextDigits
}

func main() {
	digits := make([]int, len(input))
	for i, digit := range input {
		digits[i] = must(strconv.Atoi(string(digit)))
	}
	fmt.Printf("%+v\n", digits)

	t := time.Now()
	for i := 0; i < 50; i++ {
		groupedDigits := partition(digits)
		digits = next(groupedDigits)
	}
	dur := time.Now().Sub(t)
	fmt.Printf("%+v (took %s)\n", len(digits), dur)
	// part 2 (50x): 4666278, (took 1.551496551s)
}
