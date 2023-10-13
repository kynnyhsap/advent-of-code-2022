package main

import (
	"fmt"
	"log"
	"os"
)

const PART_1 = 4
const PART_2 = 14

const SIZE = PART_2

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(content)-SIZE; i += 1 {
		if isMarker(content[i : i+SIZE]) {
			fmt.Println(i + SIZE)
			break
		}
	}

}

func isMarker(characters []byte) bool {
	unique := make([]byte, 0)

	for _, c := range characters {
		for _, u := range unique {
			if c == u {
				return false
			}
		}

		unique = append(unique, c)
	}

	return true
}
