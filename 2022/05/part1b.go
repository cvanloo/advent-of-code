package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Stack interface {
	Push(rune) Stack
	Pop() (Stack, error)
	Top() (rune, error)
}

type EmptyStack struct{}

// EmptyStack implements Stack
var _ Stack = (*EmptyStack)(nil)

func NewStack() Stack {
	return &EmptyStack{}
}

// Pop implements Stack
func (e *EmptyStack) Pop() (Stack, error) {
	return nil, errors.New("must not Pop() from empty stack")
}

// Push implements Stack
func (e *EmptyStack) Push(r rune) Stack {
	return &NonEmptyStack{
		value: r,
		last:  e,
	}
}

// Top implements Stack
func (e *EmptyStack) Top() (rune, error) {
	var zero rune
	return zero, errors.New("must not Top() from empty stack")
}

func (e EmptyStack) String() string {
	return fmt.Sprintf("END")
}

type NonEmptyStack struct {
	value rune
	last  Stack
}

// NonEmptyStack implements Stack.
var _ Stack = (*NonEmptyStack)(nil)

// Pop implements Stack
func (s *NonEmptyStack) Pop() (Stack, error) {
	return s.last, nil
}

// Push implements Stack
func (s *NonEmptyStack) Push(r rune) Stack {
	return &NonEmptyStack{
		value: r,
		last:  s,
	}
}

// Top implements Stack
func (s *NonEmptyStack) Top() (rune, error) {
	return s.value, nil
}

func (s NonEmptyStack) String() string {
	return fmt.Sprintf("%s -> %s", string(s.value), s.last)
}

type Instruction struct {
	amount int
	from   int
	to     int
}

func (ins Instruction) Apply(stacks []Stack) {
	for i := 0; i < ins.amount; i++ {
		val, _ := stacks[ins.from].Top()
		newToStack := stacks[ins.to].Push(val)
		newFromStack, _ := stacks[ins.from].Pop()

		stacks[ins.to] = newToStack
		stacks[ins.from] = newFromStack
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(input string) (stacks []Stack, instructions []Instruction, err error) {
	input = strings.TrimRight(input, "\n")

	inputParts := strings.Split(input, "\n\n")
	stackStrs := strings.Split(inputParts[0], "\n")
	instructionStrs := strings.Split(inputParts[1], "\n")

	stackStrs = stackStrs[0 : len(stackStrs)-1]

	for i := len(stackStrs)-1; i >= 0; i-- {
		stackStr := stackStrs[i]

		start, end := 0, 4
		stackPos := -1

		for start < end {
			currEl := []rune(stackStr[start:end])

			// advance position
			stackPos++
			start = end
			end += 4
			if end > len(stackStr) {
				end = len(stackStr)
			}

			if stackPos >= len(stacks) {
				stacks = append(stacks, NewStack())
			}

			// get current stack element value
			value := currEl[1]
			if value != ' ' {
				newStack := stacks[stackPos].Push(value)
				stacks[stackPos] = newStack
			}
		}
	}

	for _, insStr := range instructionStrs {
		parts := strings.Split(insStr, " ")

		amount, err := strconv.Atoi(parts[1])
		panicIf(err)
		from, err := strconv.Atoi(parts[3])
		panicIf(err)
		to, err := strconv.Atoi(parts[5])
		panicIf(err)

		ins := Instruction{
			amount: amount,
			from:   from - 1,
			to:     to - 1,
		}
		instructions = append(instructions, ins)
	}

	return stacks, instructions, nil
}

func copyToClipboard(text string) {
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
		val, _ := stack.Top()
		result += string(val)
	}

	copyToClipboard(result)
	fmt.Printf("Result: %s (copied to clipboard)\n", result)
}
