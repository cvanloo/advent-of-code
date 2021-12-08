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

		// find the length of the line
		top := int(math.Max(math.Abs(float64(eX-sX)), math.Abs(float64(eY-sY))))

		for i := 0; i <= top; i++ {
			linesMatrix[sY][sX]++
			if sY < eY {
				sY++
			} else if sY > eY {
				sY--
			}
			if sX < eX {
				sX++
			} else if sX > eX {
				sX--
			}
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
