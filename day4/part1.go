package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")

		x1, _ := strconv.Atoi(strings.Split(ranges[0], "-")[0])
		x2, _ := strconv.Atoi(strings.Split(ranges[0], "-")[1])

		y1, _ := strconv.Atoi(strings.Split(ranges[1], "-")[0])
		y2, _ := strconv.Atoi(strings.Split(ranges[1], "-")[1])

		// [x1 ... {y1 ... y2} ... x2] or [y1 ... {x1 ... y2} ... y2]
		if (x1 <= y1 && y2 <= x2) || (y1 <= x1 && x2 <= y2) {
			sum += 1
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
