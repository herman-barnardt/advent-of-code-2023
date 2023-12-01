package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"github.com/herman-barnardt/aoc"
)

func main() { 
	flag.Parse()

	command := flag.Arg(0)
	yearString := flag.Arg(1)
	year := 2023
	var err error
	if len(yearString) > 0 {
		year, err = strconv.Atoi(yearString)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}
	day := 0
	dayString := flag.Arg(2)
	if len(dayString) > 0 {
		day, err = strconv.Atoi(dayString)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}
	part := 0
	partString := flag.Arg(3)
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