package day8

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 8, solve2023Day8Part1, solve2023Day8Part2)
}

func solve2023Day8Part1(lines []string) interface{} {
	instructions := lines[0]
	lines = lines[2:]
	nodes := make(map[string]map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		name := parts[0]
		destinations := strings.Split(parts[1][1:len(parts[1])-1], ", ")
		left := destinations[0]
		right := destinations[1]
		nodes[name] = map[string]string{
			"L": left,
			"R": right,
		}
	}
	count := 0
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		for _, instruction := range strings.Split(instructions, "") {
			count++
			currentNode = nodes[currentNode][instruction]
			if currentNode == "ZZZ" {
				break
			}
		}
	}
	return count
}

func solve2023Day8Part2(lines []string) interface{} {
	instructions := lines[0]
	lines = lines[2:]
	nodes := make(map[string]map[string]string)
	startingNodes := make([]string, 0)
	endNodes := make(map[string]bool)
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		name := parts[0]
		destinations := strings.Split(parts[1][1:len(parts[1])-1], ", ")
		left := destinations[0]
		right := destinations[1]
		nodes[name] = map[string]string{
			"L": left,
			"R": right,
		}
		if strings.HasSuffix(name, "A") {
			startingNodes = append(startingNodes, name)
		}
		if strings.HasSuffix(name, "Z") {
			endNodes[name] = true
		}
	}
	steps := make([]int, 0)
	retVal := 1
	for i := 0; i < len(startingNodes); i++ {
		count := 0
		stop := false
		for !stop {
			for _, instruction := range strings.Split(instructions, "") {
				count++
				startingNodes[i] = nodes[startingNodes[i]][instruction]
				if endNodes[startingNodes[i]] {
					stop = true
					break
				}
			}
		}
		steps = append(steps, count)
		retVal *= (count)
	}
	return util.LeastCommonMultiple(steps[0], steps[1], steps[2:]...)
}
