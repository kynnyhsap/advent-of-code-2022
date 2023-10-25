package day9

import (
	"bytes"
	"fmt"
	"os"
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

	// these points represet allowed boundaries for tail to be considered adjecent to head
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

	lines := bytes.Split(input, []byte("\n"))
	for _, motion := range lines {
		direction := motion[0]
		steps := int(motion[2] - '0')

		for step := 0; step < steps; step += 1 {
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

			isTailAdjacent := false
			for _, point := range adjacent {
				if point.x+head.x == tail.x && point.y+head.y == tail.y {
					isTailAdjacent = true
					break
				}
			}
			if isTailAdjacent {
				continue
			}

			// move tail
			for from, to := range gaps {
				if from.x+head.x == tail.x && from.y+head.y == tail.y {
					tail.x += to.x
					tail.y += to.y
					break
				}
			}

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
