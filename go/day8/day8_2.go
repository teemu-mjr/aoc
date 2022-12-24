package main

import (
	"bufio"
	"fmt"
	"os"
)

type ViewDistance struct {
	left  int
	right int
	up    int
	down  int
}

func (v ViewDistance) total() int {
	return v.left * v.right * v.up * v.down
}

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

	bestView := 0

	for i := range forest {
		for j := range forest[i] {
			treeView := viewDistance(forest, i, j)
			if treeView.total() > bestView {
				bestView = treeView.total()
			}
		}
	}

	fmt.Println("BEST: ", bestView)
}

func viewDistance(forest [][]int, i int, j int) ViewDistance {
	var viewDistance ViewDistance
	// left
	for n := j; n > 0; n-- {
		viewDistance.left++
		if forest[i][n-1] >= forest[i][j] {
			break
		}
	}

	// right
	for n := j; n < len(forest[i])-1; n++ {
		viewDistance.right++
		if forest[i][n+1] >= forest[i][j] {
			break
		}
	}

	// up
	for n := i; n > 0; n-- {
		viewDistance.up++
		if forest[n-1][j] >= forest[i][j] {
			break
		}
	}

	// down
	for n := i; n < len(forest)-1; n++ {
		viewDistance.down++
		if forest[n+1][j] >= forest[i][j] {
			break
		}
	}

	return viewDistance
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
