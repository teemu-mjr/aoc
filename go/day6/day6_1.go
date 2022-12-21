package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput(true)
	readHead := []int{0, 1, 2, 3}

	readNext(readHead, input)
}

func readNext(readHead []int, input []string) {
    fmt.Println(readHead)
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
	for _, v := range readHead {
		fmt.Print(input[v])
	}
	fmt.Println(readHead[len(readHead)-1] + 1)
    return
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
