package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func transpose(matrix [][]rune) [][]rune {
	trasposedMatrix := make([][]rune, 0)

	for x := 0; x < len(matrix[0]); x++ {
		column := make([]rune, len(matrix))

		for y := 0; y < len(matrix); y++ {
			column[y] = matrix[y][x]
		}

		trasposedMatrix = append(trasposedMatrix, column)
	}

	return trasposedMatrix
}

func expandUniverseHorizontally(universe [][]rune) [][]rune {
	newUniverse := make([][]rune, 0)

	for _, row := range universe {
		str := string(row)
		
		if (strings.Contains(str, "#")) {
			newUniverse = append(newUniverse, row)
		} else {
			newUniverse = append(newUniverse, row, row)
		}
	}

	return newUniverse
}

func expandUniverse(universe [][]rune) [][]rune {
	universe = expandUniverseHorizontally(universe)
	universe = transpose(universe)
	universe = expandUniverseHorizontally(universe)
	universe = transpose(universe)

	return universe
}

func findGalaxies(universe [][]rune) []Pos {
	galaxies := make([]Pos, 0)

	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[0]); x++ {
			currentChar := universe[y][x]

			if currentChar == '#' {
				newGalaxy := Pos{ x, y }
				galaxies = append(galaxies, newGalaxy)
			}
		}
	}

	return galaxies
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func calculateDistance(galaxies []Pos) int {
	result := 0

	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		rest := galaxies[i+1:]

		for _, destination := range rest {
			result += abs(destination.x - galaxy.x) + abs(destination.y - galaxy.y)
		}
	}

	return result
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	universe := make([][]rune, 0)

	for sc.Scan() {
		line := sc.Text()

		row := make([]rune, 0)
		
		for _, char := range line {
			row = append(row, char)
		}

		universe = append(universe, row)
	}

	universe = expandUniverse(universe)
	galaxies := findGalaxies(universe)
	distance := calculateDistance(galaxies)
	fmt.Println(distance)
}
