package main

import (
	"fmt"
	"sort"
)

type MathFunc func(a, b int) int

var operationsMap = map[string]MathFunc{
	"+": func(a, b int) int {
		return a + b
	},
	"*": func(a, b int) int {
		return a * b
	},
	"^": func(a, b int) int {
		return a * a
	},
}

type Operation struct {
	op MathFunc
	b  int
}

type Test interface {
	Test(a int) bool
}

type IsDivisible struct {
	by int
}

func (t IsDivisible) Test(a int) bool {
	return a%t.by == 0
}

type Monkey struct {
	items                 []int
	operation             Operation
	test                  Test
	onTrueIdx, onFalseIdx int
}

func main() {
	//worryManagement := 96577
	worryManagement := 9699690

	// And you know what? I don't even feel bad about it!
	/*
		m0 := &Monkey{
			items: []int{79, 98},
			operation: Operation{
				op: operationsMap["*"],
				b:  19,
			},
			test:       IsDivisible{23},
			onTrueIdx:  2,
			onFalseIdx: 3,
		}
		m1 := &Monkey{
			items: []int{54, 65, 75, 74},
			operation: Operation{
				op: operationsMap["+"],
				b:  6,
			},
			test:       IsDivisible{19},
			onTrueIdx:  2,
			onFalseIdx: 0,
		}
		m2 := &Monkey{
			items: []int{79, 60, 97},
			operation: Operation{
				op: operationsMap["^"],
				b:  0,
			},
			test:       IsDivisible{13},
			onTrueIdx:  1,
			onFalseIdx: 3,
		}
		m3 := &Monkey{
			items: []int{74},
			operation: Operation{
				op: operationsMap["+"],
				b:  3,
			},
			test:       IsDivisible{17},
			onTrueIdx:  0,
			onFalseIdx: 1,
		}*/
	m1 := &Monkey{
		items: []int{89, 73, 66, 57, 64, 80},
		operation: Operation{
			op: operationsMap["*"],
			b:  3,
		},
		test:       IsDivisible{13},
		onTrueIdx:  6,
		onFalseIdx: 2,
	}
	m2 := &Monkey{
		items: []int{83, 78, 81, 55, 81, 59, 69},
		operation: Operation{
			op: operationsMap["+"],
			b:  1,
		},
		test:       IsDivisible{3},
		onTrueIdx:  7,
		onFalseIdx: 4,
	}
	m3 := &Monkey{
		items: []int{76, 91, 58, 85},
		operation: Operation{
			op: operationsMap["*"],
			b:  13,
		},
		test:       IsDivisible{7},
		onTrueIdx:  1,
		onFalseIdx: 4,
	}
	m4 := &Monkey{
		items: []int{71, 72, 74, 76, 68},
		operation: Operation{
			op: operationsMap["^"],
			b:  0, // lol
		},
		test:       IsDivisible{2},
		onTrueIdx:  6,
		onFalseIdx: 0,
	}
	m5 := &Monkey{
		items: []int{98, 85, 84},
		operation: Operation{
			op: operationsMap["+"],
			b:  7,
		},
		test:       IsDivisible{19},
		onTrueIdx:  5,
		onFalseIdx: 7,
	}
	m6 := &Monkey{
		items: []int{78},
		operation: Operation{
			op: operationsMap["+"],
			b:  8,
		},
		test:       IsDivisible{5},
		onTrueIdx:  3,
		onFalseIdx: 0,
	}
	m7 := &Monkey{
		items: []int{86, 70, 60, 88, 88, 78, 74, 83},
		operation: Operation{
			op: operationsMap["+"],
			b:  4,
		},
		test:       IsDivisible{11},
		onTrueIdx:  1,
		onFalseIdx: 2,
	}
	m8 := &Monkey{
		items: []int{81, 58},
		operation: Operation{
			op: operationsMap["+"],
			b:  5,
		},
		test:       IsDivisible{17},
		onTrueIdx:  3,
		onFalseIdx: 5,
	}
	monkeys := []*Monkey{m1, m2, m3, m4, m5, m6, m7, m8}
	//monkeys := []*Monkey{m0, m1, m2, m3}

	const numRounds = 10_000
	//const numRounds = 20
	monkeyInspects := make([]int, len(monkeys))
	for round := 0; round < numRounds; round++ {
		for monkeyIdx, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeyInspects[monkeyIdx]++
				op := monkey.operation
				worryAfter := op.op(item, op.b) % worryManagement
				if monkey.test.Test(worryAfter) {
					monkey.items = monkey.items[1:]
					other := monkeys[monkey.onTrueIdx]
					other.items = append(other.items, worryAfter)
				} else {
					monkey.items = monkey.items[1:]
					other := monkeys[monkey.onFalseIdx]
					other.items = append(other.items, worryAfter)
				}
			}
		}
		if round == 0 || round == 19 || round == 999 || round == 1999 {
			fmt.Printf("After round %d\n", round+1)
			for i, inspects := range monkeyInspects {
				fmt.Printf("Monkey: %d inspected items %d times.\n", i, inspects)
			}
		}
	}

	worries := ""
	for i, monkey := range monkeys {
		worries += fmt.Sprintf("%d: %v, ", i, monkey.items)
	}
	fmt.Println(worries)

	for i, inspects := range monkeyInspects {
		fmt.Printf("Monkey: %d inspected items %d times.\n", i, inspects)
	}

	sort.Slice(monkeyInspects, func(i, j int) bool {
		return monkeyInspects[i] > monkeyInspects[j]
	})
	fmt.Printf("Monkey business: %d\n", monkeyInspects[0]*monkeyInspects[1])
}
