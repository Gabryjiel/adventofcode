package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

	limitMap := map[string]int64{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	for sc.Scan() {
		line := sc.Text()
		
		isValid := true
		semicolonIndex := strings.Index(line, ":")
		gameId, err := strconv.Atoi(line[5:semicolonIndex])

		if (err != nil) {
			continue
		}

		rounds := strings.Split(line[semicolonIndex + 2:], ";")

		for _, round := range rounds {
			colours := strings.Split(round, ",")
			for _, colour := range colours {
				values := strings.Split(strings.TrimSpace(colour), " ")
				limit, ok := limitMap[values[1]]

				if (ok != true) {
					fmt.Println("No label in map")
					continue
				}

				count, err := strconv.ParseInt(values[0], 10, 64)

				if (err != nil) {
					fmt.Println("Cannot parse values[0]")
					continue
				}

				if (count > limit) {
					isValid = false
					break
				}
			}
		}
		
		if (isValid) {
			sum += gameId
		}
	}

	fmt.Println("Result: ", sum)
}
