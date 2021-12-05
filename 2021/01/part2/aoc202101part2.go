package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func log(err error) {
	fmt.Fprintf(os.Stderr, "encountered error: %v\n", err)
}

type window struct {
	elements []int
}

func (w *window) push(depth int) {
	w.elements = append(w.elements, depth)
}

func (w *window) count() int {
	return len(w.elements)
}

func (w *window) last() []int {
	return w.elements[w.count()-4 : w.count()-1]
}

func (w *window) curr() []int {
	return w.elements[w.count()-3:]
}

func sum(i []int) int {
	sum := 0
	for _, val := range i {
		sum += val
	}
	return sum
}

var increased int
var depthWindow window

func handle(depth int) {
	depthWindow.push(depth)

	if depthWindow.count() > 3 {
		fmt.Printf("last: %v\n", depthWindow.last())
		fmt.Printf("curr: %v\n", depthWindow.curr())

		if sum(depthWindow.curr()) > sum(depthWindow.last()) {
			increased++
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	//file, err := os.Open("test.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log(err)
		} else {
			handle(value)
		}
	}

	check(scanner.Err())

	fmt.Printf("Depth increased %d times.\n", increased)
}
