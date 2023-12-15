package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	focalLength int
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)

	boxes := make([][]Lens, 256)

	for sc.Scan() {
		line := sc.Text()
		commands := strings.Split(line, ",")

		for _, command := range commands {
			processCommand(command, boxes)
		}
	}

	sum := 0
	for boxIndex, box := range boxes {
		for lensIndex, lens := range box {
			sum += (boxIndex + 1) * (lensIndex + 1) * lens.focalLength
		}
	}

	fmt.Println("Result:", sum)
}

func hashing(command string) int {
	localValue := 0

	for _, char := range command {
		localValue += int(char)
		localValue *= 17
		localValue = localValue % 256
	}

	return localValue
}

func getLensIndexFromBoxHash(boxes [][]Lens, hash int, label string) int {
	return slices.IndexFunc(boxes[hash], func(box Lens) bool {
		return box.label == label
	})
}

func processCommand(command string, boxes [][]Lens) {
	doesContainEqual := strings.Contains(command, "=")
	if doesContainEqual {
		commandSplit := strings.Split(command, "=")

		focalLength, err := strconv.Atoi(commandSplit[1])

		if err != nil {
			fmt.Println(commandSplit[1], err)
			return
		}

		newLens := Lens{
			label: commandSplit[0],
			focalLength: focalLength,
		}

		hash := hashing(newLens.label)	
		lensIndex := getLensIndexFromBoxHash(boxes, hash, newLens.label)

		if lensIndex == -1 {
			boxes[hash] = append(boxes[hash], newLens)
		} else {
			boxes[hash][lensIndex] = newLens
		}
	}

	dashIndex := strings.Index(command, "-")
	if dashIndex != -1 {
		label := command[0:dashIndex]
		hash := hashing(label)
		lensIndex := getLensIndexFromBoxHash(boxes, hash, label)

		if lensIndex == -1 {
			return
		}

		boxes[hash] = removeItemFromBox(boxes[hash], lensIndex)
	}
}

func removeItemFromBox(box []Lens, indexToRemove int) []Lens {
	return append(box[:indexToRemove], box[indexToRemove + 1:]...)
}
