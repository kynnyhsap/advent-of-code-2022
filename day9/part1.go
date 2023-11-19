package day9

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

func Part1() {

	input, _ := os.ReadFile("./day9/input.txt")

	head := Point{0, 0}
	tail := Point{0, 0}

	visitedTailLocations := make([]Point, 0)

	visitedTailLocations = append(visitedTailLocations, Point{0, 0})

	// these points represent allowed boundaries for tail to be considered adjecent to head
	adjacent := []Point{
		{-1, 1},  // top left
		{0, 1},   // top middle
		{1, 1},   // top right
		{-1, 0},  // middle left
		{0, 0},   // middle middle
		{1, 0},   // middle right
		{-1, -1}, // bottom left
		{0, -1},  // bottom middle
		{1, -1},  // bottom right
	}

	// gaps represent a map points of how tail should adjust to head to become adjecent again
	gaps := make(map[Point]Point)

	// horizontal adjustments
	gaps[Point{-2, 0}] = Point{-1, 0}
	gaps[Point{2, 0}] = Point{1, 0}
	// vartical adjustments
	gaps[Point{0, -2}] = Point{0, -1}
	gaps[Point{0, 2}] = Point{0, 1}
	// diagonal adjustments
	gaps[Point{-1, -2}] = Point{0, -1}
	gaps[Point{1, -2}] = Point{0, -1}
	gaps[Point{-1, 2}] = Point{0, 1}
	gaps[Point{1, 2}] = Point{0, 1}
	gaps[Point{-2, -1}] = Point{-1, 0}
	gaps[Point{-2, 1}] = Point{-1, 0}
	gaps[Point{2, -1}] = Point{1, 0}
	gaps[Point{2, 1}] = Point{1, 0}

	// for from, to := range gaps {
	// 	fmt.Println("----------------------------------------------------------------")
	// 	printHeadAndTail(&head, &from)
	// 	printHeadAndTail(&head, &to)
	// 	fmt.Println("----------------------------------------------------------------")
	// }

	tailMovesCount := 0

	lines := bytes.Split(input, []byte("\n"))

	for _, motion := range lines {
		direction := motion[0]
		steps, _ := strconv.Atoi(string(motion)[2:])

		fmt.Printf("\n-----------\nMove %c %d times\n", direction, steps)

		for step := 0; step < steps; step += 1 {
			// fmt.Printf("- [Step %d]: starting position\n", step+1)
			// printHeadAndTail(&head, &tail)

			// move head
			if direction == 'U' {
				head.y += 1
			} else if direction == 'D' {
				head.y -= 1
			} else if direction == 'R' {
				head.x += 1
			} else if direction == 'L' {
				head.x -= 1
			} else {
				panic("Unkown direction")
			}

			// fmt.Printf("Moved head %c\n", direction)
			// printHeadAndTail(&head, &tail)

			isTailAdjacent := false
			for _, point := range adjacent {
				if point.x+head.x == tail.x && point.y+head.y == tail.y {
					isTailAdjacent = true
					break
				}
			}
			if isTailAdjacent {
				// fmt.Println("Tail is adjecent, no need to move")
				printHeadAndTail(&head, &tail)
				continue
			}

			// move tail
			for from, to := range gaps {

				if tail.x-head.x == from.x && tail.y-head.y == from.y {
					tail.x = head.x + to.x
					tail.y = head.y + to.y
					break
				}
			}

			tailMovesCount += 1
			// fmt.Println("Moved tail")
			printHeadAndTail(&head, &tail)

			// record tail position
			visited := false
			for _, loc := range visitedTailLocations {
				if loc.x == tail.x && loc.y == tail.y {
					visited = true
					break
				}
			}
			if !visited {
				visitedTailLocations = append(visitedTailLocations, Point{tail.x, tail.y})
			}
		}
	}

	fmt.Println(len(visitedTailLocations))
}

const S = 3

func printHeadAndTail(h *Point, t *Point) {
	// fmt.Println("")
	// for y := h.y + S; y >= h.y-S; y -= 1 {
	// 	fmt.Printf("\t")
	// 	for x := h.x - S; x <= h.x+S; x += 1 {
	// 		if x == h.x && y == h.y {
	// 			fmt.Printf("H")
	// 		} else if x == t.x && y == t.y {
	// 			fmt.Printf("T")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
	// fmt.Println("")
	fmt.Printf("\tHead (%d, %d)  Tail (%d, %d)\n", h.x, h.y, t.x, t.y)
}
