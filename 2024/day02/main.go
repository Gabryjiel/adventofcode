package main

import (
	"log"
	"os"
	"slices"
)

func main() {
	input := parseInput("input.txt")

	safeCount := task1(input)
	log.Println("Result 1:", safeCount)

	safeCount = task2(input)
	log.Println("Result 2:", safeCount)
}

type Level int
type Report []Level

func parseInput(filename string) []Report {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]Report, 0)
	helper := make(Report, 0)
	var value Level = 0

	for _, part := range input {
		if part == '\n' {
			helper = append(helper, value)
			value = 0
			result = append(result, slices.Clone(helper))
			helper = helper[:0]
		} else if part == ' ' {
			helper = append(helper, value)
			value = 0
		} else {
			value = value*10 + Level(part-'0')
		}
	}

	return result
}

func (report Report) isSafe() bool {
	if len(report) < 2 {
		return true
	}

	monotony := report[0] - report[1]
	previous := report[0]

	for i := 1; i < len(report); i++ {
		current := report[i]
		difference := previous - current

		isIdenticalAsPrevious := difference == 0
		hasChangedSign := difference*monotony < 0
		notInRequiredRange := !((difference <= 3 && difference >= 1) || (difference >= -3 && difference <= -1))

		if isIdenticalAsPrevious || hasChangedSign || notInRequiredRange {
			return false
		}

		previous = current
	}

	return true
}

func (report Report) isSafeDampended() bool {
	if len(report) < 2 {
		return true
	}

	monotony := report[0] - report[1]
	previous := report[0]
	isOnSecondChance := false

	if len(report) > 2 {
		monotony2 := report[1] - report[2]
		monotony3 := report[2] - report[3]

		if monotony*monotony2 < 0 && monotony2*monotony3 > 0 {
			return report[1:].isSafe()
		}
	}

	for i := 1; i < len(report); i++ {
		current := report[i]
		difference := previous - current

		isIdenticalAsPrevious := difference == 0
		hasChangedSign := difference*monotony < 0
		notInRequiredRange := !((difference <= 3 && difference >= 1) || (difference >= -3 && difference <= -1))

		if isIdenticalAsPrevious || hasChangedSign || notInRequiredRange {
			if isOnSecondChance {
				return false
			}

			isOnSecondChance = true
			continue
		}

		previous = current
	}

	return true
}

func task1(input []Report) int {
	safeCount := 0

	for _, report := range input {
		if report.isSafe() {
			safeCount++
		}
	}

	return safeCount
}

func task2(input []Report) int {
	safeCount := 0

	for _, report := range input {
		if report.isSafeDampended() {
			safeCount++
		}
	}

	return safeCount
}
