package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}

	table := make([]int, 0)
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

		if (len(table) <= 0) {
			table = append(table, 1)
		}
		winCount := 0
		currentCopies := table[0]
		table = table[1:]

		for _, value := range winningNumbers {
			isInYourNumbers := slices.Contains(yourNumbers, value)

			if (isInYourNumbers) {
				winCount++
			}
		}

		for i := 0; i < winCount; i++ {
			if (len(table) > i) {
				table[i] += currentCopies
			} else {
				table = append(table, 1 + currentCopies)
			}
		}

		sum += currentCopies
	}

	fmt.Println("Result: ", sum, table)
}
