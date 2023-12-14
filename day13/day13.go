package day13

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 13, solve2023Day13Part1, solve2023Day13Part2)
}

func solve2023Day13Part1(lines []string) interface{} {
	sum := 0
	patterns := util.SplitLines(lines, "")
	for _, pattern := range patterns {
		horizontalSymmetry := getHorizontalSymmetry(pattern, false, -1)

		transpose := transposeStrings(pattern)

		verticalSymmetry := getHorizontalSymmetry(transpose, false, -1)

		sum += (verticalSymmetry + 1) + (horizontalSymmetry+1)*100
	}
	return sum
}

func solve2023Day13Part2(lines []string) interface{} {
	sum := 0
	patterns := util.SplitLines(lines, "")
	for _, pattern := range patterns {
		oldHorizontalSymmetry := getHorizontalSymmetry(pattern, false, -1)
		horizontalSymmetry := getHorizontalSymmetry(pattern, true, oldHorizontalSymmetry)

		transpose := transposeStrings(pattern)

		oldVerticalSymmetry := getHorizontalSymmetry(transpose, false, -1)
		verticalSymmetry := getHorizontalSymmetry(transpose, true, oldVerticalSymmetry)

		sum += (verticalSymmetry + 1) + (horizontalSymmetry+1)*100
	}
	return sum
}

func getHorizontalSymmetry(pattern []string, allowSmudge bool, notAllowed int) int {
	for i := 0; i < len(pattern)-1; i++ {
		index1 := i
		index2 := i + 1
		mismatch := false
		done := false

		for index1 >= 0 && index2 < len(pattern) {
			diffCount := getDifference(pattern[index1], pattern[index2])
			if pattern[index1] != pattern[index2] && (!allowSmudge || diffCount > 1 || done) {
				mismatch = true
				break
			}
			if diffCount == 1 {
				done = true
			}

			index1 -= 1
			index2 += 1
		}

		if !mismatch && i != notAllowed {
			return i
		}
	}

	return -1
}

func transposeStrings(lines []string) []string {
	transpose := make([]string, len(lines[0]))
	for _, line := range lines {
		for i, c := range strings.Split(line, "") {
			transpose[i] += c
		}
	}
	return transpose
}

func getDifference(s1, s2 string) int {
	diffCount := len(s1)
	for k := 0; k < len(s1); k++ {
		if s1[k] == s2[k] {
			diffCount--
		}
	}
	return diffCount
}
