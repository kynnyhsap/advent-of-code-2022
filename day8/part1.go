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

	totalVisibleTrees := 0

	for r := range matrix {
		for c := range matrix[r] {
			if visible(matrix, r, c) {
				totalVisibleTrees += 1
			}
		}
	}

	fmt.Println(totalVisibleTrees)
}

func visible(matrix [][]int, r, c int) bool {
	if c == 0 || r == 0 || r == len(matrix)-1 || c == len(matrix)-1 {
		return true
	}

	// top (from 0 to r)
	visibleFromTop := true
	for y := 0; y < r; y += 1 {
		if matrix[y][c] >= matrix[r][c] {
			visibleFromTop = false
			break

		}
	}

	// bottom (from r+1 to len(matrix))
	visibleFromBottom := true
	for y := r + 1; y < len(matrix); y += 1 {
		if matrix[y][c] >= matrix[r][c] {
			visibleFromBottom = false
			break
		}
	}

	// left (from 0 to c)
	visibleFromLeft := true
	for x := 0; x < c; x += 1 {
		if matrix[r][x] >= matrix[r][c] {
			visibleFromLeft = false
			break
		}
	}

	// right (from c+1 to len(matrix))
	visibleFromRight := true
	for x := c + 1; x < len(matrix); x += 1 {
		if matrix[r][x] >= matrix[r][c] {
			visibleFromRight = false
			break
		}
	}

	return visibleFromTop || visibleFromBottom || visibleFromLeft || visibleFromRight
}
