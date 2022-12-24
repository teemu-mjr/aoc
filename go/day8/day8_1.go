package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput(true)
	var forest [][]int
	for _, row := range input {
		var treeRow []int
		for _, r := range row {
			tree := int(r - '0')
			treeRow = append(treeRow, tree)
		}
		forest = append(forest, treeRow)
	}

	var result int

	for i := range forest {
		for j := range forest[i] {
			if i <= 0 || i+1 >= len(forest) || j <= 0 || j+1 >= len(forest[i]) {
				fmt.Print("\033[32m", forest[i][j], "\033[0m")
				result++
				continue
			} else if canSee(forest, i, j) {
				fmt.Print("\033[35m", forest[i][j], "\033[0m")
				result++
				continue
			} else {
				fmt.Print(forest[i][j])
			}

		}
		fmt.Print("\n")
	}

	fmt.Println("RESULT: ", result)
}

func canSee(forest [][]int, i int, j int) bool {
	// left
	isGood := true
	for n := j; n > 0; n-- {
		if forest[i][n-1] >= forest[i][j] {
			isGood = false
			break
		}
	}
	if isGood {
		return true
	}

	// right
	isGood = true
	for n := j; n < len(forest[i])-1; n++ {
		if forest[i][n+1] >= forest[i][j] {
			isGood = false
			break
		}
	}
	if isGood {
		return true
	}

	// up
	isGood = true
	for n := i; n > 0; n-- {
		if forest[n-1][j] >= forest[i][j] {
			isGood = false
			break
		}
	}
	if isGood {
		return true
	}

	// down
	isGood = true
	for n := i; n < len(forest)-1; n++ {
		if forest[n+1][j] >= forest[i][j] {
			isGood = false
			break
		}
	}
	if isGood {
		return true
	}

	return false
}

func readInput(prod bool) []string {
	var result []string
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
		result = append(result, fileScanner.Text())
	}

	file.Close()

	return result
}
