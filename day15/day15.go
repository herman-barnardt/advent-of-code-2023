package day15

import (
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 15, solve2023Day15Part1, solve2023Day15Part2)
}

func solve2023Day15Part1(lines []string) interface{} {
	line := lines[0]
	sum := 0
	for _, step := range strings.Split(line, ",") {
		currentValue := 0
		for _, c := range step {
			currentValue += int(c)
			currentValue *= 17
			currentValue = currentValue % 256
		}
		sum += currentValue
	}
	return sum
}

type lens struct {
	label       string
	focalLength int
}

func solve2023Day15Part2(lines []string) interface{} {
	line := lines[0]
	boxes := make(map[int][]*lens, 256)
	for _, step := range strings.Split(line, ",") {
		boxNumber := 0
		label := ""
		focalLengthString := ""
		var operation rune
		for _, c := range step {
			if c == '=' || c == '-' {
				operation = c
			} else if operation != 0 {
				focalLengthString += string(c)
			} else {
				boxNumber += int(c)
				boxNumber *= 17
				boxNumber = boxNumber % 256
				label += string(c)
			}
		}
		if operation == '=' {
			focalLength, _ := strconv.Atoi(focalLengthString)
			foundIndex := -1
			for index, lens := range boxes[boxNumber] {
				if lens.label == label {
					foundIndex = index
				}
			}
			if foundIndex == -1 {
				boxes[boxNumber] = append(boxes[boxNumber], &lens{label, focalLength})
			} else {
				boxes[boxNumber][foundIndex].focalLength = focalLength
			}
		} else {
			deleteIndex := -1
			for index, lens := range boxes[boxNumber] {
				if lens.label == label {
					deleteIndex = index
				}
			}
			if deleteIndex == 0 {
				boxes[boxNumber] = boxes[boxNumber][1:]
			} else if deleteIndex > 0 {
				boxes[boxNumber] = append(boxes[boxNumber][:deleteIndex], boxes[boxNumber][deleteIndex+1:]...)

			}
		}
	}
	sum := 0
	for boxNumber, box := range boxes {
		for index, lens := range box {
			sum += (boxNumber + 1) * (index + 1) * lens.focalLength
		}
	}
	return sum
}
