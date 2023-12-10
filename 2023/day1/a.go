package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()

	if (err != nil) {
		fmt.Errorf("File not found")
		return
	}

	sum := 0
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		line := sc.Text()
		firstDigitIdx := strings.IndexFunc(line, unicode.IsNumber)
		lastDigitIdx := strings.LastIndexFunc(line, unicode.IsNumber)
		numberStr := string(line[firstDigitIdx]) + string(line[lastDigitIdx])
		number, err := strconv.Atoi(numberStr)

		if (err != nil) {
			continue
		}

		sum += number
	}

	fmt.Println("Everything: %d", sum)
}
