package day4

import (
	"math"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 4, solve2023Day4Part1, solve2023Day4Part2)
}

func solve2023Day4Part1(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		winningNumbers := make(map[int]int)
		line = line[strings.Index(line, ":")+1:]
		parts := strings.Split(line, "|")
		winningNumberStrings := strings.Split(strings.Trim(parts[0], " "), " ")
		for _, str := range winningNumberStrings {
			num, err := strconv.Atoi(str)
			if err == nil {
				winningNumbers[num] = 0
			}
		}
		numberStrings := strings.Split(strings.Trim(parts[1], " "), " ")
		for _, str := range numberStrings {
			num, _ := strconv.Atoi(str)
			if _, ok := winningNumbers[num]; ok {
				winningNumbers[num]++
			}
		}
		matchCount := 0
		for _, v := range winningNumbers {
			if v > 0 {
				matchCount++
			}
		}
		if matchCount > 0 {
			sum += int(math.Pow(2, float64(matchCount-1)))
		}
	}
	return sum
}

func solve2023Day4Part2(lines []string) interface{} {
	cards := make(map[int]int)
	for cardNumber, line := range lines {
		cards[cardNumber+1]++
		winningNumbers := make(map[int]int)
		line = line[strings.Index(line, ":")+1:]
		parts := strings.Split(line, "|")
		winningNumberStrings := strings.Split(strings.Trim(parts[0], " "), " ")
		for _, str := range winningNumberStrings {
			num, err := strconv.Atoi(str)
			if err == nil {
				winningNumbers[num] = 0
			}
		}
		numberStrings := strings.Split(strings.Trim(parts[1], " "), " ")
		for _, str := range numberStrings {
			num, _ := strconv.Atoi(str)
			if _, ok := winningNumbers[num]; ok {
				winningNumbers[num]++
			}
		}
		matchCount := 0
		for _, v := range winningNumbers {
			if v > 0 {
				matchCount++
			}
		}
		for i := 1; i <= matchCount; i++ {
			cards[cardNumber+1+i] += cards[cardNumber+1]
		}

	}
	sum := 0
	for _, v := range cards {
		sum += v
	}
	return sum
}
