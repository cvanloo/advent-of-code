package main

import (
	"bufio"
	"fmt"
	"os"
)

//const lineLength = 5 // test.txt
const lineLength = 12 // input.txt

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func strToBin(s string) uint {
	var bin uint

	for i, v := range s {
		if v == '1' {
			bin |= (1 << (lineLength - 1 - i))
		}
	}

	return bin
}

func calculateGamma(bits [lineLength][2]uint) uint {
	var gamma uint
	for i := 0; i < lineLength; i++ {
		if bits[i][1] > bits[i][0] {
			gamma |= (1 << i)
		}
	}
	return gamma
}

func epsilonFromGamma(gamma uint) uint {
	//mask := uint(0b11111) // test.txt
	mask := uint(0b111111111111) // input.txt

	return gamma ^ mask
}

func main() {

	// First dimension represents the columns,
	// second dimension the number of 0s and
	// 1s in that column
	bits := [lineLength][2]uint{}
	file, err := os.Open("input.txt")
	panicIf(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		asi := strToBin(line)

		// iterate over each column
		for i := 0; i < lineLength; i++ {
			// if column is 0 increment
			// the number at [column][0],
			// otherwise at [column][1]
			bits[i][(asi>>i)&1]++
		}
	}

	gammaRate := calculateGamma(bits)
	epsilonRate := epsilonFromGamma(gammaRate)

	res := gammaRate * epsilonRate

	fmt.Printf("result: %d\n", res)
}
