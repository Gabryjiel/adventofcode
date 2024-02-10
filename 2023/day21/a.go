package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	timeStart := time.Now()
	content, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	tiles := parse(content)
	result := solve(tiles, 26501365)

	fmt.Println("Result:", result)
	fmt.Println("Time:", float64(time.Since(timeStart).Microseconds()) / 1000, "ms")
}

func openFile(name string) ([]string, error) {
	file, err := os.Open(name)	
	defer file.Close()

	content := make([]string, 0)

	if (err != nil) {
		return nil, fmt.Errorf("File %s not found", name)
	}
	
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		content = append(content, sc.Text())
	}

	return content, nil
}

const (
	START = iota
	PLOT
	ROCK
)

type Position struct {
	x, y int
}

type Tile struct {
	position Position
	content int
}

func parse(content []string) ([][]Tile) {
	result := make([][]Tile, len(content))

	for i := 0; i < len(result); i++ {
		result[i] = make([]Tile, len(content[0]))
	}

	for rowIndex, row := range content {
		for colIndex, col := range row {
			result[rowIndex][colIndex] = Tile{
				content: getTileType(col),
				position: Position{
					x: colIndex,
					y: rowIndex,
				},
			}
		}
	}

	return result
}

func getTileType(char rune) int {
	switch char {
	case 'S':
		return START
	case '.':
		return PLOT
	case '#':
		return ROCK
	default:
		return ROCK
	}
}

type Step struct {
	iteration int
	position Position
}

func solve(tiles [][]Tile, maxSteps int) int{
	startingTile := getStartingTile(tiles)

	steps := make([]Step, 1)
	steps[0] = Step{
		iteration: 0,
		position: startingTile.position,
	}

	positionAdd := [][]int{ {1, 0}, {0, 1}, {-1, 0}, {0, -1} }
	maxX, maxY := len(tiles[0]) - 1, len(tiles) - 1

	for i := 0; i < maxSteps; i++ {
		thisIteration := steps
		steps = []Step{}

		for _, step := range thisIteration {

			for _, add := range positionAdd {
				newX := step.position.x + add[0]
				newY := step.position.y + add[1]

				if newX < 0 || newX > maxX || 
					newY < 0 || newY > maxY {
					continue
				}

				newTile := tiles[newY][newX]

				if newTile.content == ROCK {
					continue	
				}

				newStep := Step{
					iteration: i,
					position: Position{
						x: newX,
						y: newY, 
					},
				}

				steps = addUnique(steps, newStep)
			}
		}
	}

	return len(steps)
}

func addUnique(steps []Step, newStep Step) []Step {
	for _, step := range steps {
		if step.position.y == newStep.position.y &&
			step.position.x == newStep.position.x {
			return steps
		}
	}

	return append(steps, newStep)
}

func getStartingTile(tiles [][]Tile) Tile {
	for _, row := range tiles {
		for _, col := range row {
			if col.content == START {
				return col
			}
		}
	}

	return Tile{}
}
