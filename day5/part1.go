package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rearrangement struct {
	amount int
	from   int
	to     int
}

type CargoStack struct {
	stack []byte
}

func (cs *CargoStack) empty() bool {
	return len(cs.stack) == 0
}

func (cs *CargoStack) push(cargo byte) {
	cs.stack = append(cs.stack, cargo)
}

func (cs *CargoStack) pop() byte {

	if cs.empty() {
		panic("empty stack")
	}

	last := cs.stack[len(cs.stack)-1]

	cs.stack = cs.stack[:len(cs.stack)-1]

	return last
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	splitted := strings.Split(string(content), "\n\n")
	staksStrings := strings.Split(splitted[0], "\n")
	staksStrings = staksStrings[:len(staksStrings)-1] // removing last numerical row

	stacks := make([]CargoStack, 9)

	for j := len(staksStrings) - 1; j >= 0; j -= 1 {
		for i := 0; i < 9; i += 1 {
			cargo := staksStrings[j][i*4+1]

			if 'A' <= cargo && cargo <= 'Z' {
				stacks[i].push(cargo)
			}
		}
	}

	// fmt.Println(stacks)

	rearrangementsString := strings.Split(splitted[1], "\n")

	for _, rearrangementString := range rearrangementsString {
		r := strings.Split(rearrangementString, " ")

		amount, _ := strconv.Atoi(r[1])
		from, _ := strconv.Atoi(r[3])
		to, _ := strconv.Atoi(r[5])

		from -= 1 // adjust to indeces
		to -= 1   // adjust to indeces

		for k := 0; k < amount; k += 1 {
			stacks[to].push(stacks[from].pop())
		}
	}

	// fmt.Println(stacks)

	top := make([]byte, 9)

	for _, stack := range stacks {
		top = append(top, stack.pop())
	}

	fmt.Println(string(top))

}
