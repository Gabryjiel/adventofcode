package day04

import "testing"

func TestSolveExampleA(t *testing.T) {
	input := parseInput("example.txt")
	result := solveA(input)

	if result != 18 {
		t.Errorf("Incorrect result! Should be 18, is %d.", result)
	}
}

func TestSolveInputA(t *testing.T) {
	input := parseInput("input.txt")
	result := solveA(input)

	if result != 2569 {
		t.Error("Incorrect")
	}
}

func TestSolveExampleB(t *testing.T) {
	input := parseInput("example.txt")
	result := solveB(input)

	if result != 9 {
		t.Errorf("Incorrect result! Should be 9, is %d.", result)
	}
}
