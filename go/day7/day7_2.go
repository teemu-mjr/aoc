package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	name     string
	parent   *Directory
	children []*Directory
	files    []File
	size     int
}

type File struct {
	name string
	size int
}

func main() {
	input := readInput(true)

	root := Directory{
		name: "/",
	}

	current := &root

	for _, v := range input {
		if v[0] == "$" {
			if v[1] == "cd" {
				if v[2] == ".." {
					current = current.parent
				} else {
					for i, child := range current.children {
						if child.name == v[2] {
							current = current.children[i]
						}
					}
				}
			}
		} else {
			if v[0] == "dir" {
				current.children = append(current.children, &Directory{
					name:   v[1],
					parent: current,
				})
			} else {
				fileSize, _ := strconv.ParseInt(v[0], 10, 64)
				current.files = append(current.files, File{
					name: v[1],
					size: int(fileSize),
				})
			}
		}
	}

	for {
		if current.parent == nil {
			break
		}
		current = current.parent
	}

	fmt.Println("COOL:", calcDir(current))
	sizes := getSizes(current)
	sort.Slice(sizes, func(a, b int) bool {
		return sizes[a] < sizes[b]
	})
	target := (30000000 - (70000000 - sizes[len(sizes)-1]))
	fmt.Println("Actual coolnsee:", coolNess(sizes, target))

}

func coolNess(sizes []int, target int) int {
	for _, size := range sizes {
		if size >= target {
			return size
		}
	}
	return 0
}

func getSizes(current *Directory) []int {
	var result []int

	result = append(result, current.size)

	for i := range current.children {
		result = append(result, getSizes(current.children[i])...)
	}

	return result
}

func calcDir(current *Directory) int {
	var total int
	for i := range current.files {
		current.size += current.files[i].size
	}

	for j := range current.children {
		total += calcDir(current.children[j])
		current.size += current.children[j].size
	}

	if current.size < 100_000 {
		total += current.size
	}

	return total
}

func readInput(prod bool) [][]string {
	var result [][]string
	var file *os.File
	var err error
	if prod == true {
		file, err = os.Open("input.prod")
	} else {
		file, err = os.Open("input.test")
	}

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		result = append(result, strings.Split(fileScanner.Text(), " "))
	}

	file.Close()

	return result
}
