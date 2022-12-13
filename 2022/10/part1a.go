package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var signalStrengths []int

func calculateSignalStrength(cycle, xreg int) {
	if (cycle - 20) % 40 == 0 {
		signalStrengths = append(signalStrengths, xreg * cycle)
		fmt.Printf("cycle: %d, xreg: %d, signal: %d\n", cycle, xreg, xreg * cycle)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	xreg, cycle := 1, 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		instruction := parts[0]
		if instruction == "addx" {
			value, _ := strconv.Atoi(parts[1])
			cycle += 1
			calculateSignalStrength(cycle, xreg)
			cycle += 1
			calculateSignalStrength(cycle, xreg)
			xreg += value
		} else {
			// noop
			cycle += 1
			calculateSignalStrength(cycle, xreg)
		}
	}
	totalStrength := 0
	for _, strength := range signalStrengths {
		totalStrength += strength
	}
	fmt.Printf("%v\n", totalStrength)
}
