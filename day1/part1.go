package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentElfCalories := 0
	fatestElfCalories := 0

	for scanner.Scan() {
		calories, err := strconv.Atoi(scanner.Text())

		if err != nil {
			// if the line is not a number then it's the end of the elf's calories list

			if currentElfCalories > fatestElfCalories {
				fatestElfCalories = currentElfCalories
			}

			currentElfCalories = 0
		} else {
			currentElfCalories += calories
		}
	}

	fmt.Println("Fatest elf calories:", fatestElfCalories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
