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

type Seed struct {
	beginning, end int
}

func getSeedsFromSeedLine2(line string) []Seed {
	seeds := make([]Seed, 0)
	seedsStr := strings.Split(line, " ")[1:]

	seedBegin := -1
	for _, seedStr := range seedsStr {
		seed, err := strconv.Atoi(seedStr)

		if (err != nil) {
			fmt.Println("Could not convert seedStr to seed", seedStr)
		}

		if (seedBegin == -1) {
			seedBegin = seed
		} else {
			seeds = append(seeds, Seed{
				beginning: seedBegin,
				end: seedBegin + seed - 1,
			})
			seedBegin = -1
		}
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

func appendMany(original []Seed, newEntries []Seed) []Seed {
	appended := original

	for _, value := range newEntries {
		appended = append(appended, value)
	}

	return appended
}

func getSeedThroughSingleMap(seed Seed, singleMap []MapEntry) []Seed {
	newSeeds := make([]Seed, 1)
	newSeeds[0] = seed

	for _, mapEntry := range singleMap {
		sourceEnd := mapEntry.sourceStart + mapEntry.length
		diff := mapEntry.destinationStart - mapEntry.sourceStart

		if ((mapEntry.sourceStart < seed.beginning && sourceEnd < seed.beginning) || (mapEntry.sourceStart > seed.end && sourceEnd < seed.end)) {
			continue
		}
			

		if (seed.beginning >= mapEntry.sourceStart && seed.end <= sourceEnd) {
			// All seed inside mapEntry (0 new seeds)
			newSeeds[0].beginning += diff
			newSeeds[0].end += diff
			break
		} else if (seed.beginning < mapEntry.sourceStart && seed.end > sourceEnd) {
			// All mapEntry inside seed (2 new seeds)
			leftSeed := Seed{
				beginning: seed.beginning,
				end: mapEntry.sourceStart - 1,
			}

			middleSeed := Seed{
				beginning: mapEntry.sourceStart,
				end: sourceEnd,
			}

			rightSeed := Seed{
				beginning: sourceEnd + 1,
				end: seed.end,
			}

			newSeeds[0] = middleSeed
			newSeeds[0].beginning += diff
			newSeeds[0].end += diff

			leftSeeds := getSeedThroughSingleMap(leftSeed, singleMap)
			rightSeeds := getSeedThroughSingleMap(rightSeed, singleMap)

			newSeeds = appendMany(newSeeds, leftSeeds)
			newSeeds = appendMany(newSeeds, rightSeeds)
			break
		} else if (seed.beginning >= mapEntry.sourceStart && seed.end > sourceEnd && seed.beginning < sourceEnd) {
			// mapEntry on a left of seed
			leftSeed := Seed{
				beginning: seed.beginning,
				end: sourceEnd,
			}

			rightSeed := Seed{
				beginning: sourceEnd + 1,
				end: seed.end,
			}

			newSeeds[0] = leftSeed
			newSeeds[0].beginning += diff
			newSeeds[0].end += diff

			rightSeeds := getSeedThroughSingleMap(rightSeed, singleMap)
			newSeeds = appendMany(newSeeds, rightSeeds)
			break
		} else if (seed.beginning < mapEntry.sourceStart && seed.end <= sourceEnd && seed.end > mapEntry.sourceStart) {
			// map Entry on a right of seed
			leftSeed := Seed{
				beginning: seed.beginning,
				end: mapEntry.sourceStart - 1,
			}

			rightSeed := Seed{
				beginning: mapEntry.sourceStart,
				end: seed.end,
			}

			newSeeds[0] = rightSeed
			newSeeds[0].beginning += diff
			newSeeds[0].end += diff

			leftSeeds := getSeedThroughSingleMap(leftSeed, singleMap)
			newSeeds = appendMany(newSeeds, leftSeeds)
			break
		}
	}

	return newSeeds
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	seeds := make([]Seed, 0)
	maps := make([][]MapEntry, 0)

	for sc.Scan() {
		line := sc.Text()

		isSeedsLine := strings.Contains(line, "seeds:")
		if (isSeedsLine) {
			seeds = getSeedsFromSeedLine2(line)
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

	for _, singleMap := range maps {
	 	newSeeds := make([]Seed, 0)

	 	for _, singleSeed := range seeds {
	 		seedsFromSingleSeed := getSeedThroughSingleMap(singleSeed, singleMap)
	 		newSeeds = appendMany(newSeeds, seedsFromSingleSeed)
	 	}

		seeds = newSeeds
	}

	minSeed := seeds[0].beginning
	for _, seed := range seeds {
		if (seed.beginning < minSeed) {
			minSeed = seed.beginning
		}
	}

	fmt.Println("Result: ", seeds, minSeed)
}
