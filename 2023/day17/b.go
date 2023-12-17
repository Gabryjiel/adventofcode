// taken from "https://github.com/keriati/aoc/blob/master/2023/day17.ts"
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	hashset "github.com/ugurcsen/gods-generic/sets/hashset"
	heap "github.com/ugurcsen/gods-generic/trees/binaryheap"
)

type City [][]int

func main() {
	content, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	city := makeCity(content)
	heatLoss := getLeastHeatLoss(city, 10, 4)

	fmt.Println("Result:", heatLoss)
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

func makeCity(content []string) City {
	city := make(City, len(content))

	for rowIndex, row := range content {
		city[rowIndex] = make([]int, len(content[0]))

		for colIndex, char := range row {
			value, err := strconv.Atoi(string(char))

			if err != nil {
				value = 0
			}

			city[rowIndex][colIndex] = value
		}
	}

	return city
}

type Position struct {
	x, y int
	direction int
}

type Step struct {
	heuristic int
	position Position
	heatLoss int
	steps int
}

func getLeastHeatLoss(city City, maxStepsInDirection, minStepsInDirection int) int {
	endX, endY := len(city[0]) - 1, len(city) - 1
	startX, startY := 0, 0

	getNextPositionsMap := map[int](func (x, y int) []Position){
		1: func (x, y int) []Position {
			positions := make([]Position, 3)

			positions[0] = Position{
				x: x,
				y: y - 1,
				direction: 1,
			}

			positions[1] = Position{
				x: x + 1,
				y: y,
				direction: 3,
			}

			positions[2] = Position{
				x: x - 1,
				y: y,
				direction: 4,
			}

			return positions
		},
		2: func (x, y int) []Position {
			positions := make([]Position, 3)

			positions[0] = Position{
				x: x,
				y: y + 1,
				direction: 2,
			}

			positions[1] = Position{
				x: x + 1,
				y: y,
				direction: 3,
			}

			positions[2] = Position{
				x: x - 1,
				y: y,
				direction: 4,
			}

			return positions
		},
		3: func (x, y int) []Position {
			positions := make([]Position, 3)

			positions[0] = Position{
				x: x,
				y: y + 1,
				direction: 2,
			}

			positions[1] = Position{
				x: x,
				y: y - 1,
				direction: 1,
			}

			positions[2] = Position{
				x: x + 1,
				y: y,
				direction: 3,
			}

			return positions
		},
		4: func (x, y int) []Position {
			positions := make([]Position, 3)

			positions[0] = Position{
				x: x,
				y: y + 1,
				direction: 2,
			}

			positions[1] = Position{
				x: x,
				y: y - 1,
				direction: 1,
			}

			positions[2] = Position{
				x: x - 1,
				y: y,
				direction: 4,
			}

			return positions
		},
	}

	startingStepEast := Step{
		heuristic: 0,
		position: Position{
			x: startX,
			y: startY,
			direction: 3,
		},
		heatLoss: 0,
		steps: 0,
	}

	startingStepSouth := Step{
		heuristic: 0,
		position: Position{
			x: startX,
			y: startY,
			direction: 2,
		},
		heatLoss: 0,
		steps: 0,
	}

	comparator := func(a, b Step) int {
		return a.heuristic - b.heuristic
	}
	heap := heap.NewWith(comparator)
	heap.Push(startingStepEast)
	heap.Push(startingStepSouth)

	visited := hashset.New[int]() 
	visited.Add(cacheKey(startingStepEast))
	visited.Add(cacheKey(startingStepSouth))

	for heap.Size() > 0 {
		value, ok := heap.Pop()

		if ok == false {
			break
		}

		if value.position.x == endX && value.position.y == endY {
			if value.steps < minStepsInDirection {
				continue
			}

			return value.heatLoss
		}

		nextPositionsFunc, ok := getNextPositionsMap[value.position.direction]

		if ok == false {
			continue
		}

		potentialNextPositions := nextPositionsFunc(value.position.x, value.position.y)
		nextPositions := make([]Position, 0)
		for _, potentialPosition := range potentialNextPositions {
			if value.steps < minStepsInDirection {
				if potentialPosition.direction != value.position.direction {
					continue
				}
			} else if value.steps >= maxStepsInDirection {
				if potentialPosition.direction == value.position.direction {
					continue
				}
			}

			if potentialPosition.x < 0 || 
				potentialPosition.y < 0 || 
				potentialPosition.y > endY || 
				potentialPosition.x > endX {
				continue
			}

			nextPositions = append(nextPositions, potentialPosition)
		}

		for _, nextPosition := range nextPositions {
			newSteps := 1

			if nextPosition.direction == value.position.direction {
				newSteps = value.steps + 1
			}

			nextStep := Step{
				heuristic: value.heatLoss + city[nextPosition.y][nextPosition.x] + endX - nextPosition.x + endY - nextPosition.y,
				position: nextPosition,
				heatLoss: value.heatLoss + city[nextPosition.y][nextPosition.x],
				steps: newSteps,
			}

			key := cacheKey(nextStep)

			if !visited.Contains(key) {
				visited.Add(key)
				heap.Push(nextStep)
			}
		}
	}

	return -1
}

func cacheKey(step Step) int {
	return (step.position.y << 16) | 
		(step.position.x << 8) | 
		(step.position.direction << 4) | 
		step.steps
}
