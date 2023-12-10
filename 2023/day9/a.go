package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isZeroSequence(sequence []int) bool {
	for _, value := range sequence {
		if (value != 0) {
			return false
		}
	}

	return true
}

func getAdditionalSequences(originalSequence []int) [][]int {
	allSequences := make([][]int, 1)
	allSequences[0] = originalSequence

	for ;; {
		currentSequence := allSequences[len(allSequences) - 1]
		if (isZeroSequence(currentSequence)) {
			break
		}

		newSequence := make([]int, 0)

		for i := 0; i < len(currentSequence) - 1; i++ {
			newValue := currentSequence[i + 1] - currentSequence[i]
			newSequence = append(newSequence, newValue)
		}

		allSequences = append(allSequences, newSequence)
	}	

	return allSequences
}

func addNewValuesToAllSequences(allSequences [][]int) [][]int {
	allSequencesLength := len(allSequences)
	allSequences[allSequencesLength - 1] = append(allSequences[allSequencesLength - 1], 0)

	for i := allSequencesLength - 2; i >= 0; i-- {
		newValue := allSequences[i][len(allSequences[i]) - 1] + allSequences[i + 1][len(allSequences[i + 1]) - 1]
		allSequences[i] = append(allSequences[i], newValue)
	}

	return allSequences
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	sum := 0

	for sc.Scan() {
		line := sc.Text()
		sequence := make([]int, 0)

		sequenceStr := strings.Split(line, " ")

		for _, valueStr := range sequenceStr {
			value, err := strconv.Atoi(valueStr)

			if (err != nil) {
				fmt.Println("Value could not be converted to int", valueStr)
				return
			}

			sequence = append(sequence, value)
		}

		allSequences := getAdditionalSequences(sequence)
		predictedSequences := addNewValuesToAllSequences(allSequences)
		newValue := predictedSequences[0][len(predictedSequences[0]) - 1]
		sum += newValue
	}

	fmt.Println("Result: ", sum)
}
