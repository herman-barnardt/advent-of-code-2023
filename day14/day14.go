package day14

import (
	"math"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 14, solve2023Day14Part1, solve2023Day14Part2)
}

var directions = map[string]util.Point{
	"NORTH": {X: 0, Y: -1},
	"WEST":  {X: -1, Y: 0},
	"SOUTH": {X: 0, Y: 1},
	"EAST":  {X: 1, Y: 0},
}

func moveRocks(world map[util.Point]string, direction, min, max util.Point, flipped bool) map[util.Point]string {
	if !flipped {
		for y := min.Y; y < max.Y; y++ {
			for x := min.X; x < max.X; x++ {
				if c := world[util.Point{X: x, Y: y}]; c == "O" {
					p := util.Point{X: x, Y: y}
					newPoint := p.Add(direction)
					for newPoint.Y >= min.Y && newPoint.Y < max.Y && newPoint.X >= min.X && newPoint.X < max.X && world[newPoint] == "." {
						world[p] = "."
						p = newPoint
						world[p] = "O"
						newPoint = p.Add(direction)
					}
				}
			}
		}
	} else {
		for y := max.Y - 1; y >= min.Y; y-- {
			for x := max.X - 1; x >= min.X; x-- {
				if c := world[util.Point{X: x, Y: y}]; c == "O" {
					p := util.Point{X: x, Y: y}
					newPoint := p.Add(direction)
					for newPoint.Y >= min.Y && newPoint.Y < max.Y && newPoint.X >= min.X && newPoint.X < max.X && world[newPoint] == "." {
						world[p] = "."
						p = newPoint
						world[p] = "O"
						newPoint = p.Add(direction)
					}
				}
			}
		}
	}

	return world
}

func solve2023Day14Part1(lines []string) interface{} {
	world := make(map[util.Point]string)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			world[util.Point{X: x, Y: y}] = c
		}
	}
	for y := range lines {
		for x := range lines[0] {
			c := world[util.Point{X: x, Y: y}]
			if c == "O" {
				world = moveRocks(world, directions["NORTH"], util.Point{X: 0, Y: 0}, util.Point{X: len(lines[0]), Y: len(lines)}, false)
			}
		}
	}

	sum := 0
	for point, rock := range world {
		if rock == "O" {
			sum += len(lines) - point.Y
		}
	}
	return sum
}

type state struct {
	visitedCount int
	cycleNumber  int
	load         int
}

func solve2023Day14Part2(lines []string) interface{} {
	world := make(map[util.Point]string)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			world[util.Point{X: x, Y: y}] = c
		}
	}
	visitedStates := make(map[string]state)

	cycleCount := 0
	loopCount := 0
	for true {
		cycleCount++
		world = moveRocks(world, directions["NORTH"], util.Point{X: 0, Y: 0}, util.Point{X: len(lines[0]), Y: len(lines)}, false)
		world = moveRocks(world, directions["WEST"], util.Point{X: 0, Y: 0}, util.Point{X: len(lines[0]), Y: len(lines)}, false)
		world = moveRocks(world, directions["SOUTH"], util.Point{X: 0, Y: 0}, util.Point{X: len(lines[0]), Y: len(lines)}, true)
		world = moveRocks(world, directions["EAST"], util.Point{X: 0, Y: 0}, util.Point{X: len(lines[0]), Y: len(lines)}, true)

		keyString := mapToString(world, len(lines[0]), len(lines))

		load := 0
		for point, rock := range world {
			if rock == "O" {
				load += len(lines) - point.Y
			}
		}

		if _, ok := visitedStates[keyString]; ok {
			if visitedStates[keyString].visitedCount == 2 {
				break
			}
			visitedStates[keyString] = state{2, cycleCount, load}
			loopCount++
		} else {
			visitedStates[keyString] = state{1, cycleCount, load}
		}
	}

	cycleMaps := make(map[int]int)
	loopStart := math.MaxInt
	for _, value := range visitedStates {
		if value.visitedCount == 2 {
			cycleMaps[value.cycleNumber] = value.load
			if value.cycleNumber < loopStart {
				loopStart = value.cycleNumber
			}
		}
	}

	offset := len(visitedStates) - len(cycleMaps)

	index := (1000000000 - offset) % len(cycleMaps)

	return cycleMaps[loopStart+index-1]
}

func mapToString(world map[util.Point]string, maxX, maxY int) string {
	retVal := ""
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			retVal += world[util.Point{X: x, Y: y}]
		}
		retVal += "\n"
	}
	return retVal
}
