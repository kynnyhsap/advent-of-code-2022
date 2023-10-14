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

const SIZE_LIMIT = 100000

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

	// walkTree(root)

	sum := 0
	dirSize(root, &sum)
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func dirSize(dir *Directory, sum *int) int {
	size := 0

	for _, childFile := range dir.files {
		size += childFile.size
	}

	for _, childDir := range dir.directories {
		size += dirSize(childDir, sum)
	}

	if size < SIZE_LIMIT {
		fmt.Printf("add dir %s with size %d to sum\n", dir.name, size)
		*sum += size
	}

	return size
}

func walkTree(root *Directory) {
	visitDir(root, 0)
}

func visitDir(dir *Directory, level int) {
	indent := strings.Repeat("  ", level)
	fmt.Printf(indent+"- %s (dir)\n", dir.name)

	for _, childDir := range dir.directories {
		visitDir(childDir, level+1)
	}

	for _, childFile := range dir.files {
		visitFile(childFile, level+1)
	}
}

func visitFile(file *File, level int) {
	indent := strings.Repeat("  ", level)
	fmt.Printf(indent+"- %s (file, size=%d)\n", file.name, file.size)
}
