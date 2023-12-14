package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	platform = cycle(platform, 1_000_000_000)
	result := calculateLoad(platform)
	fmt.Println("Result:", result)
}

func cycle(platform [][]rune, times int) [][]rune {
	platforms := make([]string, 0)

	for i := 0; i < times; i++ {
		moveRocksNorth(platform)
		moveRocksWest(platform)
		moveRocksSouth(platform)
		moveRocksEast(platform)

		hash := makeHash(platform)
		isIn := slices.Contains(platforms, hash)

		if isIn == false {
			platforms = append(platforms, hash)
		} else {
			isInIndex := slices.Index(platforms, hash)
			cycleLength := i - isInIndex
			rest := (times - isInIndex) % cycleLength
			platform = makePlatform(platforms[isInIndex + rest - 1])
			return platform
		}
	}

	return makePlatform(platforms[len(platforms) - 1])
}

func makeHash(platform [][]rune) string {
	lines := make([]string, len(platform))

	for index, row := range platform {
		lines[index] = string(row)
	}

	return strings.Join(lines, "?")
}

func makePlatform(hash string) [][]rune {
	lines := strings.Split(hash, "?")
	platform := make([][]rune, len(lines))

	for index, line := range lines {
		platform[index] = []rune(line)
	}
	

	return platform
}


func calculateLoad(platform [][]rune) int {
	sum := 0

	for index, line := range platform {
		str := string(line)
		sum += strings.Count(str, "O") * (len(platform) - index)
	}

	return sum
}
func moveRocksNorth(platform [][]rune) {
	for y := 0; y < len(platform); y++ {
		for x := 0; x < len(platform[0]); x++ {
			char := platform[y][x]

			if char != 'O' {
				continue
			}

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
	}
}

func moveRocksSouth(platform [][]rune) {
	for y := len(platform) - 1; y >= 0; y-- {
		for x := 0; x < len(platform[0]); x++ {
			char := platform[y][x]

			if char != 'O' {
				continue
			}

			for i := y + 1; i < len(platform); i++ {
				charToSwap := platform[i][x]

				if charToSwap == '.' {
					platform[i][x] = char
					platform[i - 1][x] = charToSwap
				} else {
					break
				}
			}
		}
	}
}
func moveRocksEast(platform [][]rune) {
	for x := len(platform[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(platform); y++ {
			char := platform[y][x]

			if char != 'O' {
				continue
			}

			for i := x + 1; i < len(platform[0]); i++ {
				charToSwap := platform[y][i]

				if charToSwap == '.' {
					platform[y][i] = char
					platform[y][i - 1] = charToSwap
				} else {
					break
				}
			}
		}
	}
}
func moveRocksWest(platform [][]rune) {
	for x := 0; x < len(platform[0]); x++ {
		for y := 0; y < len(platform); y++ {
			char := platform[y][x]

			if char != 'O' {
				continue
			}

			for i := x - 1; i >= 0; i-- {
				charToSwap := platform[y][i]

				if charToSwap == '.' {
					platform[y][i] = char
					platform[y][i + 1] = charToSwap
				} else {
					break
				}
			}
		}
	}
}
