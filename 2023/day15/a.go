package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	for sc.Scan() {
		line := sc.Text()

		commands := strings.Split(line, ",")

		for _, command := range commands {
			result := hashing(command)
			sum += result
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
