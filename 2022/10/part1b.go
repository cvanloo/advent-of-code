package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cpu struct {
	xreg, cycle int
	signals []int
}

func (c *Cpu) Step() {
	c.cycle++
	if (c.cycle - 20) % 40 == 0 {
		c.signals = append(c.signals, c.xreg * c.cycle)
	}
}

type Instruction interface {
	Execute(cpu *Cpu)
}

type AddX struct {
	value int
}

func (a AddX) Execute(cpu *Cpu) {
	cpu.Step()
	cpu.Step()
	cpu.xreg += a.value
}

type Noop struct{}

func (n Noop) Execute(cpu *Cpu) {
	cpu.Step()
}

func parseInstruction(instructionStr string) (Instruction, error) {
	parts := strings.Fields(instructionStr)
	instruction := parts[0]

	switch instruction {
	case "addx":
		value, err := strconv.Atoi(parts[1])
		return AddX{value}, err
	case "noop":
		return Noop{}, nil
	}
	return nil, fmt.Errorf("instruction not recognized: %s", instruction)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cpu := &Cpu{1,0, nil}
	for scanner.Scan() {
		line := scanner.Text()
		instruction, err := parseInstruction(line)
		if err != nil {
			fmt.Printf("instruction error: %v\n", err)
			continue
		}

		instruction.Execute(cpu)
	}

	totalStrength := 0
	for _, strength := range cpu.signals {
		totalStrength += strength
	}
	fmt.Printf("%v\n", totalStrength)
}
