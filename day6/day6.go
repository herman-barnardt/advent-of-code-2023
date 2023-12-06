package day6

import (
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 6, solve2023Day6Part1, solve2023Day6Part2)
}

type race struct {
	maxTime int
}

func solve2023Day6Part1(lines []string) interface{} {
	times := make([]int, 0)
	for _, s := range strings.Split(strings.ReplaceAll(lines[0][10:], "  ", " "), " ")[1:] {
		n, err := strconv.Atoi(s)
		if err == nil {
			times = append(times, n)
		}
	}
	distances := make([]int, 0)
	for _, s := range strings.Split(strings.ReplaceAll(lines[1][10:], "  ", " "), " ")[1:] {
		n, err := strconv.Atoi(s)
		if err == nil {
			distances = append(distances, n)
		}
	}
	total := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		count := 0
		for t := 1; t <= time/2; t++ {
			if t*(time-t) > distance {
				count += 2
				if t == time-t {
					count--
				}
			}
		}
		total *= count
	}
	return total
}

func solve2023Day6Part2(lines []string) interface{} {
	time, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(lines[0], " ", ""), ":")[1])
	distance, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(lines[1], " ", ""), ":")[1])
	count := 0
	for t := 1; t <= time/2; t++ {
		if t*(time-t) > distance {
			count += 2
			if t == time-t {
				count--
			}
		}
	}
	return count
}
