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

type Plane [][]Tile

func (plane *Plane) disenergize() {
	for y := 0; y < len(*plane); y++ {
		for x := 0; x < len((*plane)[0]); x++ {
			if (*plane)[y][x].isEnergized == true {
				(*plane)[y][x].isEnergized = false
			}
		}
	}
}

func (plane *Plane) countEnergizedTiles() int {
	sum := 0

	for _, row := range (*plane) {
		for _, tile := range row {
			if tile.isEnergized {
				sum += 1
			}
		}
	}

	return sum
}

type Beam struct {
	direction rune
	x, y int
}

func (beam Beam) getNextBeam() Beam {
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

func (beam Beam) isBeamValid(tiles Plane) bool {
	if beam.x >= 0 && beam.x < len(tiles[0]) && beam.y >= 0 && beam.y < len(tiles) {
		return true
	}

	return false 
}

func main() {
	content, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	plane := getPlane(content)
	maximum := tracker(plane)
	fmt.Println("Result:", maximum)
}

func getPlane(content []string) Plane {
	plane := make(Plane, len(content))

	for lineIndex, line := range content {
		row := make([]Tile, len(line))
		
		for charIndex, char := range line {
			row[charIndex] = Tile {
				isEnergized: false,
				symbol: char,
			}
		}

		plane[lineIndex] = row
	}

	return plane
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

func tracker(tiles Plane) int {
	maxLeft := trackOneEdge(tiles, len(tiles), '>', -1, -2)
	maxRight := trackOneEdge(tiles, len(tiles), '<', len(tiles[0]), -2)
	maxTop := trackOneEdge(tiles, len(tiles[0]), 'v', -2, -1)
	maxBottom := trackOneEdge(tiles, len(tiles[0]), '^', -2, len(tiles))
	
	return max(maxLeft, maxRight, maxTop, maxBottom)
}

func trackOneEdge(plane Plane, maxIter int, direction rune, xStart, yStart int) int {
	maximum := 0
	x, y := xStart, yStart

	for i := 0; i < maxIter; i++ {
		if xStart == -2 {
			x = i
		}

		if yStart == -2 {
			y = i
		}

		startingBeam := Beam{ direction, x, y }
		sum := trackFromStartingBeam(plane, startingBeam)

		if sum > maximum {
			maximum = sum
		}

		plane.disenergize()
	}

	return maximum
}

func trackFromStartingBeam(tiles Plane, startingBeam Beam) int {
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

	return tiles.countEnergizedTiles()
}

func checkIfBeamExisted(currentBeam Beam, historyBeams []Beam) bool {
	for _, prevBeam := range historyBeams {
		if currentBeam.x == prevBeam.x &&
			currentBeam.y == prevBeam.y &&
			(currentBeam.direction == prevBeam.direction) {
			return true
		}
	}

	return false 
}

func trackSingleBeam(beam Beam, plane Plane) ([]Beam, []Beam) {
	newBeam := Beam{ 
		direction: beam.direction, 
		x: beam.x,
		y: beam.y,
	}

	var newBeams []Beam = nil
	historyBeams := make([]Beam, 0)

	for ;; {
		newBeam = newBeam.getNextBeam() 

		if newBeam.isBeamValid(plane) == false {
			break
		}

		historyBeams = append(historyBeams, newBeam)

		currentTile := plane[newBeam.y][newBeam.x]
		plane[newBeam.y][newBeam.x].isEnergized = true
		
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
			newBeams = make([]Beam, 2)
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

			break
		} else if currentTile.symbol == '-' {
			newBeams = make([]Beam, 2)
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

			break	
		}
	}

	return newBeams, historyBeams
}

func isBeamNotChanging(beam Beam, tile Tile) bool {
	if tile.symbol == '.' ||
		(tile.symbol == '-' && (beam.direction == '<' || beam.direction == '>')) ||
		(tile.symbol == '|' && (beam.direction == '^' || beam.direction == 'v')) {
		return true
	}

	return false
}
