package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	panic("todo")
}

func part2(input string) {
	panic("todo")
}
