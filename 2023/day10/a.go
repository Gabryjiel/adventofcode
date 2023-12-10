package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkConnectionFromEast(chars [][]string, startX, startY int) bool {
	connectFromWest := "|7J"
	
	if (startX < len(chars[0])) {
		nextChar := chars[startY][startX + 1]

		if (strings.Contains(connectFromWest, nextChar)) {
			return true
		}
	}

	return false
}

func checkConnectionFromWest(chars [][]string, startX, startY int) bool {
	connectFromEast := "-LF"
	
	if (startX > 0) {
		nextChar := chars[startY][startX - 1]

		if (strings.Contains(connectFromEast, nextChar)) {
			return true
		}
	}

	return false
}

func checkConnectionFromNorth(chars [][]string, startX, startY int) bool {
	connectFromSouth := "|F7"
	
	if (startY > 0) {
		nextChar := chars[startY - 1][startX]

		if (strings.Contains(connectFromSouth, nextChar)) {
			return true
		}
	}

	return false
}

func checkConnectionFromSouth(chars [][]string, startX, startY int) bool {
	connectFromNorth := "|LJ"
	
	if (startY < len(chars)) {
		nextChar := chars[startY + 1][startX]

		if (strings.Contains(connectFromNorth, nextChar)) {
			return true
		}
	}

	return false
}

func checkDirectionForStart(chars [][]string, curX, curY int) (int, int) {
	if (checkConnectionFromNorth(chars, curX, curY)) {
		return 0, -1
	} else if (checkConnectionFromSouth(chars, curX, curY)) {
		return 0, 1
	} else if (checkConnectionFromEast(chars, curX, curY)) {
		return 1, 0
	} else if (checkConnectionFromWest(chars, curX, curY)) {
		return -1, 0
	}

	return 0, 0
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	chars := make([][]string, 0)
	curX, curY := 0, 0

	for sc.Scan() {
		line := sc.Text()
		splitLine := strings.Split(line, "")

		if (strings.Contains(line, "S")) {
			curX = strings.Index(line, "S")
			curY = len(chars)
		}

		chars = append(chars, splitLine)
	}

	moveX, moveY := checkDirectionForStart(chars, curX, curY)
	steps := 0

	for ;; {
		curX += moveX
		curY += moveY
		nextChar := chars[curY][curX]

		if (nextChar == ".") {
			fmt.Println("ERROR")
			break
		} else if (nextChar == "S") {
			break
		} else if (nextChar == "L") {
			if (moveX == 0) {
				moveX = 1
				moveY = 0
			} else {
				moveX = 0
				moveY = -1
			}
		} else if (nextChar == "J") {
			if (moveX == 0) {
				moveX = -1
				moveY = 0
			} else {
				moveX = 0
				moveY = -1
			}
		} else if (nextChar == "F") {
			if (moveX == 0) {
				moveX = 1
				moveY = 0
			} else {
				moveX = 0
				moveY = 1
			}
		} else if (nextChar == "7") {
			if (moveX == 0) {
				moveX = -1
				moveY = 0
			} else {
				moveX = 0
				moveY = 1
			}
		}

		steps += 1
	}

	fmt.Println("Result: ", (steps + 1) / 2)
}
