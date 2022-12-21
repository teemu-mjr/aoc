package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput(true)
	var readHead []int
	for i := 0; i < 14; i++ {
		readHead = append(readHead, i)
	}

	fmt.Println(readNext(readHead, input))
}

func readNext(readHead []int, input []string) int {
	for _, i := range readHead {
		for _, j := range readHead {
			if i == j {
				continue
			}
			if input[i] == input[j] {
				readHead = moveReadHead(readHead)
				readNext(readHead, input)
				break
			}
		}
	}
	return readHead[len(readHead)-1] + 1
}

func moveReadHead(readHead []int) []int {
	for i, v := range readHead {
		readHead[i] = v + 1
	}
	return readHead
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

	fileScanner.Split(bufio.ScanBytes)

	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	file.Close()

	return result
}
