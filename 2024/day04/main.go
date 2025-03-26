package day04

import (
	"log"
	"os"
	"strings"
)

func parseInput(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Incorrect path", path)
	}

	dataStr := string(data)
	dataStr = strings.TrimSpace(dataStr)
	return strings.Split(dataStr, "\n")
}

type Position struct {
	X, Y int
}

type Direction struct {
	Horizontal, Vertical int
	Position             Position
}

type Guess [2]int

var OmniGuesses []Guess = []Guess{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

var CrossGuesses []Guess = []Guess{
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func solveA(input []string) int {
	xPositions := getCharacterPositions(input, 'X')
	directions := getDirectionsFromPositions(input, 'M', xPositions, OmniGuesses)
	As := moveAlongTheRoad(input, 'A', directions)
	Ss := moveAlongTheRoad(input, 'S', As)
	return len(Ss)
}

func solveB(input []string) int {
	// aPositions := getCharacterPositions(input, 'A')
	// mDirections := getDirectionsFromPositions(input, 'M', aPositions, CrossGuesses)
	// sDirections := getDirectionsFromPositions(input, 'S', aPositions, CrossGuesses)
	// result := countCrosses(mDirections, sDirections)
	mPositions := getCharacterPositions(input, 'M')
	aDirections := getDirectionsFromPositions(input, 'A', mPositions, CrossGuesses)
	sDirections := moveAlongTheRoad(input, 'S', aDirections)

	return 0
}

func getCharacterPositions(input []string, char rune) []Position {
	result := make([]Position, 0)

	for yIndex, row := range input {
		for xIndex, value := range row {
			if value == char {
				result = append(result, Position{X: xIndex, Y: yIndex})
			}
		}
	}

	return result
}

func getDirectionsFromPositions(input []string, char rune, characterPositions []Position, guesses []Guess) []Direction {
	result := make([]Direction, 0)

	for _, position := range characterPositions {
		for _, guess := range guesses {
			newX := position.X + guess[0]
			newY := position.Y + guess[1]

			if newX < 0 || newX >= len(input[0]) {
				continue
			}

			if newY < 0 || newY >= len(input) {
				continue
			}

			guessedChar := input[newY][newX]

			if rune(guessedChar) == char {
				result = append(result, Direction{
					Horizontal: guess[0],
					Vertical:   guess[1],
					Position:   Position{X: newX, Y: newY},
				})
			}
		}
	}

	return result
}

func moveAlongTheRoad(input []string, char rune, directions []Direction) []Direction {
	result := make([]Direction, 0)

	for _, direction := range directions {
		newX := direction.Position.X + direction.Horizontal
		newY := direction.Position.Y + direction.Vertical

		if newX < 0 || newX >= len(input[0]) {
			continue
		}

		if newY < 0 || newY >= len(input) {
			continue
		}

		guessedChar := input[newY][newX]

		if rune(guessedChar) == char {
			result = append(result, Direction{
				Horizontal: direction.Horizontal,
				Vertical:   direction.Vertical,
				Position:   Position{X: newX, Y: newY},
			})
		}
	}

	return result
}

func countCrosses(mDirections, sDirections []Direction) int {
	result := 0

	for _, mDirection := range mDirections {
		guessedX := mDirection.Position.X + (mDirection.Horizontal * 2 * -1)
		guessedY := mDirection.Position.Y + (mDirection.Vertical * 2 * -1)

		for _, sDirection := range sDirections {
			if sDirection.Position.X == guessedX && sDirection.Position.Y == guessedY {
				log.Println("M Position:", mDirection.Position, "S Position:", sDirection.Position)
				result++
				break
			}
		}
	}

	return result
}

func countCrosses2(sDirections []Direction) int {
	result := 0

	for _, sDirection := range sDirections {

	}

	return result
}
