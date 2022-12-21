package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Box struct {
	index int
	stack []string
}

func tryParseUint(str string) int {
	res, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	return int(res)
}

func parseInstruction(instruction string) [3]int {
	var result [3]int

	words := strings.Split(instruction, " ")

	result[0] = tryParseUint(words[1])
	result[1] = tryParseUint(words[3])
	result[2] = tryParseUint(words[5])

	return result
}

func parseBoxLines(lines []string) [16][]string {
	var result [16][]string

	for _, line := range lines {
		for i, char := range line {
			if (i+3)%4 != 0 || char == 32 || unicode.IsNumber(char) {
				continue
			}
			result[i/4] = append(result[i/4], strconv.QuoteRune(char))
		}
	}

	for _, arr := range result {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return result
}

func main() {
	readFile, err := os.Open("input.prod")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var boxLines []string

	var readInstruct bool = false
	var instructions [][3]int

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			readInstruct = true
			continue
		}

		if !readInstruct {
			boxLines = append(boxLines, fileScanner.Text())
		}

		if readInstruct {
			instructions = append(instructions, parseInstruction(fileScanner.Text()))
		}

	}
	readFile.Close()

	boxStacks := parseBoxLines(boxLines)

	for _, instruction := range instructions {
		for i := 0; i < instruction[0]; i++ {
			moveOne(&boxStacks[instruction[1]-1], &boxStacks[instruction[2]-1])
		}
	}

	var result string

	for _, stack := range boxStacks {
		if len(stack) > 0 {
			result = result + strings.Split(stack[len(stack)-1], "'")[1]
		}
	}

	fmt.Println(result)

}
func moveOne(a *[]string, b *[]string) {
	newA := *a
	*b = append(*b, newA[len(newA)-1])
	newA = newA[:len(newA)-1]
	*a = newA
}
