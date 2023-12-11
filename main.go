package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	totalPoints := 0
	for _, card := range strings.Split(input, "\r\n") {
		numbers := strings.Split(card, ": ")[1]
		winningAndMyNumbersStr := strings.Split(numbers, " | ")
		winningNumbersStr := winningAndMyNumbersStr[0]
		myNumbersStr := winningAndMyNumbersStr[1]

		winningNumbers := map[int]struct{}{}
		for i := 0; i < len(winningNumbersStr); i += 3 {
			number, err := strconv.Atoi(strings.Trim(winningNumbersStr[i:i+2], " "))
			if err != nil {
				log.Fatalln("Failed to parse winning number: ", err)
			}
			winningNumbers[number] = struct{}{}
		}

		points := 0
		for i := 0; i < len(myNumbersStr); i += 3 {
			number, err := strconv.Atoi(strings.Trim(myNumbersStr[i:i+2], " "))
			if err != nil {
				log.Fatalln("Failed to parse my number: ", err)
			}

			if _, ok := winningNumbers[number]; !ok {
				continue
			}

			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
		totalPoints += points
	}

	fmt.Println("Result: ", totalPoints)
}

func part2(input string) {
	cards := strings.Split(input, "\r\n")
	cardQuantityArr := make([]int, len(cards))
	for cardIdx, card := range cards {
		cardQuantityArr[cardIdx] += 1
		numbersSection := strings.Split(card, ": ")[1]
		winningAndMyNumbers := strings.Split(numbersSection, " | ")
		winningNumbersStr := winningAndMyNumbers[0]
		myNumbersStr := winningAndMyNumbers[1]

		winningNumbers := map[int]struct{}{}
		for i := 0; i < len(winningNumbersStr); i += 3 {
			number, err := strconv.Atoi(strings.TrimLeft(winningNumbersStr[i:i+2], " "))
			if err != nil {
				log.Fatalln("Failed to parse winning number: ", err)
			}
			winningNumbers[number] = struct{}{}
		}

		numbersMatched := 0
		for i := 0; i < len(myNumbersStr); i += 3 {
			number, err := strconv.Atoi(strings.Trim(myNumbersStr[i:i+2], " "))
			if err != nil {
				log.Fatalln("Failed to parse my number: ", err)
			}

			if _, ok := winningNumbers[number]; !ok {
				continue
			}

			copyIdx := cardIdx + numbersMatched + 1
			if copyIdx >= len(cards) {
				break
			}
			cardQuantityArr[copyIdx] += 1 * cardQuantityArr[cardIdx]
			numbersMatched += 1
		}
	}

	totalNumScratchcards := 0
	for _, numScratchcard := range cardQuantityArr {
		totalNumScratchcards += numScratchcard
	}
	fmt.Println("Result: ", totalNumScratchcards)
}
