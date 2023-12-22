package day21

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 21, solve2023Day21Part1, solve2023Day21Part2)
}

var directions = map[string]util.Point{
	"NORTH": {X: 0, Y: -1},
	"SOUTH": {X: 0, Y: 1},
	"EAST":  {X: 1, Y: 0},
	"WEST":  {X: -1, Y: 0},
}

var grid = make(map[util.Point]string)
var pointNeighbours = make(map[util.Point][]util.Point)
var gridSize = 0

func getNeighbours(point util.Point) []util.Point {
	if neighbours, ok := pointNeighbours[point]; ok {
		return neighbours
	}
	retVal := make([]util.Point, 0)
	for _, d := range directions {
		gridPoint := point.Add(d)
		if gridPoint.X < 0 {
			gridPoint.X = gridSize + (gridPoint.X % gridSize)
		}
		if gridPoint.X >= gridSize {
			gridPoint.X = gridPoint.X % gridSize
		}
		if gridPoint.Y < 0 {
			gridPoint.Y = gridSize + (gridPoint.Y % gridSize)
		}
		if gridPoint.Y >= gridSize {
			gridPoint.Y = gridPoint.Y % gridSize
		}
		if c, ok := grid[gridPoint]; ok && c != "#" {
			retVal = append(retVal, point.Add(d))
		}
	}
	pointNeighbours[point] = retVal
	return retVal
}

func solve2023Day21Part1(lines []string) interface{} {
	start := util.Point{X: 0, Y: 0}
	for y, row := range lines {
		for x, c := range strings.Split(row, "") {
			if c == "S" {
				c = "."
				start = util.Point{X: x, Y: y}
			}
			grid[util.Point{X: x, Y: y}] = c
		}
	}
	gridSize = len(lines)
	uniqueNeighbours := map[util.Point]bool{start: true}
	for i := 0; i < 64; i++ {
		newNeighbours := make([]util.Point, 0)
		for n := range uniqueNeighbours {
			newNeighbours = append(newNeighbours, getNeighbours(n)...)
		}
		uniqueNeighbours = make(map[util.Point]bool)
		for _, n := range newNeighbours {
			uniqueNeighbours[n] = true
		}
	}
	return len(uniqueNeighbours)
}

func solve2023Day21Part2(lines []string) interface{} {
	start := util.Point{X: 0, Y: 0}
	for y, row := range lines {
		for x, c := range strings.Split(row, "") {
			if c == "S" {
				c = "."
				start = util.Point{X: x, Y: y}
			}
			grid[util.Point{X: x, Y: y}] = c
		}
	}
	gridSize = len(lines)
	rem := 26501365 % gridSize
	uniqueNeighbours := make(map[util.Point]bool)
	sequence := make([]int, 0)

	uniqueNeighbours = map[util.Point]bool{start: true}
	for i := 0; i < 2*gridSize+rem; i++ {
		newNeighbours := make([]util.Point, 0)
		for n := range uniqueNeighbours {
			newNeighbours = append(newNeighbours, getNeighbours(n)...)
		}
		uniqueNeighbours = make(map[util.Point]bool)
		for _, n := range newNeighbours {
			uniqueNeighbours[n] = true
		}
		if i+1 == rem || i+1 == rem+gridSize || i+1 == rem+2*gridSize {
			sequence = append(sequence, len(uniqueNeighbours))
		}
	}

	//ax^2 + bx + c = sequence[x]
	//a(1)^2 + b(1) + c = sequence[1] = a + b + c
	//a + b = sequence[1] - c
	//a(2)^2 + b(2) + c = sequence[2] = 4a + 2b + c
	//4a + 2b = sequence[2] - c
	//4a + 2b - 2(a + b) = 2a = sequence[2] - c - 2(sequence[1] - c)
	//a = (sequence[2] - c - 2(sequence[1] - c))/2
	c := sequence[0]
	a := (sequence[2] - c - 2*(sequence[1]-c)) / 2
	b := sequence[1] - c - a

	n := 26501365 / gridSize

	return a*(n*n) + b*n + c
}
