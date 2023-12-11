package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Pos struct {
	x, y int
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func calculateDistanceInALine(start, end int, line []int) int {
	result := 0

	startM := start
	endM := end

	if startM > endM {
		startM = end
		endM = start
	}

	for x := startM; x < endM; x++ {
		isNotExpandedRow := slices.Contains(line, x)

		if isNotExpandedRow {
			result += 1
		} else {
			result += 1_000_000
		}
	}

	return result
}

func calculateDistance(galaxies []Pos, rows, cols []int) int {
	result := 0

	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		rest := galaxies[i+1:]

		for _, destination := range rest {
			result += calculateDistanceInALine(galaxy.x, destination.x, rows)
			result += calculateDistanceInALine(galaxy.y, destination.y, cols)
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
	galaxies := make([]Pos, 0)

	rows := make([]int, 0)
	cols := make([]int, 0)

	for sc.Scan() {
		line := sc.Text()

		row := make([]rune, 0)
		
		for _, char := range line {
			row = append(row, char)

			if char == '#' {
				x := len(row)
				y := len(universe)

				newGalaxy := Pos{ x, y }
				galaxies = append(galaxies, newGalaxy)
				rows = append(rows, x)
				cols = append(cols, y)
			}
		}

		universe = append(universe, row)
	}

	distance := calculateDistance(galaxies, rows, cols)
	fmt.Println(distance)
}
