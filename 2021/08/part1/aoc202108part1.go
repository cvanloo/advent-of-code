package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(input *os.File) (signals, outputs []string) {
	for scanner := bufio.NewScanner(input); scanner.Scan(); {
		line := scanner.Text()

		halfs := strings.Split(line, "|")
		s := strings.Fields(halfs[0])
		o := strings.Fields(halfs[1])

		signals = append(signals, s...)
		outputs = append(outputs, o...)
	}

	return
}

func contains(slice []int, el int) bool {
	for _, v := range slice {
		if v == el {
			return true
		}
	}
	return false
}

// 1, 4, 7, 8 are segments conisting of a unique number of chars.
var uniqueSegments = []int{2, 3, 4, 7}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, outputs := parseInput(file)

	result := 0

	for i := 0; i < len(outputs); i++ {
		strLen := len(outputs[i])

		if contains(uniqueSegments, strLen) {
			result++
		}
	}

	fmt.Printf("result: %d\n", result)
}
