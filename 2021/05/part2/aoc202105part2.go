package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkif(err error) {
	if err != nil {
		panic(err)
	}
}

// sign returns `-1` if the number is negative, `0` if the number is `0` or `1`
// if the number is positive.
func sign(i int) int {
	switch {
	case i < 0:
		return -1
	case i > 0:
		return 1
	}
	return 0
}

//const gridSize = 10

const gridSize = 1000

func main() {
	file, err := os.Open("input.txt")
	checkif(err)
	defer file.Close() // ignore error of `Close`

	linesMatrix := [gridSize][gridSize]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readLine := scanner.Text()
		parts := strings.Split(readLine, " -> ")

		startPoints := strings.Split(parts[0], ",")
		sX, err := strconv.Atoi(startPoints[0])
		checkif(err)
		sY, err := strconv.Atoi(startPoints[1])
		checkif(err)

		endPoints := strings.Split(parts[1], ",")
		eX, err := strconv.Atoi(endPoints[0])
		checkif(err)
		eY, err := strconv.Atoi(endPoints[1])
		checkif(err)

		displacementX := eX - sX
		displacementY := eY - sY

		// find the length of the line
		top := int(math.Max(math.Abs(float64(displacementX)),
			math.Abs(float64(displacementY))))

		for i := 0; i <= top; i++ {
			// NOTE: By using the `sign` function we will continue to increment
			// when the displacement is positive, continue to decrement if it
			// is negative or do nothing if it is 0.
			x := sX + sign(displacementX)*i
			y := sY + sign(displacementY)*i
			linesMatrix[y][x]++
		}

		/*fmt.Printf("%v\n", readLine)
		for i := 0; i < len(linesMatrix); i++ {
			fmt.Printf("%v\n", linesMatrix[i])
		}*/
	}

	checkif(scanner.Err())

	overlaps := 0

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if linesMatrix[i][j] >= 2 {
				overlaps++
			}
		}
	}

	fmt.Printf("There are %d overlapping points\n", overlaps)
}
