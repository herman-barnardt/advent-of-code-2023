package day11

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 11, solve2023Day11Part1, solve2023Day11Part2)
}

func solve2023Day11Part1(lines []string) interface{} {
	newLines := make([]string, 0)
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			newLines = append(newLines, line)
		}
		newLines = append(newLines, line)
	}
	emptyCols := make([]int, 0)
	for x := 0; x < len(lines[0]); x++ {
		contains := false
		for _, line := range lines {
			contains = string(line[x]) == "#"
			if contains {
				break
			}
		}
		if !contains {
			emptyCols = append(emptyCols, x)
		}
	}
	for y := range newLines {
		for i, x := range emptyCols {
			newLines[y] = newLines[y][:x+i] + "." + newLines[y][x+i:]
		}
	}
	points := make([]util.Point, 0)
	lineMap := util.LinesToMap(newLines)
	for y, row := range lineMap {
		for x, c := range row {
			if c == "#" {
				points = append(points, util.Point{X: x, Y: y})
			}
		}
	}
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < (len(points)); j++ {
			distance := util.DistanceBetween(&points[i], &points[j])
			sum += distance
		}
	}
	return sum
}

func solve2023Day11Part2(lines []string) interface{} {
	emptyRows := make([]int, 0)
	for i, line := range lines {
		if !strings.Contains(line, "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	emptyCols := make([]int, 0)
	for x := 0; x < len(lines[0]); x++ {
		contains := false
		for _, line := range lines {
			contains = string(line[x]) == "#"
			if contains {
				break
			}
		}
		if !contains {
			emptyCols = append(emptyCols, x)
		}
	}
	mult := 1000000 - 1
	points := make([]util.Point, 0)
	lineMap := util.LinesToMap(lines)
	currentEmptyRow := 0
	for y := 0; y < len(lineMap); y++ {
		if currentEmptyRow < len(emptyRows) && emptyRows[currentEmptyRow]+1 == y {
			currentEmptyRow++
		}
		currentEmptyCol := 0
		for x := 0; x < len(lineMap[y]); x++ {
			if currentEmptyCol < len(emptyCols) && emptyCols[currentEmptyCol]+1 == x {
				currentEmptyCol++
			}
			if lineMap[y][x] == "#" {
				points = append(points, util.Point{X: x + (currentEmptyCol * mult), Y: y + (currentEmptyRow * mult)})
			}
		}
	}
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < (len(points)); j++ {
			distance := util.DistanceBetween(&points[i], &points[j])
			sum += distance
		}
	}
	return sum
}
