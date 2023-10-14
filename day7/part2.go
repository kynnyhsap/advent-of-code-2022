package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name   string
	parent *Directory

	directories []*Directory
	files       []*File
}

const FREE_SPACE_REQUIRED = 30000000
const TOTAL_SPACE = 70000000

func dirSize(dir *Directory) int {
	size := 0

	for _, childFile := range dir.files {
		size += childFile.size
	}

	for _, childDir := range dir.directories {
		size += dirSize(childDir)
	}

	return size
}

func calcSmallestDirToBeDeleted(dir *Directory, rootDirSize int, smallestDirToBeDeleted *int) int {
	size := 0

	for _, childFile := range dir.files {
		size += childFile.size
	}

	for _, childDir := range dir.directories {
		size += calcSmallestDirToBeDeleted(childDir, rootDirSize, smallestDirToBeDeleted)
	}

	if rootDirSize-size < FREE_SPACE_REQUIRED && size < *smallestDirToBeDeleted {
		// fmt.Printf("%s: 70000000 - %d = %d is less than 30000000\n", dir.name, size, TOTAL_SPACE-size)
		*smallestDirToBeDeleted = size
	}

	return size
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := &Directory{
		name:   "/",
		parent: nil,
	}

	currentDir := root

	for scanner.Scan() {
		line := scanner.Text()

		// cd
		if line[0] == '$' && line[2] == 'c' {

			if line[5] == '/' {
				// switch to root dir
				currentDir = root
			} else if line[5] == '.' {
				// move out
				currentDir = currentDir.parent
			} else {
				// move in
				for _, dir := range currentDir.directories {
					if strings.Compare(dir.name, string(line[5:])) == 0 {
						currentDir = dir
						break
					}
				}
			}

		}

		// listing
		if line[0] != '$' {
			if line[0] == 'd' {
				currentDir.directories = append(currentDir.directories, &Directory{
					name:   string(line[4:]),
					parent: currentDir,
				})
			} else {
				splitted := strings.Split(string(line), " ")

				size, _ := strconv.Atoi(splitted[0])

				currentDir.files = append(currentDir.files, &File{
					size: size,
					name: splitted[1],
				})
			}
		}
	}

	rootSize := dirSize(root)
	gap := FREE_SPACE_REQUIRED - (TOTAL_SPACE - rootSize)

	smallest := rootSize

	traverse(root, gap, &smallest)

	fmt.Println(smallest)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func traverse(dir *Directory, gap int, smallest *int) {
	size := dirSize(dir)

	if size >= gap {
		if size < *smallest {
			*smallest = size
		}
	}

	for _, childDir := range dir.directories {
		traverse(childDir, gap, smallest)
	}

}
