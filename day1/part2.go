package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves := make([]int, 0)
	current := 0

	for scanner.Scan() {
		calories, err := strconv.Atoi(scanner.Text())

		if err != nil {
			// if the line is not a number then it's the end of the elf's calories list

			elves = append(elves, current)

			current = 0
		} else {
			current += calories
		}
	}

	sort.Ints(elves)

	total := 0
	top3Elves := elves[len(elves)-3:]
	for _, value := range top3Elves {
		total += value
	}

	println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
