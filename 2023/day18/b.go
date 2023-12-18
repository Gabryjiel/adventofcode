// taken from "https://github.com/dannyvankooten/advent-of-code/blob/main/2023/18-lavaduct-lagoon/main.go"
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	x, y int
}

type Command struct {
	direction rune
	num int
	color string
}

type Plan []Command

func main() {
	timeStart := time.Now()
	content, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	plan := makePlan(content)
	result := solve(plan)

	fmt.Println("Result:", result)
	fmt.Println("Time:", float64(time.Since(timeStart).Microseconds()) / 1000, "ms")
}

func openFile(name string) ([]string, error) {
	file, err := os.Open(name)	
	defer file.Close()

	content := make([]string, 0)

	if (err != nil) {
		return nil, fmt.Errorf("File %s not found", name)
	}
	
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		content = append(content, sc.Text())
	}

	return content, nil
}

func makePlan(content []string) Plan {
	result := make(Plan, len(content))

	directionMap := map[rune]rune {
		'0': 'R',
		'1': 'D',
		'2': 'L',
		'3': 'U',
	}

	for index, row := range content {
		rowElements := strings.Split(row, " ")

		color := rowElements[2]
		directionRune := rune(color[7])

		direction, ok := directionMap[directionRune] 

		if ok == false {
			fmt.Println("[Error] Could not extract value from map", directionRune)
			break
		}

		numStr := color[2:7]
		num64, err := strconv.ParseInt(numStr, 16, 64)

		if err != nil {
			fmt.Println("[Error] Cound not convert to int", numStr)
		}

		num := int(num64)

		result[index] = Command{
			direction,
			num,
			color,
		}
	}

	return result
}

func solve(plan Plan) int {
	vectors := make([]Position, 0)
	x, y := 0, 0
	vectors = append(vectors, Position{ x, y })
	lineLength := 0

	for _, item := range plan {
		dy, dx := 0, 0

		switch item.direction {
		case 'R':
			dx = 1
			break
		case 'L':
			dx = -1
			break
		case 'U':
			dy = -1
			break
		case 'D':
			dy = 1
			break
		}

		x += dx * item.num
		y += dy * item.num
		lineLength += item.num
		vectors = append(vectors, Position{ x, y })
	}

	return (lineLength / 2 + area(vectors) + 1)
}

func area(positions []Position) int {
	result := 0

	i := 0
	j := len(positions) - 1
	for i < len(positions)  {
		result += positions[i].x * positions[j].y - positions[j].x * positions[i].y
		j = i
		i += 1
	}

	if result < 0 {
		result = -result
	}

	return result / 2
}

