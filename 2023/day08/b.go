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

func filterNodesBasedOnLastLetter(nodes []Node, letter byte) []Node {
	newNodes := make([]Node, 0)

	for _, node := range nodes {
		if (node.name[2] == letter) {
			newNodes = append(newNodes, node)
		}
	}

	return newNodes
}

func checkIfTravelIsFinished(nodes []Node) bool {
	for _, node := range nodes {
		if (node.name[2] != 'Z') {
			return false
		}
	}

	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(nums []int) int {
	if (len(nums) < 2) {
		return -1
	}

	a := nums[0]
	b := nums[1]
	rest := nums[2:]

	result := a * b / GCD(a, b)

	for i := 0; i < len(rest); i++ {
		temp := make([]int, 2)
		temp[0] = result
		temp[1] = rest[i]

		result = LCM(temp)
	}

	return result
}

func travel(nodes []Node, direction string) int {
	currentNodes := filterNodesBasedOnLastLetter(nodes, 'A')
	stepCounts := make([]int, 0)
	directionLength := len(direction)

	for _, currentNode := range currentNodes {
		stepCount := 0

		for ;; {
			if (currentNode.name[2] == 'Z') {
				break
			}

			currentDirection := direction[stepCount % directionLength]

			if (currentDirection == 'L') {
				newIndex := slices.IndexFunc(nodes, func(node Node) bool {
					return node.name == currentNode.left
				})
				currentNode = nodes[newIndex]
			} else if (currentDirection == 'R') {
				newIndex := slices.IndexFunc(nodes, func(node Node) bool {
					return node.name == currentNode.right
				})
				currentNode = nodes[newIndex]
			}

			stepCount += 1
		}

		stepCounts = append(stepCounts, stepCount)
	}

	result := LCM(stepCounts)

	return result
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

	stepCount := travel(nodes, direction)

	fmt.Println("Result: ", stepCount)
}
