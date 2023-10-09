package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	priorities := 0
	for scanner.Scan() {
		rucksack := []byte(scanner.Text())

		compartment1 := rucksack[0 : len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2 : len(rucksack)]

		for _, item1 := range compartment1 {
			found := false

			for _, item2 := range compartment2 {
				if item2 == item1 {
					found = true
					priorities += getPriority(item1)
					break
				}
			}

			if found {
				break
			}
		}
	}

	fmt.Println(priorities)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getPriority(item byte) int {
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}

	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}

	panic("uknown item")
}
