package main

import (
	_ "advent-of-code-2023/day1"
	_ "advent-of-code-2023/day10"
	_ "advent-of-code-2023/day11"
	_ "advent-of-code-2023/day12"
	_ "advent-of-code-2023/day13"
	_ "advent-of-code-2023/day14"
	_ "advent-of-code-2023/day15"
	_ "advent-of-code-2023/day17"
	_ "advent-of-code-2023/day18"
	_ "advent-of-code-2023/day2"
	_ "advent-of-code-2023/day3"
	_ "advent-of-code-2023/day4"
	_ "advent-of-code-2023/day5"
	_ "advent-of-code-2023/day6"
	_ "advent-of-code-2023/day7"
	_ "advent-of-code-2023/day8"
	_ "advent-of-code-2023/day9"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/herman-barnardt/aoc"
)

func main() {
	flag.Parse()

	command := flag.Arg(0)
	year := 2023
	_, _, day := time.Now().Date()
	var err error
	dayString := flag.Arg(1)
	if len(dayString) > 0 && dayString != "0" {
		day, err = strconv.Atoi(dayString)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}
	part := 0
	partString := flag.Arg(2)
	if len(partString) > 0 {
		part, err = strconv.Atoi(partString)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}

	err = aoc.Run(command, year, day, part)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
