package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile struct {
	symbol rune
	isEnergized bool
}

type Beam struct {
	direction rune
	x, y int
}

func main() {
	file, err := os.Open(os.Args[1])	
	defer file.Close()

	if (err != nil) {
		fmt.Printf("File %s not found", os.Args[1])
		return
	}
	
	sc := bufio.NewScanner(file)
	tiles := make([][]Tile, 0)

	for sc.Scan() {
		line := sc.Text()

		row := make([]Tile, len(line))
		for index, value := range line {
			row[index] = Tile {
				isEnergized: false,
				symbol: value,
			}
		}
		tiles = append(tiles, row)
	}

	maximum := tracker(tiles)
	fmt.Println("Result:", maximum)
}

func tracker(tiles [][]Tile) int {
	maxLeft := trackOneEdge(tiles, len(tiles), '>', -1, -2)
	maxRight := trackOneEdge(tiles, len(tiles), '<', len(tiles[0]), -2)
	maxTop := trackOneEdge(tiles, len(tiles[0]), 'v', -2, -1)
	maxBottom := trackOneEdge(tiles, len(tiles[0]), '^', -2, len(tiles))
	
	return max(maxLeft, maxRight, maxTop, maxBottom)
}

func trackOneEdge(tiles [][]Tile, maxIter int, direction rune, xStart, yStart int) int {
	maximum := 0
	x, y := xStart, yStart

	for i := 0; i < maxIter; i++ {
		if xStart == -2 {
			x = i
		}

		if yStart == -2 {
			y = i
		}

		startingBeam := Beam{ 
			direction,
			x, 
			y,
		}
		sum := trackFromStartingBeam(tiles, startingBeam)

		if sum > maximum {
			maximum = sum
		}

		resetTiles(tiles)
	}

	return maximum
}

func sumEnergizedTiles(tiles [][]Tile) int {
	sum := 0

	for _, row := range tiles {
		for _, tile := range row {
			if tile.isEnergized {
				sum += 1
			}
		}
	}

	return sum
}

func resetTiles (tiles [][]Tile) {
	for y := 0; y < len(tiles); y++ {
		for x := 0; x < len(tiles[0]); x++ {
			if tiles[y][x].isEnergized == true {
				tiles[y][x].isEnergized = false
			}
		}
	}
}

func trackFromStartingBeam(tiles [][]Tile, startingBeam Beam) int {
	beams := make([]Beam, 1)
	beams[0] = startingBeam

	historyBeams := make([]Beam, 0)

	for ;; {
		if len(beams) == 0 {
			break
		}

		currentBeam := beams[0]
		beams = beams[1:]

		if checkIfBeamExisted(currentBeam, historyBeams) == true {
			continue
		}

		newBeams, newHistoryBeams := trackSingleBeam(currentBeam, tiles)
		if newBeams != nil {
			beams = append(beams, newBeams...)
		}

		historyBeams = append(historyBeams, currentBeam)
		historyBeams = append(historyBeams, newHistoryBeams...)
	}

	return sumEnergizedTiles(tiles)
}

func checkIfBeamExisted(currentBeam Beam, historyBeams []Beam) bool {
	for _, prevBeam := range historyBeams {
		if currentBeam.x == prevBeam.x &&
			currentBeam.y == prevBeam.y &&
			currentBeam.direction == prevBeam.direction {
			return true
		}
	}

	return false 
}

func trackSingleBeam(beam Beam, tiles [][]Tile) ([]Beam, []Beam) {
	newBeam := Beam{ 
		direction: beam.direction, 
		x: beam.x,
		y: beam.y,
	}

	historyBeams := make([]Beam, 0)

	for ;; {
		newBeam = getBeamFromNextStep(newBeam)

		if checkIfBeamIsValid(newBeam, tiles) == false {
			return nil, historyBeams
		}

		historyBeams = append(historyBeams, newBeam)

		currentTile := tiles[newBeam.y][newBeam.x]
		tiles[newBeam.y][newBeam.x].isEnergized = true
		
		if isBeamNotChanging(newBeam, currentTile) == true {
			continue
		} else if currentTile.symbol == '/' {
			if newBeam.direction == '>' {
				newBeam.direction = '^'
			} else if newBeam.direction == '<' {
				newBeam.direction = 'v'
			} else if newBeam.direction == '^' {
				newBeam.direction = '>'
			} else if newBeam.direction == 'v' {
				newBeam.direction = '<'
			}
		} else if currentTile.symbol == '\\' {
			if newBeam.direction == '>' {
				newBeam.direction = 'v'
			} else if newBeam.direction == '<' {
				newBeam.direction = '^'
			} else if newBeam.direction == '^' {
				newBeam.direction = '<'
			} else if newBeam.direction == 'v' {
				newBeam.direction = '>'
			}	
		} else if currentTile.symbol == '|' {
			newBeams := make([]Beam, 2)
			newBeams[0] = Beam{
				direction: '^',
				x: newBeam.x,
				y: newBeam.y,
			}
			newBeams[1] = Beam{
				direction: 'v',
				x: newBeam.x,
				y: newBeam.y,
			}

			return newBeams, historyBeams
		} else if currentTile.symbol == '-' {
			newBeams := make([]Beam, 2)
			newBeams[0] = Beam{
				direction: '<',
				x: newBeam.x,
				y: newBeam.y,
			}
			newBeams[1] = Beam{
				direction: '>',
				x: newBeam.x,
				y: newBeam.y,
			}

			return newBeams, historyBeams
		}
	}
}

func isBeamNotChanging(beam Beam, tile Tile) bool {
	if tile.symbol == '.' ||
		(tile.symbol == '-' && (beam.direction == '<' || beam.direction == '>')) ||
		(tile.symbol == '|' && (beam.direction == '^' || beam.direction == 'v')) {
		return true
	}

	return false
}

func checkIfBeamIsValid(beam Beam, tiles [][]Tile) bool {
	if beam.x >= 0 && beam.x < len(tiles[0]) && beam.y >= 0 && beam.y < len(tiles) {
		return true
	}

	return false 
}

func getBeamFromNextStep(beam Beam) Beam {
	addX, addY := 0, 0

	if beam.direction == '<' {
		addX = -1
	} else if beam.direction == '>' {
		addX = 1
	} else if beam.direction == '^' {
		addY = -1
	} else if beam.direction == 'v' {
		addY = 1
	}

	newBeam := Beam{
		direction: beam.direction,
		y: beam.y + addY,
		x: beam.x + addX,
	}

	return newBeam
}
