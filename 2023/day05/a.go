package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MapEntry struct {
	destinationStart, sourceStart, length int
}

func getSeedsFromSeedLine(line string) []int {
	seeds := make([]int, 0)
	seedsStr := strings.Split(line, " ")[1:]

	for _, seedStr := range seedsStr {
		seed, err := strconv.Atoi(seedStr)

		if (err != nil) {
			fmt.Println("Could not convert seedStr to seed", seedStr)
		}
		seeds = append(seeds, seed) 
	}

	return seeds
}

func getMapEntryFromLine(line string) MapEntry {
	numsStr := strings.Split(line, " ")
	mapEntry := MapEntry {
		destinationStart: 0,
		sourceStart: 0,
		length: 0,
	}
	
	for index, numStr := range numsStr {
		num, err := strconv.Atoi(numStr)

		if (err != nil) {
			continue
		}

		if (index == 0) {
			mapEntry.destinationStart = num
		} else if (index == 1) {
			mapEntry.sourceStart = num
		} else {
			mapEntry.length = num
		}
	}

	return mapEntry
}

func getSeedThrough(seed int, mapEntries []MapEntry) int {
	for _, mapEntry := range mapEntries {
		sourceEnd := mapEntry.sourceStart + mapEntry.length

		if (seed >= mapEntry.sourceStart && seed <= sourceEnd) {
			return seed + (mapEntry.destinationStart - mapEntry.sourceStart)
		}
	}

	return seed
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	seeds := make([]int, 0)
	maps := make([][]MapEntry, 0)

	for sc.Scan() {
		line := sc.Text()

		isSeedsLine := strings.Contains(line, "seeds:")
		if (isSeedsLine) {
			seeds = getSeedsFromSeedLine(line)
			continue
		}

		isMapLine := strings.Contains(line, "map:")
		if (isMapLine) {
			maps = append(maps, []MapEntry{})
			continue
		}

		isEmptyLine := len(line) == 0
		if (isEmptyLine) {
			continue
		}

		lastMapIndex := len(maps) - 1
		mapEntry := getMapEntryFromLine(line)
		maps[lastMapIndex] = append(maps[lastMapIndex], mapEntry)
	}

	locations := make([]int, 0)
	for _, seed := range seeds {
		temp := seed
		for _, mapEntries := range maps {
			temp = getSeedThrough(temp, mapEntries)
		}
		locations = append(locations, temp)
	}

	minLocation := locations[0]
	for _, location := range locations {
		if (location < minLocation) {
			minLocation = location
		}
	}

	fmt.Println("Result: ", minLocation)

}
