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

func c(counts[][]int, lineIdx, numIdx int) int {
	if lineIdx == -1 && numIdx == 0 {
		return 1
	}

	if lineIdx >= 0 && numIdx >= 0 {
		return counts[lineIdx][numIdx]
	}

	return 0
}

func solve(question string, info []int, r int) int {
	line := repeatNTimes(question, r)
	nums := repeatNTimesSlice(info, r)
	nums = append([]int{0}, nums...)

	counts := make([][]int, len(line))

	for numIdx := 0; numIdx < len(nums); numIdx++ {
		for lineIdx := 0; lineIdx < len(line); lineIdx++ {
			current := 0

			if line[lineIdx] != '#' {
				current += c(counts, lineIdx - 1, numIdx)
			}

			if numIdx > 0 {
				shouldCount := true
				
				if line[lineIdx] == '#' {
					shouldCount = false
				}

				if shouldCount {
					for k := 1; k <= nums[numIdx]; k++ {
						if lineIdx - k >= 0 && line[lineIdx - k] == '.' {
							shouldCount = false
						}
					}
				}

				if shouldCount {
					current += c(counts, lineIdx - nums[numIdx] - 1, numIdx - 1)
				}
			}

			counts[lineIdx] = append(counts[lineIdx], current)
		}

	}	

	return counts[len(line) - 1][len(nums) - 1]
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

func repeatNTimes(line string, n int) string {
	result := ""

	for i := 0; i < n; i++ {
		result += line + "?"
	}

	return result
}

func repeatNTimesSlice(slice []int, n int) []int {
	result := make([]int, len(slice) * n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(slice); j++ {
			result[i * len(slice) + j] = slice[j]
		}
	}

	return result
}

func countPositions(question string, info []int) int {
	result := 0

	result += solve(question, info, 5)

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
	fmt.Println(question,availableIndexes, requiredChoice, len(choices))
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
