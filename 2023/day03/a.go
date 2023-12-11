package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	current := ""
	for sc.Scan() {
		botLine = sc.Text()
		
		for pos, char := range strings.Split(midLine, "") {
			_, err := strconv.Atoi(char)

			if (err != nil) {
				current = ""
				continue
			}

			current += char
			end := false

			if (pos < len(midLine) - 1) {
				_, err := strconv.Atoi(string(midLine[pos + 1]))

				if (err != nil) {
					end = true
				}
			} else if (pos == len(midLine) - 1) {
				end = true
			}

			if (end == true) {
				topLineSlice, botLineSlice, midLineLeft, midLineRight := "", "", ".", "."
				sliceStart := max(0, pos - len(current))
				sliceEnd := min(len(midLine), pos + 2)

				if (topLine != "") {
					topLineSlice = topLine[sliceStart:sliceEnd]
				}

				if (botLine != "") {
					botLineSlice = botLine[sliceStart:sliceEnd]
				}

				if (sliceStart > 0) {
					midLineLeft = string(midLine[sliceStart])
				}

				if (sliceEnd < len(midLine)) {
					midLineRight = string(midLine[sliceEnd - 1])
				}

				topCount := strings.Count(topLineSlice, ".")
				botCount := strings.Count(botLineSlice, ".") 

				if (topCount < len(topLineSlice) || botCount < len(botLineSlice) || midLineLeft != "." || midLineRight != ".") {
					value, err := strconv.Atoi(current)

					if (err != nil) {
						fmt.Println(err, current)
						continue
					}

					sum += value
				}
			
				current = ""
			}
		}

		topLine = midLine
		midLine = botLine
	}

	fmt.Println("Result: ", sum)
}
