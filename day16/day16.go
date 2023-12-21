package day16

import (
	"math"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 16, solve2023Day16Part1, solve2023Day16Part2)
}

var directions = map[string]util.Point{
	"RIGHT": {X: 1, Y: 0},
	"LEFT":  {X: -1, Y: 0},
	"UP":    {X: 0, Y: -1},
	"DOWN":  {X: 0, Y: 1},
}

type beamType struct {
	point     util.Point
	direction string
}

func solve2023Day16Part1(lines []string) interface{} {
	grid := make(map[util.Point]string)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			grid[util.Point{X: x, Y: y}] = c
		}
	}
	energized := make(map[util.Point]string)
	beams := []beamType{{util.Point{X: 0, Y: 0}, "RIGHT"}}

	for len(beams) > 0 {
		beamPoint := beams[0].point
		direction := beams[0].direction
		if d, ok := energized[beamPoint]; ok && d == direction {
			beams = beams[1:]
			continue
		}
		if _, ok := grid[beamPoint]; ok {
			energized[beamPoint] = direction
		}

		if grid[beamPoint] == "." {
			beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
		} else if grid[beamPoint] == "\\" {
			if direction == "RIGHT" {
				direction = "DOWN"
			} else if direction == "LEFT" {
				direction = "UP"
			} else if direction == "UP" {
				direction = "LEFT"
			} else if direction == "DOWN" {
				direction = "RIGHT"
			}
			beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
		} else if grid[beamPoint] == "/" {
			if direction == "RIGHT" {
				direction = "UP"
			} else if direction == "LEFT" {
				direction = "DOWN"
			} else if direction == "UP" {
				direction = "RIGHT"
			} else if direction == "DOWN" {
				direction = "LEFT"
			}
			beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
		} else if grid[beamPoint] == "|" {
			if direction == "RIGHT" || direction == "LEFT" {
				direction = "UP"
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
				direction = "DOWN"
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			} else {
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			}
		} else if grid[beamPoint] == "-" {
			if direction == "UP" || direction == "DOWN" {
				direction = "LEFT"
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
				direction = "RIGHT"
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			} else {
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			}
		}

		beams = beams[1:]
	}

	return len(energized)
}

func solve2023Day16Part2(lines []string) interface{} {
	grid := make(map[util.Point]string)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			grid[util.Point{X: x, Y: y}] = c
		}
	}
	startPoints := make(map[util.Point]string)
	for x := 0; x < len(lines[0]); x++ {
		startPoints[util.Point{X: x, Y: 0}] = "DOWN"
		startPoints[util.Point{X: x, Y: len(lines) - 1}] = "UP"
	}
	for y := 0; y < len(lines[0]); y++ {
		startPoints[util.Point{X: 0, Y: y}] = "RIGHT"
		startPoints[util.Point{X: len(lines[0]) - 1, Y: y}] = "LEFT"
	}
	max := math.MinInt
	for start, startDir := range startPoints {
		energized := make(map[util.Point]string)
		beams := []beamType{{start, startDir}}

		for len(beams) > 0 {
			beamPoint := beams[0].point
			direction := beams[0].direction
			if d, ok := energized[beamPoint]; ok && d == direction {
				beams = beams[1:]
				continue
			}
			if _, ok := grid[beamPoint]; ok {
				energized[beamPoint] = direction
			}

			if grid[beamPoint] == "." {
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			} else if grid[beamPoint] == "\\" {
				if direction == "RIGHT" {
					direction = "DOWN"
				} else if direction == "LEFT" {
					direction = "UP"
				} else if direction == "UP" {
					direction = "LEFT"
				} else if direction == "DOWN" {
					direction = "RIGHT"
				}
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			} else if grid[beamPoint] == "/" {
				if direction == "RIGHT" {
					direction = "UP"
				} else if direction == "LEFT" {
					direction = "DOWN"
				} else if direction == "UP" {
					direction = "RIGHT"
				} else if direction == "DOWN" {
					direction = "LEFT"
				}
				beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
			} else if grid[beamPoint] == "|" {
				if direction == "RIGHT" || direction == "LEFT" {
					direction = "UP"
					beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
					direction = "DOWN"
					beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
				} else {
					beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
				}
			} else if grid[beamPoint] == "-" {
				if direction == "UP" || direction == "DOWN" {
					direction = "LEFT"
					beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
					direction = "RIGHT"
					beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
				} else {
					beams = append(beams, beamType{beams[0].point.Add(directions[direction]), direction})
				}
			}

			beams = beams[1:]
		}

		if len(energized) > max {
			max = len(energized)
		}
	}
	return max
}
