package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Hand struct {
	cards string
	bid int
	value int
}

type SortedHands []Hand

func (hands SortedHands) Len() int {
	return len(hands)
}

func (hands SortedHands) Less(i, j int) bool {
	return hands[i].value < hands[j].value
}

func (hands SortedHands) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func recognizeHandType(cards string) int {
	result := 0
	localMap := map[rune]int{}

	for _, card := range cards {
		value, ok := localMap[card]

		if (ok) {
			localMap[card] = value + 1
		} else {
			localMap[card] = 1
		}
	}

	mapLength := len(localMap)
	power := int(math.Pow(13, 5))

	if (mapLength == 1) {
		// Five of a kind
		result = 6 * power 
	} else if (mapLength == 2) {
		value, ok := localMap[rune(cards[0])]

		if (ok == false) {
			fmt.Println("ERROR")
			return result;
		}

		if (value == 1 || value == 4) {
			// Four of a kind
			result = 5 * power 
		} else {
			// Full house
			result = 4 * power 
		}
	} else if (mapLength == 3) {
		value, ok := localMap[rune(cards[0])]

		if (ok == false) {
			fmt.Println("ERROR")
			return result;
		}

		if (value == 1) {
			value, ok = localMap[rune(cards[1])]

			if (ok == false) {
				fmt.Println("ERROR")
				return result;
			}
		}

		if (value == 1 || value == 3) {
			// Three of a kind
			result = 3 * power
		} else {
			// Two pair
			result = 2 * power
		}
	} else if (mapLength == 4) {
		// One pair
		result = 1 * power 
	}
	
	return result
}

func calculateCardsHandValue(cards string) int {
	sum := recognizeHandType(cards)

	figureValueMap := map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}

	for index, card := range cards {
		value, ok := figureValueMap[string(card)]

		if (ok == false) {
			continue
		}

		sum += value * int(math.Pow(13, float64(4 - index)))
	}
	
	return sum
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	sum := 0
	hands := make(SortedHands, 0)

	for sc.Scan() {
		line := sc.Text()
		splitLine := strings.Split(line, " ")

		if (len(splitLine) < 2) {
			continue
		}

		bid, err := strconv.Atoi(splitLine[1])

		if (err != nil) {
			continue
		}

		hands = append(hands, Hand{
			cards: splitLine[0],
			bid: bid,
			value: calculateCardsHandValue(splitLine[0]),
		})
	}

	sort.Sort(hands)

	for index, hand := range hands {
		sum += (index + 1) * hand.bid
	}

	fmt.Println("Result: ", sum)
}
