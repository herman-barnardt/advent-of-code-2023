package day9

import (
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 9, solve2023Day9Part1, solve2023Day9Part2)
}

func diffSlice(sequence []int) (diffSlice []int, zeroes bool) {
	sum := 0
	for i := 1; i < len(sequence); i++ {
		n := sequence[i] - sequence[i-1]
		diffSlice = append(diffSlice, n)
		sum += n
	}
	return diffSlice, sum == 0
}

func solve2023Day9Part1(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		sequences := make([][]int, 0)
		sequence := make([]int, 0)
		for _, n := range parts {
			num, _ := strconv.Atoi(n)
			sequence = append(sequence, num)
		}
		index := 0
		sequences = append(sequences, sequence)
		diff, zeroed := diffSlice(sequences[index])
		for !zeroed {
			sequences = append(sequences, diff)
			index++
			diff, zeroed = diffSlice(sequences[index])
		}

		for i := index; i > 0; i-- {
			lastDelta := sequences[i][len(sequences[i])-1]
			sequences[i-1] = append(sequences[i-1], sequences[i-1][len(sequences[i-1])-1]+lastDelta)
		}
		sum += sequences[0][len(sequences[0])-1]
	}
	return sum
}

func solve2023Day9Part2(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		sequences := make([][]int, 0)
		sequence := make([]int, 0)
		for _, n := range parts {
			num, _ := strconv.Atoi(n)
			sequence = append(sequence, num)
		}
		index := 0
		sequences = append(sequences, sequence)
		diff, zeroed := diffSlice(sequences[index])
		for !zeroed {
			sequences = append(sequences, diff)
			index++
			diff, zeroed = diffSlice(sequences[index])
		}

		for i := index; i > 0; i-- {
			firstDelta := sequences[i][0]
			sequences[i-1] = append([]int{sequences[i-1][0] - firstDelta}, sequences[i-1]...)
		}
		sum += sequences[0][0]
	}
	return sum
}
