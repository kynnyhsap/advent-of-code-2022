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

const ROPE_SIZE = 10

func Part2() {
	input, _ := os.ReadFile("./day9/input.txt")

	rope := make([]Point, ROPE_SIZE)
	for i := range rope {
		rope[i] = Point{0, 0}
	}

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

	gaps[Point{-2, -2}] = Point{-1, -1}
	gaps[Point{-2, 2}] = Point{-1, 1}
	gaps[Point{2, -2}] = Point{1, -1}
	gaps[Point{2, 2}] = Point{1, 1}

	lines := bytes.Split(input, []byte("\n"))

	fmt.Println("START")
	printRope(rope)

	for _, motion := range lines {
		direction := motion[0]
		steps, _ := strconv.Atoi(string(motion)[2:])

		fmt.Printf("\n\n== %c %d ==\n\n", direction, steps)

		for step := 0; step < steps; step += 1 {

			for i := range rope[:ROPE_SIZE-1] {
				h := i
				t := i + 1

				// fmt.Printf("\n- Direction: %c    Step: %d   Head Index: %d     Tail Index: %d  \n", direction, step+1, h, t)

				// move head
				if i == 0 {
					if direction == 'U' {
						rope[h].y += 1
					} else if direction == 'D' {
						rope[h].y -= 1
					} else if direction == 'R' {
						rope[h].x += 1
					} else if direction == 'L' {
						rope[h].x -= 1
					} else {
						panic("Unkown direction")
					}
				}

				// printRope(rope)

				isTailAdjacent := false

				for _, point := range adjacent {
					if point.x+rope[h].x == rope[t].x && point.y+rope[h].y == rope[t].y {
						isTailAdjacent = true
						break
					}
				}

				// fmt.Println("isTailAdjacent", isTailAdjacent, rope[h], rope[t])
				if isTailAdjacent {
					continue
				}

				// move tail
				for from, to := range gaps {
					if rope[t].x-rope[h].x == from.x && rope[t].y-rope[h].y == from.y {
						rope[t].x = rope[h].x + to.x
						rope[t].y = rope[h].y + to.y
						break
					}
				}

				// printRope(rope)
			}

			tail := rope[ROPE_SIZE-1]
			fmt.Println("tail", tail)

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

		printRope(rope)

	}

	fmt.Println(len(visitedTailLocations))
	printTrail(visitedTailLocations)
}

const P = 20

func printRope(rope []Point) {
	fmt.Println("")

	// h := rope[0]

	for y := P; y >= -P; y -= 1 {
		fmt.Printf("\t")
		for x := -P; x <= P; x += 1 {

			isEmpty := true
		Inner:
			for i, knot := range rope {
				if x == knot.x && y == knot.y {
					if i == 0 {
						fmt.Print("H")
					} else {
						fmt.Printf("%d", i)
					}
					isEmpty = false
					break Inner
				}
			}

			if isEmpty {
				if x == 0 && y == 0 {
					fmt.Printf("s")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Println("")
	}
	// fmt.Println(rope)
}

func printTrail(visited []Point) {
	for y := P; y >= -P; y -= 1 {
		fmt.Printf("\t")
		for x := -P; x <= P; x += 1 {

			isEmpty := true
			for _, v := range visited {
				if v.x == x && v.y == y {
					fmt.Printf("#")
					isEmpty = false
					break
				}
			}

			if isEmpty {
				if x == 0 && y == 0 {
					fmt.Printf("s")
				} else {
					fmt.Printf(".")
				}
			}

		}
		fmt.Printf("\n")
	}
}
