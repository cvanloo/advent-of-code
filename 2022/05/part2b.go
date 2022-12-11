package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Stack struct {
	elements []rune
}

func (s *Stack) Push(r []rune) {
	s.elements = append(s.elements, r...)
}

func (s *Stack) Pop(amount int) []rune {
	length := len(s.elements)
	elements := s.elements[length-amount : length]
	s.elements = s.elements[0 : length-amount]
	return elements
}

func (s Stack) Peek() rune {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Move(to *Stack, amount int) {
	to.Push(s.Pop(amount))
}

func (s Stack) String() string {
	return fmt.Sprintf("%s", string(s.elements))
}

type Instruction struct {
	amount int
	from   int
	to     int
}

func (ins Instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", ins.amount, ins.from, ins.to)
}

func (ins Instruction) Apply(stacks []Stack) {
	from := &stacks[ins.from]
	to := &stacks[ins.to]

	from.Move(to, ins.amount)
}

type intParser struct {
	err error
}

func (ip *intParser) Parse(s string) (res int) {
	if ip.err == nil {
		res, ip.err = strconv.Atoi(s)
	}
	return
}

func parseInput(input string) ([]Stack, []Instruction, error) {
	input = strings.TrimRight(input, "\n")
	inputParts := strings.Split(input, "\n\n")

	stackStrs := strings.Split(inputParts[0], "\n")
	stackStrs = stackStrs[0 : len(stackStrs)-1] // Ignore last line (it only numbers the stacks)

	// Based on the last stack line we can get the total number of crates per stack.
	// The last line is the only stack line where every stack must have at least one element.
	lastStackLine := stackStrs[len(stackStrs)-1]
	numberOfCrates := len(strings.Split(lastStackLine, " "))

	stacks := make([]Stack, numberOfCrates)
	stackLength := len(stackStrs) - 1

	for i := stackLength; i >= 0; i-- {
		stackStr := stackStrs[i]

		for j := 0; j < numberOfCrates; j++ {
			crate := []rune(stackStr)[j*4+1]

			if crate != ' ' {
				stacks[j].Push([]rune{crate})
			}
		}
	}

	instructionStrs := strings.Split(inputParts[1], "\n")
	instructions := make([]Instruction, len(instructionStrs))
	intParser := intParser{}

	for i, insStr := range instructionStrs {
		parts := strings.Split(insStr, " ")

		amount := intParser.Parse(parts[1])
		from := intParser.Parse(parts[3])
		to := intParser.Parse(parts[5])

		if err := intParser.err; err != nil {
			return nil, nil, fmt.Errorf("failed to parse instructions: %v", err)
		}

		ins := Instruction{
			amount: amount,
			from:   from - 1,
			to:     to - 1,
		}
		instructions[i] = ins
	}

	return stacks, instructions, nil
}

func copyToClipboard(text string) {
	panicIf := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	if runtime.GOOS != "linux" {
		panic("have you heard of our lord and saviour, GNU/Linux?")
	}

	if _, err := exec.LookPath("xsel"); err != nil {
		panic("paru -Syu xsel")
	}

	command := exec.Command("xsel", "-bi")
	stdin, err := command.StdinPipe()
	panicIf(err)

	err = command.Start()
	panicIf(err)

	io.WriteString(stdin, text)
	stdin.Close() // Command won't exit until stdin is closed.
	err = command.Wait()
	panicIf(err)
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	stacks, instructions, err := parseInput(string(input))
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	for _, instruction := range instructions {
		instruction.Apply(stacks)
	}

	var result string
	for _, stack := range stacks {
		result += string(stack.Peek())
	}

	copyToClipboard(result)
	fmt.Printf("Result: %s (copied to clipboard)\n", result)
}
