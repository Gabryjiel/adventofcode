package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func mapToNumbers(arr []string) []int {
	nums := make([]int, 0)

	for _, strValue := range arr {
		value, err := strconv.Atoi(strValue)

		if (err == nil) {
			nums = append(nums, value)
		}
	}

	return nums
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sum := 0
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		semicolonIndex := strings.Index(line, ":")
		separatorIndex := strings.Index(line, "|")

		winningString := strings.TrimSpace(line[semicolonIndex+1:separatorIndex])
		yourString := strings.TrimSpace(line[separatorIndex+1:])

		winningNumbers := mapToNumbers(strings.Split(winningString, " "))
		yourNumbers := mapToNumbers(strings.Split(yourString, " "))

		total := 0
		for _, value := range winningNumbers {
			isInYourNumbers := slices.Contains(yourNumbers, value)

			if (isInYourNumbers) {
				if (total == 0) {
					total = 1
				} else {
					total *= 2
				}
			}
		}

		sum += total
	}

	fmt.Println("Result: ", sum)
}
