package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("example.txt")
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

	occurencesMap := makeMap(rightTable)
	result := calculateSimilarityScore(leftTable, occurencesMap)

	log.Println("Result:", result)
}

func calculateSimilarityScore(leftTable []int, occurencesMap map[int]int) int {
	result := 0

	for _, value := range leftTable {
		occurences, ok := occurencesMap[value]

		if ok {
			result += occurences * value
		}
	}

	return result
}

func makeMap(table []int) map[int]int {
	result := make(map[int]int)

	for _, value := range table {
		mapValue, ok := result[value]

		if ok {
			result[value] = mapValue + 1
		} else {
			result[value] = 1
		}
	}

	return result
}
