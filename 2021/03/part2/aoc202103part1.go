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
			bin |= (1 << (lineLength - 1 - i)) // start at left-most bit
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
	//mask := uint(0b11111)
	mask := uint(0b111111111111)

	return gamma ^ mask
}

func countBitsAt(bitIndex int, values []uint) map[uint]uint {
	count := map[uint]uint{}
	for _, v := range values {
		count[uint((v>>bitIndex)&1)]++
	}
	return count
}

func searchValue(values []uint, invert bool) uint {
	keep := values

	for i := lineLength - 1; i >= 0; i-- {
		if len(keep) == 1 {
			return keep[0]
		}
		bitCount := countBitsAt(i, keep)

		if (!invert && bitCount[1] >= bitCount[0]) || (invert && bitCount[1] < bitCount[0]) {
			// keep all with a 1 at position i
			lk := []uint{}
			for _, v := range keep {
				// check if bit is set
				if v&(1<<i) != 0 {
					lk = append(lk, v)
				}
			}
			keep = lk
		} else {
			// keep all with a 0 at position i
			lk := []uint{}
			for _, v := range keep {
				// check if bit is not set
				if v&(1<<i) == 0 {
					lk = append(lk, v)
				}
			}
			keep = lk
		}
	}

	return keep[0]
}

func main() {
	// First dimension represents the columns,
	// second dimension the number of 0s and
	// 1s in that column
	bits := [lineLength][2]uint{}
	values := []uint{}

	file, err := os.Open("input.txt")
	panicIf(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		asi := strToBin(line)
		values = append(values, asi)

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
	fmt.Printf("result for part 1: %d\n", res)

	fmt.Printf("bits: %v\n", bits)

	oxygen := searchValue(values, false)
	co2 := searchValue(values, true)
	fmt.Printf("found: %d\n", oxygen*co2)
}
