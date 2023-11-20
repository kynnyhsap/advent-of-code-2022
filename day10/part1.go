package day10

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const OFFSET = 20
const JUMP = 40

func Part1() {
	input, _ := os.ReadFile("./day10/input.txt")
	lines := bytes.Split(input, []byte("\n"))

	x := 1
	cycle := 1

	signalStrengths := 0

	for _, instruction := range lines {
		// increment cycle both for noop and addx
		cycle += 1
		signalStrengths += checkpoint(cycle, x)

		if instruction[0] == 'a' {
			// addx
			v, _ := strconv.Atoi(string(instruction[5:]))

			cycle += 1
			x += v
			signalStrengths += checkpoint(cycle, x)
		}
	}

	fmt.Println(signalStrengths)
}

func checkpoint(cycle int, x int) int {
	if cycle%OFFSET == 0 && (cycle/OFFSET)%2 != 0 {
		fmt.Printf("cycle [%d] X: %d\n", cycle, x)
		return cycle * x
	}

	return 0
}
