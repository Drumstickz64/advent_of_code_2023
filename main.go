package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Drumstickz64/advent_of_code_2023/utils"
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
	sequences := parseInput(input)
	sum := 0
	for _, sequence := range sequences {
		sum += predictNextValue(sequence)
	}

	fmt.Println("Result: ", sum)
}

func part2(input string) {
	panic("todo")
}

func parseInput(input string) [][]int {
	sequences := [][]int{}
	for lineIdx, line := range utils.Lines(input) {
		valueStrs := strings.Split(line, " ")
		values := []int{}
		for valueIdx, valueStr := range valueStrs {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Fatalf("Failed to parse value %v at position %v on line %v", valueStr, valueIdx, lineIdx)
			}

			values = append(values, value)
		}
		sequences = append(sequences, values)
	}

	return sequences
}

func predictNextValue(sequence []int) int {
	diffArr := []int{}
	diffAllZeros := true
	for i := 0; i < len(sequence)-1; i++ {
		diff := sequence[i+1] - sequence[i]
		if diff != 0 {
			diffAllZeros = false
		}
		diffArr = append(diffArr, diff)
	}

	if diffAllZeros {
		return sequence[len(sequence)-1]
	}

	return sequence[len(sequence)-1] + predictNextValue(diffArr)
}
