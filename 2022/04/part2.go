package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numOfOverlaps int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == '-'
		})

		// Theoretically, for single digit numbers, we could just compare the
		// ascii value directly.
		ints := make([]int, len(parts))
		for i, part := range parts {
			ints[i], _ = strconv.Atoi(part)
		}

		if (ints[0] >= ints[2] && ints[0] <= ints[3]) || (ints[2] >= ints[0] && ints[2] <= ints[1]) {
			// start comes after others start, but also before others end
			numOfOverlaps++
		}
	}

	if scanner.Err() != nil {
		log.Fatalf("failed to scan input: %v", scanner)
	}

	fmt.Printf("A total of %d ranges overlap.\n", numOfOverlaps)
}
