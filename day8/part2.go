package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	bytesMatrix := bytes.Split(input, []byte("\n"))

	size := len(bytesMatrix) // matrix size

	matrix := make([][]int, size)

	// construct trees matrix
	for r := range bytesMatrix {
		matrix[r] = make([]int, size)
		for c := range bytesMatrix[r] {
			matrix[r][c] = int(bytesMatrix[r][c] - '0')
		}
	}

	highestScenicScore := 0

	for r := range matrix {
		for c := range matrix[r] {
			score := scenicScore(matrix, r, c)
			if score > highestScenicScore {
				highestScenicScore = score
			}
		}
	}

	fmt.Println(highestScenicScore)
}

func scenicScore(matrix [][]int, r, c int) int {
	h := matrix[r][c]

	// top (from r - 1 to 0)
	top := 0

	for y := r - 1; y >= 0; y -= 1 {
		top += 1

		if matrix[y][c] >= h {
			break
		}
	}

	// bottom (from r+1 to len(matrix))
	bottom := 0
	for y := r + 1; y < len(matrix); y += 1 {
		bottom += 1

		if matrix[y][c] >= h {
			break
		}
	}

	// left (from c-1 to 0)
	left := 0
	for x := c - 1; x >= 0; x -= 1 {
		left += 1

		if matrix[r][x] >= h {
			break
		}
	}

	// right (from c+1 to len(matrix))
	right := 0
	for x := c + 1; x < len(matrix); x += 1 {
		right += 1

		if matrix[r][x] >= h {
			break
		}
	}

	return top * bottom * right * left
}
