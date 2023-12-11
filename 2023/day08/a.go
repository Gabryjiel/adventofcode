package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Node struct {
	name string
	left string
	right string
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	nodes := make([]Node, 0)
	direction := ""

	for sc.Scan() {
		line := sc.Text()

		isNodeLine := strings.Contains(line, "=")

		if (line == "") {
			continue
		} else if (isNodeLine == false) {
			direction = line	
		} else {
			split := strings.Split(line, " ")
			name := split[0]
			left := split[2][1:4]
			right := split[3][0:3]

			node := Node{
				name: name,
				left: left,
				right: right,
			}

			nodes = append(nodes, node)
		}
	}

	currentNodeName := "AAA"
	stepCount := 0
	directionLength := len(direction)
	for ;; {
		if (currentNodeName == "ZZZ") {
			break
		}

		currentIndex := slices.IndexFunc(nodes, func(node Node) bool {
			return node.name == currentNodeName
		})

		currentNode := nodes[currentIndex]
		currentDirection := direction[stepCount % directionLength]

		if (currentDirection == 'L') {
			currentNodeName = currentNode.left
		} else if (currentDirection == 'R') {
			currentNodeName = currentNode.right
		}

		stepCount += 1
	}

	fmt.Println("Result: ", stepCount)
}
