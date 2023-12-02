package day1

import (
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 1, solve2023Day1Part1, solve2023Day1Part2)
}

func solve2023Day1Part1(lines []string) interface{} {

	sum := 0
	for _, line := range lines {
		lineNums := make([]string, 0)
		for _, c := range strings.Split(line, "") {
			_, err := strconv.Atoi(c)
			if err == nil {
				lineNums = append(lineNums, c)
			}
		}
		lineVal := lineNums[0] + lineNums[len(lineNums)-1]
		value, err := strconv.Atoi(lineVal)
		if err != nil {
			panic(err)
		}
		sum += value
	}
	return sum
}

func solve2023Day1Part2(lines []string) interface{} {
	sum := 0
	var validNumbers = map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for _, line := range lines {
		lineNums := make([]string, 0)
		for len(line) > 0 {
			for k, v := range validNumbers {
				if strings.HasPrefix(line, k) {
					lineNums = append(lineNums, v)
					break
				}
			}
			line = line[1:]
		}
		lineVal := lineNums[0] + lineNums[len(lineNums)-1]
		value, _ := strconv.Atoi(lineVal)
		sum += value
	}
	return sum
}
