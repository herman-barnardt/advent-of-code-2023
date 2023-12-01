package day1
import (
	"strings"
	"strconv"

	aoc "github.com/herman-barnardt/aoc"
)
func init() {
	aoc.Register(2023, 1, solve2023Day1Part1, solve2023Day1Part2)
}

func solve2023Day1Part1(lines []string) interface{} {
	values := make([]int, 0)
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
		values = append(values, value)
	}
	sum := 0
	for _, v := range values {
		sum += v 
	}	
	return sum
}

func solve2023Day1Part2(lines []string) interface{} {
	values := make([]int, 0)
	for _, line := range lines {
		lineNums := make([]string, 0)
		lineDup := line
		for len(lineDup) > 0 {
			_, err := strconv.Atoi((string)(lineDup[0]))
			if err == nil {
				lineNums = append(lineNums, (string)(lineDup[0]))
			} else if strings.HasPrefix(lineDup, "one") {
				lineNums = append(lineNums, "1")
			} else if strings.HasPrefix(lineDup, "two") {
				lineNums = append(lineNums, "2")
			} else if strings.HasPrefix(lineDup, "three") {
				lineNums = append(lineNums, "3")
			} else if strings.HasPrefix(lineDup, "four") {
				lineNums = append(lineNums, "4")
			} else if strings.HasPrefix(lineDup, "five") {
				lineNums = append(lineNums, "5")
			} else if strings.HasPrefix(lineDup, "six") {
				lineNums = append(lineNums, "6")
			} else if strings.HasPrefix(lineDup, "seven") {
				lineNums = append(lineNums, "7")
			} else if strings.HasPrefix(lineDup, "eight") {
				lineNums = append(lineNums, "8")
			} else if strings.HasPrefix(lineDup, "nine") {
				lineNums = append(lineNums, "9")
			}
			if len(lineDup) > 0 {
				lineDup = lineDup[1:]
			}
		}
		lineVal := lineNums[0] + lineNums[len(lineNums)-1]
		value, err := strconv.Atoi(lineVal)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}
	sum := 0
	for _, v := range values {
		sum += v 
	}	
	return sum
}
