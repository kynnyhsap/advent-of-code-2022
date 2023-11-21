package day10

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const H = 6
const W = 40

type CRT [H * W]int

func Part2() {
	input, _ := os.ReadFile("./day10/input.txt")
	lines := bytes.Split(input, []byte("\n"))

	var crt CRT

	x := 1

	cycle := 0
	drawPixel(&crt, cycle, x)

	for _, instruction := range lines {
		// increment cycle both for noop and addx
		cycle += 1
		drawPixel(&crt, cycle, x)

		if instruction[0] == 'a' {
			v, _ := strconv.Atoi(string(instruction[5:])) // addx

			x += v
			cycle += 1

			drawPixel(&crt, cycle, x)
		}
	}

	printCRT(&crt)
}

func drawPixel(crt *CRT, pixel int, x int) {
	// sprite leght is 3, x marking the middle [ x-1, x, x+1 ]
	row := pixel / W

	t := x + row*W // this is required, but it is not described in the problem statement -_-

	if pixel == t-1 || pixel == t+1 || pixel == t {
		crt[pixel] = 1
	}
}

func printCRT(crt *CRT) {
	fmt.Println()
	for i := 0; i < H*W; i += 1 {
		if i > 0 && i%(W) == 0 {
			fmt.Println()
		}

		if crt[i] == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("#")
		}
	}
	fmt.Println()
}
