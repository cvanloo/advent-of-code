package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fd, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}
	defer fd.Close()

	var errors []error
	var calories []int = []int{0}
	currentElf := 0
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, 0)
			currentElf++
		} else {
			cal, err := strconv.Atoi(line)
			if err != nil {
				errors = append(errors, err)
				log.Printf("cannot parse %s as int: %v\n", line, err)
				continue
			}
			calories[currentElf] += cal
		}
	}

	scanErr := scanner.Err()
	if scanErr != nil {
		log.Fatalf("scanning error: %v\n", scanErr)
	}

	if ec := len(errors); ec > 0 {
		log.Printf("%d errors encountered while parsing input\n", ec)
	}

	sort.Slice(calories, func(i, j int) bool {
		// Sort in decending order
		return calories[i] > calories[j]
	})
	fmt.Printf("Part one solution: %d\n", calories[0])

	largestThree := calories[0] + calories[1] + calories[2]
	fmt.Printf("Part two solution: %d\n", largestThree)
}
