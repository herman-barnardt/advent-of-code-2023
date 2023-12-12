package day12

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 12, solve2023Day12Part1, solve2023Day12Part2)
}

type key struct {
	position         int
	currentHashCount int
	groupIndex       int
}

func getCount(memo map[key]int, line string, groups []int, position int, currentHashCount int, groupIndex int) int {
	key := key{position, currentHashCount, groupIndex}
	if v, ok := memo[key]; ok {
		return v
	}
	retVal := 0
	if position == len(line) {
		// we are at the end of the line and everything is valid
		if len(groups) == groupIndex {
			retVal = 1
		}
	} else if line[position] == '#' {
		// increase our hash count and continue
		retVal = getCount(memo, line, groups, position+1, currentHashCount+1, groupIndex)
	} else if line[position] == '.' || groupIndex == len(groups) {
		if groupIndex < len(groups) && currentHashCount == groups[groupIndex] {
			//we have completed one group, move on to the next
			retVal = getCount(memo, line, groups, position+1, 0, groupIndex+1)
		} else if currentHashCount == 0 {
			//we are searching for a new group
			retVal = getCount(memo, line, groups, position+1, 0, groupIndex)
		} else {
			//this can't be a . this iteration is not valid
			retVal = 0
		}
	} else {
		//how many valid options do we have if this is a #
		hashCount := getCount(memo, line, groups, position+1, currentHashCount+1, groupIndex)
		dotCount := 0
		//how many valid options do we have if this is a .
		if currentHashCount == groups[groupIndex] {
			//we have completed a group
			dotCount = getCount(memo, line, groups, position+1, 0, groupIndex+1)
		} else if currentHashCount == 0 {
			//we are searching for a new group
			dotCount = getCount(memo, line, groups, position+1, 0, groupIndex)
		}
		retVal = hashCount + dotCount
	}
	memo[key] = retVal
	return retVal
}

func solve2023Day12Part1(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		record := parts[0] + "."
		groups := util.StringToIntSlice(parts[1], ",")
		sum += getCount(make(map[key]int), record, groups, 0, 0, 0)
	}
	return sum
}

func solve2023Day12Part2(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		record := parts[0]
		groups := util.StringToIntSlice(parts[1], ",")
		record += "?"
		newRecord := ""
		newGroup := make([]int, 0)
		for i := 0; i < 5; i++ {
			newRecord += record
			newGroup = append(newGroup, groups...)
		}
		newRecord = newRecord[:len(newRecord)-1] + "."
		sum += getCount(make(map[key]int), newRecord, newGroup, 0, 0, 0)
	}
	return sum
}
