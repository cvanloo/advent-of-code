package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func parseToBin(s string) uint {
	var b uint

	l := len(s) - 1

	for i, v := range s {
		vi, err := strconv.Atoi(string(v))
		panicIf(err)

		if vi == 0 {
			continue
		}
		b |= (1 << (l - i))
	}

	return b
}

func getMostCommonAt(bin []uint, i uint) uint {
	occurrence := map[uint]uint{}
	for _, b := range bin {
		bit := uint((b >> i) & 1)
		occurrence[bit]++
	}

	if occurrence[0] > occurrence[1] {
		return 0
	}

	return 1
}

//const lineLength = 5  // test.txt
const lineLength = 12 // input.txt

func main() {
	file, err := os.Open("input.txt")
	panicIf(err)

	scanner := bufio.NewScanner(file)
	binaries := []uint{}

	for scanner.Scan() {
		line := scanner.Text()

		b := parseToBin(line)
		binaries = append(binaries, b)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("encountered error while scanning input: %v\n", err)
	}

	var mostCommons uint

	for i := 0; i < lineLength; i++ {
		mc := getMostCommonAt(binaries, uint(i))
		if mc == 1 {
			mostCommons |= (1 << i)
		}
	}

	gammaRate := mostCommons

	fmt.Printf("gamma rate: 0b%b\n", gammaRate)
	fmt.Printf("gamma rate: %d\n", gammaRate)

	for i := 0; i < lineLength; i++ {
		mostCommons = (mostCommons ^ (1 << i))
	}

	epsilonRate := mostCommons

	fmt.Printf("epsilon rate: 0b%b\n", epsilonRate)
	fmt.Printf("epsilon rate: %d\n", epsilonRate)

	powerConsumption := gammaRate * epsilonRate

	fmt.Printf("power consumption: %d\n", powerConsumption)
}
