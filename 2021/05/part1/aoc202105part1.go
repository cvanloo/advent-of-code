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

//const gridSize = 10
const gridSize = 1000

func main() {
	file, err := os.Open("input.txt")
	checkif(err)
	defer file.Close()

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

		maxX := int(math.Max(float64(eX), float64(sX)))
		minX := int(math.Min(float64(eX), float64(sX)))

		maxY := int(math.Max(float64(eY), float64(sY)))
		minY := int(math.Min(float64(eY), float64(sY)))

		// TODO: y = mx+b ?

		if maxY-minY == 0 {
			// horizontal line
			for minX <= maxX {
				linesMatrix[sY][minX]++
				minX++
			}
		} else if maxX-minX == 0 {
			// vertical line
			for minY <= maxY {
				linesMatrix[minY][sX]++
				minY++
			}
		}
	}

	checkif(scanner.Err())

	for i := 0; i < len(linesMatrix); i++ {
		fmt.Printf("%v\n", linesMatrix[i])
	}

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
