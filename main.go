package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	isTestPtr := flag.Bool("test", false, "Whether to use test_input.txt or input.txt")
	part := flag.Int("part", 1, "Which part to run: 1 or 2")

	flag.Parse()

	inputFileName := "input.txt"
	if *isTestPtr {
		inputFileName = fmt.Sprintf("test_input_part%v.txt", *part)
	}

	input, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalln("Error reading input: ", err)
	}

	if *part == 1 {
		part1(string(input))

	} else {
		part2(string(input))
	}
}

func part1(input string) {
	sections := strings.Split(input, "\r\n\r\n")
	seedsStr := sections[0]
	seedsStr = seedsStr[7:]
	seeds := []int{}
	for _, seedStr := range strings.Split(seedsStr, " ") {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatalln("Failed to parse seed: ", err)
		}
		seeds = append(seeds, seed)
	}
	items := seeds
	for sectionIdx := 1; sectionIdx < len(sections); sectionIdx++ {
		section := sections[sectionIdx]
		ranges := []NumberRange{}
		sectionLines := strings.Split(section, "\r\n")
		// ignore title
		for i := 1; i < len(sectionLines); i++ {
			ranges = append(ranges, parseRange(sectionLines[i]))
		}
		// map current items into their currosponding next item
		for itemIdx := 0; itemIdx < len(items); itemIdx++ {
			newItem := mapItemWithRanges(items[itemIdx], ranges)
			items[itemIdx] = newItem
		}
	}

	minLocation := slices.Min(items)
	fmt.Println("Result: ", minLocation)
}

func part2(input string) {
	panic("todo")
}

type NumberRange struct {
	srcStart  int
	destStart int
	length    int
}

func (r *NumberRange) Contains(item int) bool {
	return item >= r.srcStart && item <= r.srcStart+r.length
}

func (r *NumberRange) Map(item int) int {
	diff := r.destStart - r.srcStart
	return item + diff
}

func parseRange(rangeStr string) NumberRange {
	rangeComponents := strings.Split(rangeStr, " ")
	destStartStr := rangeComponents[0]
	destStart, err := strconv.Atoi(destStartStr)
	if err != nil {
		log.Fatalln("Failed to parse destination range start: ", err)
	}
	srcStartStr := rangeComponents[1]
	srcStart, err := strconv.Atoi(srcStartStr)
	if err != nil {
		log.Fatalln("Failed to parse source range start: ", err)
	}
	lengthStr := rangeComponents[2]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		log.Fatalln("Failed to parse range length: ", err)
	}

	return NumberRange{
		srcStart:  srcStart,
		destStart: destStart,
		length:    length,
	}
}

func mapItemWithRanges(item int, ranges []NumberRange) int {
	for _, numRange := range ranges {
		if numRange.Contains(item) {
			return numRange.Map(item)
		}
	}

	return item
}
