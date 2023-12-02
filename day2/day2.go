package day2

import (
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 2, solve2023Day2Part1, solve2023Day2Part2)
}

func solve2023Day2Part1(lines []string) interface{} {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
	for gameId, l := range lines {
		l = l[strings.Index(l, ":")+1:]
		rounds := strings.Split(l, ";")
		validGame := true
		for _, round := range rounds {
			round = strings.Trim(round, " ")
			cubes := strings.Split(round, ",")
			validRound := true
			for _, c := range cubes {
				c = strings.Trim(c, " ")
				spaceIndex := strings.Index(c, " ")
				count, _ := strconv.Atoi(string(c[:spaceIndex]))
				colour := strings.Trim(c[spaceIndex:], " ")
				validRound = validRound && maxCubes[colour] >= count
			}
			validGame = validGame && validRound
		}
		if validGame {
			sum += gameId + 1
		}
	}
	return sum
}

func solve2023Day2Part2(lines []string) interface{} {
	sum := 0
	for _, l := range lines {
		l = l[strings.Index(l, ":")+1:]
		rounds := strings.Split(l, ";")
		maxCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, round := range rounds {
			round = strings.Trim(round, " ")
			cubes := strings.Split(round, ",")
			for _, c := range cubes {
				c = strings.Trim(c, " ")
				spaceIndex := strings.Index(c, " ")
				count, _ := strconv.Atoi(string(c[:spaceIndex]))
				colour := strings.Trim(c[spaceIndex:], " ")
				if count > maxCubes[colour] {
					maxCubes[colour] = count
				}
			}
		}
		sum += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
	}
	return sum
}
