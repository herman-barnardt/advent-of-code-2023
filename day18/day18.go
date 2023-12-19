package day18

import (
	"fmt"
	"strconv"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 18, solve2023Day18Part1, solve2023Day18Part2)
}

var directions = map[string]util.Point{
	"R": {X: 1, Y: 0},
	"0": {X: 1, Y: 0},
	"L": {X: -1, Y: 0},
	"2": {X: -1, Y: 0},
	"D": {X: 0, Y: 1},
	"1": {X: 0, Y: 1},
	"U": {X: 0, Y: -1},
	"3": {X: 0, Y: -1},
}

func solve2023Day18Part1(lines []string) interface{} {
	start := util.Point{X: 0, Y: 0}
	points := []util.Point{start}
	perimeter := 0
	for _, line := range lines {
		var direction string
		var number int64
		var hex string
		fmt.Sscanf(line, "%s %d %s", &direction, &number, &hex)
		start = start.Add(util.Point{X: directions[direction].X * int(number), Y: directions[direction].Y * int(number)})
		points = append(points, start)
		perimeter += int(number)
	}
	a := 0
	for i := 0; i < len(points)-1; i++ {
		a += (points[i].Y + points[i+1].Y) * (points[i].X - points[i+1].X)
	}
	return perimeter/2 + a/2 + 1
}

func solve2023Day18Part2(lines []string) interface{} {
	start := util.Point{X: 0, Y: 0}
	points := []util.Point{start}
	perimeter := 0
	for _, line := range lines {
		var direction string
		var number int64
		var hex string
		fmt.Sscanf(line, "%s %d %s", &direction, &number, &hex)
		hex = hex[2 : len(hex)-1]
		direction = string(hex[len(hex)-1])
		hex = hex[:len(hex)-1]
		number, _ = strconv.ParseInt(hex, 16, 64)
		start = start.Add(util.Point{X: directions[direction].X * int(number), Y: directions[direction].Y * int(number)})
		points = append(points, start)
		perimeter += int(number)
	}
	a := 0
	for i := 0; i < len(points)-1; i++ {
		a += (points[i].Y + points[i+1].Y) * (points[i].X - points[i+1].X)
	}
	return perimeter/2 + a/2 + 1
}
