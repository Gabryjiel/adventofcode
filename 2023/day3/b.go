package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findNumInLine(line string, pos int) int {
	cursor := pos
	current := string(line[pos])

	cursor -= 1
	for cursor > -1 {
		_, err := strconv.Atoi(string(line[cursor]))

		if (err != nil) {
			break
		}

		current = string(line[cursor]) + current
		cursor -= 1
	}

	cursor = pos + 1
	for cursor < len(line) {
		_, err := strconv.Atoi(string(line[cursor]))

		if (err != nil) {
			break
		}

		current = current + string(line[cursor])
		cursor += 1
	}
	
	result, err := strconv.Atoi(string(current))

	if (err != nil) {
		return 0
	}

	return result
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

	topLine, midLine, botLine := "", "", ""
	for sc.Scan() {
		botLine = sc.Text()
		
		for pos, char := range strings.Split(midLine, "") {
			if (char != "*") {
				continue
			}

			var nums []int
			sliceStart := max(0, pos - 1)
			sliceEnd := min(len(midLine), pos + 2)

			topSlice, botSlice := "", ""
			midLeft, midRight := "", ""

			if (topLine != "") {
				topSlice = topLine[sliceStart:sliceEnd]
			}

			if (botLine != "") {
				botSlice = botLine[sliceStart:sliceEnd]
			}

			if (sliceStart > 0) {
				midLeft = string(midLine[sliceStart])
			}

			if (sliceEnd < len(midLine)) {
				midRight = string(midLine[sliceEnd - 1])
			}

			_, err := strconv.Atoi(midLeft)
			if (err == nil) {
				num := findNumInLine(midLine, pos - 1)	
				nums = append(nums, num)
			}

			_, err = strconv.Atoi(midRight)
			if (err == nil) {
				num := findNumInLine(midLine, pos + 1)	
				nums = append(nums, num)
			}

			isDigitPresent := false
			for idx, val := range topSlice {
				_, err := strconv.Atoi(string(val))
				if (err == nil) {
					if (isDigitPresent == false) {
						num := findNumInLine(topLine, pos - 1 + idx)
						nums = append(nums, num)
						isDigitPresent = true
					}
				} else {
					isDigitPresent = false
				}
			}


			isDigitPresent = false
			for idx, val := range botSlice {
				_, err := strconv.Atoi(string(val))

				if (err == nil) {
					if (isDigitPresent == false) {
						num := findNumInLine(botLine, pos - 1 + idx)
						nums = append(nums, num)
						isDigitPresent = true
					}
				} else {
					isDigitPresent = false
				}
			}

			if (len(nums) < 2) {
				continue
			}

			sum += nums[0] * nums[1]
		}

		topLine = midLine
		midLine = botLine
	}

	fmt.Println("Result: ", sum)
}
