package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const WIN_POINTS = 6
const DRAW_POINTS = 3

func calculateMyFightPoints(enemyChoise byte, myChoice byte) int {
	points := int(myChoice - 'X' + 1)

	if enemyChoise == 'A' {
		if myChoice == 'X' {
			// draw
			points += DRAW_POINTS
		}

		if myChoice == 'Y' {
			// win
			points += WIN_POINTS
		}
	}

	if enemyChoise == 'B' {
		if myChoice == 'Y' {
			// draw
			points += DRAW_POINTS
		}

		if myChoice == 'Z' {
			// win
			points += WIN_POINTS
		}
	}

	if enemyChoise == 'C' {
		if myChoice == 'X' {
			// win
			points += WIN_POINTS
		}

		if myChoice == 'Z' {
			// draw
			points += DRAW_POINTS
		}
	}

	return points
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()

		totalPoints += calculateMyFightPoints(line[0], line[2])
	}

	fmt.Println(totalPoints)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
