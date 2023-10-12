package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	splitted := strings.Split(string(content), "\n\n")
	staksStrings := strings.Split(splitted[0], "\n")
	staksStrings = staksStrings[:len(staksStrings)-1] // removing last numerical row

	stacks := make([][]byte, 9)

	for j := len(staksStrings) - 1; j >= 0; j -= 1 {
		for i := 0; i < 9; i += 1 {
			cargo := staksStrings[j][i*4+1]

			if 'A' <= cargo && cargo <= 'Z' {
				stacks[i] = append(stacks[i], cargo)
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

		l := len(stacks[from]) - amount

		topCargos := stacks[from][l:]

		stacks[to] = append(stacks[to], topCargos...)
		stacks[from] = stacks[from][:l]
	}

	// fmt.Println(stacks)

	top := make([]byte, 9)

	for _, stack := range stacks {
		top = append(top, stack[len(stack)-1])
	}

	fmt.Println(string(top))

}
