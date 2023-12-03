package day3

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 3, solve2023Day3Part1, solve2023Day3Part2)
}

func solve2023Day3Part1(lines []string) interface{} {
	sum := 0
	schematic := make(map[int]map[int]string, 0)
	maxY := len(lines)
	maxX := len(lines[0])
	for y, line := range lines {
		schematic[y] = make(map[int]string)
		for x, c := range strings.Split(line, "") {
			schematic[y][x] = c
		}
	}
	for y := 0; y <= maxY; y++ {
		numString := ""
		isValid := false
		for x := 0; x <= maxX; x++ {
			_, err := strconv.Atoi(schematic[y][x])
			if err == nil {
				numString += schematic[y][x]
				for i := y - 1; i <= y+1; i++ {
					for j := x - 1; j <= x+1; j++ {
						_, err = strconv.Atoi(schematic[i][j])
						isValid = isValid || (err != nil && schematic[i][j] != "." && schematic[i][j] != "")
					}
				}
			}
			_, err = strconv.Atoi(schematic[y][x])
			if schematic[y][x] == "." || x == maxX-1 || (err != nil && schematic[y][x] != "") {
				if isValid {
					fmt.Println("Valid", numString)
					val, _ := strconv.Atoi(numString)
					sum += val
				}
				numString = ""
				isValid = false
			}
		}
	}
	return sum
}

func solve2023Day3Part2(lines []string) interface{} {
	sum := 0
	schematic := make(map[int]map[int]string, 0)
	maxY := len(lines)
	maxX := len(lines[0])
	for y, line := range lines {
		schematic[y] = make(map[int]string)
		for x, c := range strings.Split(line, "") {
			schematic[y][x] = c
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if schematic[y][x] == "*" {
				gearNums := make([]int, 0)
				for i := y - 1; i <= y+1; i++ {
					for j := x - 1; j <= x+1; j++ {
						_, err := strconv.Atoi(schematic[i][j])
						if err == nil {
							numString := schematic[i][j]
							leftSearch := true
							leftIndex := j - 1
							for leftSearch {
								_, err = strconv.Atoi(schematic[i][leftIndex])
								if err == nil {
									numString = schematic[i][leftIndex] + numString
									schematic[i][leftIndex] = "."
									leftIndex--
								}
								if leftIndex < 0 || schematic[i][leftIndex] == "." || schematic[i][leftIndex] == "" || schematic[i][leftIndex] == "*" {
									leftSearch = false
								}

							}
							rightSearch := true
							rightIndex := j + 1
							for rightSearch {
								_, err = strconv.Atoi(schematic[i][rightIndex])
								if err == nil {
									numString = numString + schematic[i][rightIndex]
									schematic[i][rightIndex] = "."
									rightIndex++
								}
								if rightIndex > maxX || schematic[i][rightIndex] == "." || schematic[i][rightIndex] == "" || schematic[i][rightIndex] == "*" {
									rightSearch = false
								}

							}
							num, _ := strconv.Atoi(numString)
							gearNums = append(gearNums, num)
						}
					}
				}
				if len(gearNums) >= 2 {
					sum += gearNums[0] * gearNums[1]
				}
			}
		}
	}
	return sum
}
