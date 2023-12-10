package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time int
	distance int
}

func trimAllSpaces(line string) string {
	newLine := ""

	for _, char := range line {
		if (string(char) == " ") {
			continue
		}

		newLine += string(char)
	}

	return newLine
}

func getRangeForRace(race Race) (int, int) {
	minTime := 0
	maxTime := race.time

	for i := 0; i <= race.time; i++ {
		chargeTime := i
		runTime := race.time - i

		coveredDistance := chargeTime * runTime

		if (coveredDistance > race.distance) {
			minTime = chargeTime
			break
		}
	}

	for i := 0; i <= race.time; i++ {
		chargeTime := race.time - i
		runTime := i

		coveredDistance := chargeTime * runTime

		if (coveredDistance > race.distance) {
			maxTime = chargeTime
			break
		}
	}

	return minTime, maxTime
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	race := Race{
		time: 0,
		distance: 0,
	}

	for sc.Scan() {
		line := sc.Text()
		line = trimAllSpaces(line)
		semicolonIndex := strings.Index(line, ":") + 1

		isTimeLine := strings.Contains(line, "Time:")
		if (isTimeLine) {
			line = line[semicolonIndex:]
			timeNum, err := strconv.Atoi(line)

			if (err != nil) {
				fmt.Println("Cannot convert to int", line)
			}

			race = Race{
				time: timeNum,
				distance: 0,
			}
		}

		isDistanceLine := strings.Contains(line, "Distance:")
		if (isDistanceLine) {
			line = line[semicolonIndex:]
			distanceNum, err := strconv.Atoi(line)
			
			if (err != nil) {
				fmt.Println("Cannot convert to int", line)
			}

			race.distance = distanceNum
		}
	}

	minTime, maxTime := getRangeForRace(race)
	sum := maxTime - minTime + 1

	fmt.Println("Result: ", sum)
}
