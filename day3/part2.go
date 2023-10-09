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

	var group [3][]byte

	group[0] = []byte{}
	group[1] = []byte{}
	group[2] = []byte{}

	i := 0
	for scanner.Scan() {

		group[i] = []byte(scanner.Text())

		i += 1

		if i == 3 {
			// calc group priorities

			fmt.Println("\nGroup:")
			fmt.Println(" - elf 1  ", string(group[0]))
			fmt.Println(" - elf 2  ", string(group[1]))
			fmt.Println(" - elf 3  ", string(group[2]))

		OuterLoop:
			for _, item1 := range group[0] {
				for _, item2 := range group[1] {
					for _, item3 := range group[2] {
						if item1 == item2 && item2 == item3 {
							fmt.Println("[badge]: ", string(item1))

							priorities += getPriority(item1)
							break OuterLoop
						}
					}
				}
			}

			// cleanup
			i = 0

			group[0] = []byte{}
			group[1] = []byte{}
			group[2] = []byte{}
		}
	}

	fmt.Println("\npriorities: ", priorities)

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
