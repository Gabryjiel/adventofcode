package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("File not found")
	}

	defer input.Close()

	sc := bufio.NewScanner(input)

	leftTable := make([]int, 0)
	rightTable := make([]int, 0)

	for sc.Scan() {
		text := sc.Text()
		text = strings.ReplaceAll(text, "   ", " ")

		split := strings.Split(text, " ")

		leftNum, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalln("Failed conversion with", split[0])
		}

		rightNum, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalln("Failed conversion with", split[1])
		}

		leftTable = append(leftTable, leftNum)
		rightTable = append(rightTable, rightNum)
	}

	slices.Sort(leftTable)
	slices.Sort(rightTable)

	result := calculateDistances(leftTable, rightTable)
	fmt.Println("Result:", result)
}

func calculateDistances(left, right []int) int {
	result := 0

	for index, num := range left {
		diff := num - right[index]

		if diff < 0 {
			result -= diff
		} else {
			result += diff
		}
	}

	return result
}
