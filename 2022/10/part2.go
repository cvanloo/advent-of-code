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
	video VideoBuffer
}

func (c *Cpu) Step() {
	c.cycle++
	c.video.Draw(c.xreg)
}

type VideoBuffer interface {
	Draw(pos int)
}

const height, width = 6, 40

type Crt struct {
	screenBuffer [height*width]string
	frameIndex int
}

func (c *Crt) Draw(hpos int) {
	pixel := "."
	idx := c.frameIndex

	lineIdx := idx % width
	if lineIdx >= hpos-1 && lineIdx <= hpos+1 {
		pixel = "#"
	}

	c.screenBuffer[idx] = pixel

	idx++
	if idx > height*width {
		c.frameIndex = 0
	} else {
		c.frameIndex = idx
	}
}

func (c *Crt) String() string {
	output := ""
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			output += c.screenBuffer[h*width+w]
		}
		output += "\n"
	}
	return output
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
	video := &Crt{}
	cpu := &Cpu{1,0, nil, video}
	for scanner.Scan() {
		line := scanner.Text()
		instruction, err := parseInstruction(line)
		if err != nil {
			fmt.Printf("instruction error: %v\n", err)
			continue
		}

		instruction.Execute(cpu)
	}

	fmt.Printf("%s", video)
}
