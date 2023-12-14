package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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
	lines := strings.Split(input, "\r\n")
	timesStr := lines[0][5:]
	distancesStr := lines[1][9:]
	re := regexp.MustCompile(`\s*(\d+)`)
	timesMatches := re.FindAllStringSubmatch(timesStr, -1)
	distancesMatches := re.FindAllStringSubmatch(distancesStr, -1)

	result := 1
	for i := 0; i < len(timesMatches); i++ {
		timeStr := timesMatches[i][1]
		distanceStr := distancesMatches[i][1]

		time, err := strconv.Atoi(timeStr)
		if err != nil {
			log.Fatalln("Failed to parse time: ", err)
		}
		recordDistance, err := strconv.Atoi(distanceStr)
		if err != nil {
			log.Fatalln("Failed to parse distance: ", err)
		}

		numWays := calculateNumWays(time, recordDistance)
		result *= numWays
	}

	fmt.Println("Result: ", result)
}

func part2(input string) {
	panic("todo")
}

func calculateNumWays(time int, recordDistance int) int {
	numWays := 0

	for i := 1; i < time; i++ {
		timeRemaining := time - i
		distanceTraveled := i * timeRemaining
		if distanceTraveled > recordDistance {
			numWays += 1
		}
	}

	return numWays
}
