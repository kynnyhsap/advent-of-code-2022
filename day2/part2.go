package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const WIN_POINTS = 6
const DRAW_POINTS = 3

func calculateMyFightPoints(enemyChoise byte, expectedOutcome byte) int {
	points := 0

	if expectedOutcome == 'X' {
		// expected to loose

		if enemyChoise == 'A' {
			// I need to choose Z (Scissors)
			points += 3
		}

		if enemyChoise == 'B' {
			// I need to chooze X (Rock)
			points += 1
		}

		if enemyChoise == 'C' {
			// I need to choose Y (Paper)
			points += 2
		}
	}

	if expectedOutcome == 'Y' {
		// expected to draw
		points += 3 + int(enemyChoise-'A'+1)
	}

	if expectedOutcome == 'Z' {
		// expected to win
		points += 6

		if enemyChoise == 'A' {
			// I need to choose Y (Paper)
			points += 2
		}

		if enemyChoise == 'B' {
			// I need to choose Z (Scissors)
			points += 3
		}

		if enemyChoise == 'C' {
			// I need to chooze X (Rock)
			points += 1
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
