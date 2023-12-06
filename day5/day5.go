package day5

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 5, solve2023Day5Part1, solve2023Day5Part2)
}

func solve2023Day5Part1(lines []string) interface{} {
	parts := strings.Split(lines[0], " ")
	seeds := make([]int, 0)
	for _, seed := range parts[1:] {
		seedNum, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedNum)
	}
	changed := make(map[int]bool)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			destinationStart, _ := strconv.Atoi(parts[0])
			sourceStart, _ := strconv.Atoi(parts[1])
			rangeCount, _ := strconv.Atoi(parts[2])
			for i, seed := range seeds {
				if _, ok := changed[i]; !ok {
					if seed >= sourceStart && seed < sourceStart+rangeCount {
						seeds[i] = destinationStart + seed - sourceStart
						changed[i] = true
					}
				}

			}
		}
		if len(line) == 0 {
			changed = map[int]bool{}
		}
	}
	slices.Sort(seeds)
	return seeds[0]
}

type seedRange struct {
	start int
	end   int
}

type almanacMap struct {
	destinationStart int
	sourceStart      int
	mapRange         int
}

func solve2023Day5Part2(lines []string) interface{} {
	parts := strings.Split(lines[0], " ")
	seeds := make([]seedRange, 0)
	i := 1
	for i < len(parts[1:]) {
		seedNum, _ := strconv.Atoi(parts[i])
		rangeCount, _ := strconv.Atoi(parts[i+1])
		seeds = append(seeds, seedRange{seedNum, seedNum + rangeCount})
		i += 2
	}
	maps := make([][]almanacMap, 0)
	maps = append(maps, make([]almanacMap, 0))
	index := 0
	for _, line := range lines[3:] {
		parts := strings.Split(line, " ")
		if len(line) == 0 {
			maps = append(maps, make([]almanacMap, 0))
		} else if len(parts) == 3 {
			destinationStart, _ := strconv.Atoi(parts[0])
			sourceStart, _ := strconv.Atoi(parts[1])
			rangeCount, _ := strconv.Atoi(parts[2])
			maps[index] = append(maps[index], almanacMap{destinationStart, sourceStart, rangeCount})
		} else {
			index++
		}
	}

	for _, group := range maps {
		updatedSeeds := make([]seedRange, 0)
		for _, almanacMap := range group {
			sourceEnd := almanacMap.sourceStart + almanacMap.mapRange
			newSeedRanges := make([]seedRange, 0)
			for _, seed := range seeds {
				rangeBefore := seedRange{seed.start, util.IntMin(seed.end, almanacMap.sourceStart)}
				rangeIncluded := seedRange{util.IntMax(seed.start, almanacMap.sourceStart), util.IntMin(seed.end, sourceEnd)}
				rangeAfter := seedRange{util.IntMax(sourceEnd, seed.start), seed.end}
				if rangeBefore.end > rangeBefore.start {
					newSeedRanges = append(newSeedRanges, rangeBefore)
				}
				if rangeIncluded.end > rangeIncluded.start {
					updatedSeeds = append(updatedSeeds, seedRange{rangeIncluded.start - almanacMap.sourceStart + almanacMap.destinationStart, rangeIncluded.end - almanacMap.sourceStart + almanacMap.destinationStart})
				}
				if rangeAfter.end > rangeAfter.start {
					newSeedRanges = append(newSeedRanges, rangeAfter)
				}
			}
			seeds = newSeedRanges
		}
		seeds = append(updatedSeeds, seeds...)
	}

	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].start < seeds[j].start
	})
	return seeds[0].start
}
