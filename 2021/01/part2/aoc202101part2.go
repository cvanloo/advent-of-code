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

// window is a wrap-around buffer
type window struct {
	elements [4]int
	next     int
	wrapped  bool
}

func (w *window) push(depth int) {
	w.elements[w.next] = depth
	if w.next >= 3 {
		w.next = 0
		w.wrapped = true
	} else {
		w.next++
	}
}

func (w *window) last() []int {
	upperLimit := w.next + 3
	var h1, h2 []int

	if upperLimit > 4 {
		h1 = w.elements[w.next:]
		lowerLimit := upperLimit - 4
		h2 = w.elements[:lowerLimit]
	} else {
		h1 = w.elements[w.next:upperLimit]
	}

	return append(h1, h2[:]...)
}

func (w *window) curr() []int {
	lowerLimit := w.next + 1
	var h1, h2 []int

	if lowerLimit == 4 {
		h1 = w.elements[0:3]
	} else {
		h1 = w.elements[lowerLimit:]
		h2 = w.elements[:w.next]
	}

	return append(h1, h2[:]...)
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

	if depthWindow.wrapped {
		//fmt.Printf("full: %v\n", depthWindow.elements)
		//fmt.Printf("last: %v\n", depthWindow.last())
		//fmt.Printf("curr: %v\n", depthWindow.curr())

		curr := sum(depthWindow.curr())
		last := sum(depthWindow.last())
		if curr > last {
			increased++
			fmt.Println("+increased:", curr)
		} else if curr == last {
			fmt.Println("~no change:", curr)
		} else {
			fmt.Println("-decreased:", curr)
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
