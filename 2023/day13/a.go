package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	puzzles := make([][]string, 0)
	puzzle := make([]string, 0)
	for sc.Scan() {
		line := sc.Text()
		
		if line == "" {
			puzzles = append(puzzles, puzzle)
			puzzle = make([]string, 0)
		} else {
			puzzle = append(puzzle, line)
		}
	}
	puzzles = append(puzzles, puzzle)

	sum := 0

	for _, singlePuzzle := range puzzles {
		sum += checkPuzzle(singlePuzzle) * 100
		sum += checkPuzzle(transpose(singlePuzzle))
	}

	fmt.Println("Result:", sum)
}

func checkIfMirror(pattern []string) bool {
	if len(pattern) % 2 != 0 {
		return false
	}

	for y := 0; y < len(pattern); y++ {
		if pattern[y] != pattern[len(pattern) - 1 - y] {
			return false
		}
	}

	return true 
}

func takeColumnFromPuzzle(puzzle []string, n int) string {
	result := ""

	if n > len(puzzle[0]) {
		return result
	}

	for _, line := range puzzle {
		result = result + string(line[n])
	}

	return result
}

func transpose(puzzle []string) []string {
	result := make([]string, 0)

	for x := 0; x < len(puzzle[0]); x++ {
		result = append(result, takeColumnFromPuzzle(puzzle, x))
	}

	return result
}

func checkPuzzle(puzzle []string) int {
	linesFromTop := make([]string, 0)
	linesFromBot := make([]string, 0)
	for y := 0; y < len(puzzle); y++ {
		linesFromTop = append(linesFromTop, puzzle[y])
		linesFromBot = append(linesFromBot, puzzle[len(puzzle) - 1 - y])

		if checkIfMirror(linesFromTop) {
			return (y + 1) / 2
		} else if checkIfMirror(linesFromBot) {
			return len(puzzle) - ((y + 1) / 2)
		}
	}

	return 0
}
