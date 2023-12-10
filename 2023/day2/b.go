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
	total := 0

	for sc.Scan() {
		line := sc.Text()
		
		limitMap := map[string]int64{
			"red": 0,
			"green": 0,
			"blue": 0,
		}
		semicolonIndex := strings.Index(line, ":")

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
					limitMap[values[1]] = count
				}
			}
		}
		
		
		total += int(limitMap["red"]) * int(limitMap["green"]) * int(limitMap["blue"])
	}

	fmt.Println("Result: ", total)
}
