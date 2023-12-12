package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

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

		splitLine := strings.SplitN(line, " ", 2)
		question := splitLine[0]
		infoStr := splitLine[1]
		info := mapInfoStrToInfo(infoStr)		

		localCount := countPositions(question, info)
		sum += localCount
	}

	fmt. Println("Result:", sum)
}

func splitQuestion(question string) []string {
	result := make([]string, 0)

	tmp := ""
	for _, char := range question {
		if char == '.' {
			if tmp != "" {
				result = append(result, tmp)
				tmp = ""
			}
		} else {
			tmp += string(char)
		}
	}
	
	if tmp != "" {
		result = append(result, tmp)
	}

	return result
}

func countPositions(question string, info []int) int {
	result := 0

	result += bruteForceSolution(question, info)

	return result
}

func getIndexesOf(str string, char rune) []int {
	result := make([]int, 0)

	for index, value := range str {
		if value == char {
			result = append(result, index)
		}
	}

	return result
}

func sum(list []int) int {
	result := 0

	for _, value := range list {
		result += value
	}

	return result
}

func bruteForceSolution(question string, info []int) int {
	result := 0
	availableIndexes := getIndexesOf(question, '?')
	requiredChoice := sum(info) - strings.Count(question, "#")
	choices := combin.Combinations(len(availableIndexes), requiredChoice)

	for _, choice := range choices {
		solution := strings.ReplaceAll(question, "?", ".")
		for _, singleChoice := range choice {
			index := availableIndexes[singleChoice]
			solution = solution[:index] + "#" + solution[index + 1:]
		}


		isValidSolution := checkSolution(solution, info)
		if isValidSolution {
			result += 1
		}
	}

	return result
}

func checkSolution(solution string, info []int) bool {
	solutionInfo := make([]int, 0)
	
	count := 0
	for _, char := range solution {
		if char == '.' || char == '?' {
			if count != 0 {
				solutionInfo = append(solutionInfo, count)
				count = 0
			}
		} else {
			count += 1
		}
	}

	if count != 0 {
		solutionInfo = append(solutionInfo, count)
	}

	for index, solutionInfoValue := range solutionInfo {
		if solutionInfoValue != info[index] {
			return false
		}
	}

	return true
}

func mapInfoStrToInfo(infoStr string) []int {
	splitInfoStr := strings.Split(infoStr, ",")
	info := make([]int, 0)

	for _, value := range splitInfoStr {
		num, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("Could not convert", num)
			continue
		}

		info = append(info, num)
	}

	return info
}
