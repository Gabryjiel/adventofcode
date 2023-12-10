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

func trimInnerSpaces(line string) string {
	newLine := ""

	for index, char := range line {
		if (string(char) == " " && len(line) > index + 1 && string(line[index + 1]) == " ") {
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
	races := make([]Race, 0)

	for sc.Scan() {
		line := sc.Text()
		line = trimInnerSpaces(line)
		semicolonIndex := strings.Index(line, ":") + 2

		isTimeLine := strings.Contains(line, "Time:")
		if (isTimeLine) {
			line = line[semicolonIndex:]
			times := strings.Split(line, " ")
	
			for _, time := range times {
				timeNum, err := strconv.Atoi(time)

				if (err != nil) {
					fmt.Println("Cannot convert to int", time, times)
				}

				newRace := Race{
					time: timeNum,
					distance: 0,
				}
				races = append(races, newRace)
			}
		}

		isDistanceLine := strings.Contains(line, "Distance:")
		if (isDistanceLine) {
			line = line[semicolonIndex:]
			distances := strings.Split(line, " ")

			for idx, distance := range distances {
				distanceNum, err := strconv.Atoi(distance)
				
				if (err != nil) {
					fmt.Println("Cannot convert to int", distance)
				}

				races[idx].distance = distanceNum
			}
		}
	}

	sum := 1
	for _, race := range races {
		minTime, maxTime := getRangeForRace(race)
		sum *= maxTime - minTime + 1
	}

	fmt.Println("Result: ", sum)

}
