package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Drumstickz64/advent_of_code_2023/utils"
)

type Command int

const (
	COMMAND_LEFT Command = iota
	COMMAND_RIGHT
)

type Tree map[string][2]string

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
	cmds, tree := parseInput(input)
	currNode := "AAA"
	steps := 0
	for i := 0; true; i = (i + 1) % len(cmds) {
		if currNode == "ZZZ" {
			break
		}

		cmd := cmds[i]
		currNode = tree[currNode][cmd]
		steps++
	}

	fmt.Println("Result: ", steps)
}

func part2(input string) {
	panic("todo")
}

func parseInput(input string) ([]Command, Tree) {
	sections := strings.Split(input, "\r\n\r\n")
	cmdSection := sections[0]
	cmds := []Command{}
	for _, char := range cmdSection {
		if char == 'L' {
			cmds = append(cmds, COMMAND_LEFT)
		} else {
			cmds = append(cmds, COMMAND_RIGHT)
		}
	}

	nodesSection := sections[1]
	tree := Tree{}
	re := regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)
	for _, nodeStr := range utils.Lines(nodesSection) {
		match := re.FindStringSubmatch(nodeStr)
		nodeName := match[1]
		leftNodeName := match[2]
		rightNodeName := match[3]
		tree[nodeName] = [2]string{leftNodeName, rightNodeName}
	}

	return cmds, tree
}
