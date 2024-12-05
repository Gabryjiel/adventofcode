package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	input := parseInput("input.txt")

	result := task1(input)
	log.Println("Task 1 result:", result)

	result = task2(input)
	log.Println("Task 2 result:", result)

}

func parseInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("File not found", path)
	}

	return string(data)
}

func task1(input string) int {
	result := 0
	target := "mul(,)"
	cursor := 0
	X := 0
	Y := 0

	for _, char := range input {
		if byte(char) == target[cursor] {
			if target[cursor] == ')' {
				result += X * Y
				X = 0
				Y = 0
				cursor = 0
				continue
			}

			cursor++
		} else if cursor == 4 && (char >= '0' && char <= '9') {
			num, _ := strconv.Atoi(string(char))
			X = X*10 + num
		} else if cursor == 5 && (char >= '0' && char <= '9') {
			num, _ := strconv.Atoi(string(char))
			Y = Y*10 + num
		} else {
			cursor = 0
			X = 0
			Y = 0
		}
	}

	return result
}

type State struct {
	IsEnabled bool
	Sum       int
}

type Needle struct {
	Pattern      string
	Action       func(state State, needle Needle) State
	Cursor       int
	HelperValues []int
}

func (needle *Needle) Reset() {
	needle.Cursor = 0
	needle.HelperValues = make([]int, 2)
	needle.HelperValues[0] = 0
	needle.HelperValues[1] = 0
}

func (needle *Needle) CheckIfCharEquals(char byte) bool {
	if needle.Pattern[needle.Cursor] == char {
		return true
	} else if needle.Pattern[needle.Cursor] == 'X' {
		if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			needle.HelperValues[0] = needle.HelperValues[0]*10 + num
			return true
		} else {
			needle.Cursor++
			return needle.CheckIfCharEquals(char)
		}
	} else if needle.Pattern[needle.Cursor] == 'Y' {
		if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			needle.HelperValues[1] = needle.HelperValues[1]*10 + num
			return true
		} else {
			needle.Cursor++
			return needle.CheckIfCharEquals(char)
		}
	}

	return false
}

func task2(input string) int {
	state := State{
		Sum:       0,
		IsEnabled: true,
	}

	needles := [...]Needle{
		Needle{
			Pattern: "do()",
			Action: func(state State, needle Needle) State {
				state.IsEnabled = true
				return state
			},
		},
		Needle{
			Pattern: "don't()",
			Action: func(state State, needle Needle) State {
				state.IsEnabled = false
				return state
			},
		},
		Needle{
			Pattern: "mul(XXX,YYY)",
			Action: func(state State, needle Needle) State {
				if state.IsEnabled && len(needle.HelperValues) == 2 {
					state.Sum += needle.HelperValues[0] * needle.HelperValues[1]
				}

				return state
			},
			HelperValues: make([]int, 2),
		},
	}

	for _, char := range input {
		charByte := byte(char)

		for n := 0; n < len(needles); n++ {
			if needles[n].CheckIfCharEquals(charByte) {
				needles[n].Cursor++

				if needles[n].Cursor >= len(needles[n].Pattern) {
					state = needles[n].Action(state, needles[n])
					needles[n].Reset()
				}
			} else {
				needles[n].Reset()
			}
		}
	}

	return state.Sum
}
