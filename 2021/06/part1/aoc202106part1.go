package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(input string) []laternfish {
	fish := []laternfish{}
	input = strings.Trim(input, "\n")
	parts := strings.Split(input, ",")

	for _, v := range parts {
		i, err := strconv.Atoi(v)
		panicif(err)
		fish = append(fish, laternfish{i})
	}

	return fish
}

type laternfish struct {
	timer int
}

func (f *laternfish) reproduce() laternfish {
	return laternfish{8}
}

type fishSim struct {
	fish []laternfish
}

func (s *fishSim) tick() {
	newFish := []laternfish{}

	for i := 0; i < len(s.fish); i++ {
		f := &s.fish[i]

		if f.timer == 0 {
			f.timer = 6
			newFish = append(newFish, f.reproduce())
		} else {
			f.timer--
		}
	}

	s.fish = append(s.fish, newFish...)
}

//const simDays = 18
//const simDays = 80
const simDays = 256

func main() {
	bytes, err := os.ReadFile("input.txt")
	panicif(err)

	initialFish := parseInput(string(bytes))

	sim := fishSim{initialFish}

	for i := 0; i < simDays; i++ {
		sim.tick()
	}

	fmt.Printf("There are %d laternfish after %d days\n", len(sim.fish), simDays)
}
