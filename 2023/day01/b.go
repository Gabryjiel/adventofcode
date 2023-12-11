package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"strconv"
)

func InsertLast(x, y, z string) (x2 string) {
	i := strings.LastIndex(x, y)
	if i == -1 {
		return x
	}
	return x[:i] + z + x[i:]
}

func InsertFirst(x, y, z string) string {
	i := strings.Index(x, y)
	if i == -1 {
		return x
	}
	return x[:i] + z + x[i:]
}

func main() {
	input, err := os.Open("input.txt")

	defer input.Close()

	if (err != nil) {
		fmt.Errorf("File not found")
		return
	}

	sum := 0
	sc := bufio.NewScanner(input)

	symbols := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for sc.Scan() {
		line := sc.Text()

		firstSymbolIdx, lastSymbolIdx := -1, -1
		firstSymbolNum, lastSymbolNum := -1, -1
		for index, symbol := range symbols {
			first := strings.Index(line, symbol)
			last := strings.LastIndex(line, symbol)

			if (first != -1 && (first < firstSymbolIdx || firstSymbolIdx == -1)) {
				firstSymbolIdx = first
				firstSymbolNum = index
			}

			if (last != -1 && (last > lastSymbolIdx || lastSymbolIdx == -1)) {
				lastSymbolIdx = last
				lastSymbolNum = index
			}
		}

		if (firstSymbolNum != -1) {
			line = InsertFirst(line, symbols[firstSymbolNum], strconv.FormatInt(int64(firstSymbolNum + 1), 10))
		}

		if (lastSymbolNum != -1) {
			line = InsertLast(line, symbols[lastSymbolNum], strconv.FormatInt(int64(lastSymbolNum + 1), 10))
		}

		firstDigitIdx := strings.IndexFunc(line, unicode.IsNumber)
		lastDigitIdx := strings.LastIndexFunc(line, unicode.IsNumber)

		if (firstDigitIdx == -1 || lastDigitIdx == -1) {
			continue
		}

		numberStr := string(line[firstDigitIdx]) + string(line[lastDigitIdx])
		number, err := strconv.Atoi(numberStr)

		if (err != nil) {
			continue
		}

		sum += number
	}

	fmt.Println("Result:", sum)
}
