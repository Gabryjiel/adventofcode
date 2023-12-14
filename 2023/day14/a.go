package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	platform := make([][]rune, 0)

	for sc.Scan() {
		line := sc.Text()
		platform = append(platform, []rune(line))
	}

	moveRocks(platform)

	result := calculateLoad(platform)
	fmt.Println("Result:", result)
}

func moveRocks(platform [][]rune) {
	for y := 0; y < len(platform); y++ {
		for x := 0; x < len(platform[0]); x++ {
			char := platform[y][x]

			if char == 'O' {
				moveNorth(platform, y, x, char)
			}
		}
	}
}

func moveNorth(platform [][]rune, y, x int, char rune) {
	for i := y - 1; i >= 0; i-- {
		charToSwap := platform[i][x]

		if charToSwap == '.' {
			platform[i][x] = char
			platform[i + 1][x] = charToSwap
		} else {
			break
		}
	}
}

func calculateLoad(platform [][]rune) int {
	sum := 0

	for index, line := range platform {
		str := string(line)
		sum += strings.Count(str, "O") * (len(platform) - index)
	}

	return sum
}
