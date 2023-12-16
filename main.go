package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Drumstickz64/advent_of_code_2023/utils"
)

const (
	HANDTYPE_HIGHCARD uint64 = iota + 1
	HANDTYPE_ONEPAIR
	HANDTYPE_TWOPAIR
	HANDTYPE_THREE_OF_A_KIND
	HANDTYPE_FULLHOUSE
	HANDTYPE_FOUR_OF_A_KIND
	HANDTYPE_FIVE_OF_A_KIND
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
	hands := parseInput(input)

	slices.SortFunc(hands, func(handA Hand, handB Hand) int {
		return int(handA.score - handB.score)
	})

	totalWinnings := 0
	for i, hand := range hands {
		winning := hand.bid * (i + 1)
		totalWinnings += winning
	}

	fmt.Println("Result: ", totalWinnings)
}

func part2(input string) {
	panic("todo")
}

type Hand struct {
	cardsStr string
	bid      int
	score    uint64
}

func parseInput(input string) []Hand {
	hands := []Hand{}
	for _, line := range utils.Lines(input) {
		parts := strings.Split(line, " ")
		cardsStr := parts[0]
		bidStr := parts[1]
		bid, err := strconv.Atoi(strings.Trim(bidStr, " "))
		if err != nil {
			log.Fatalln("Failed to parse bid: ", err)
		}

		hands = append(hands, Hand{
			cardsStr: cardsStr,
			bid:      bid,
			score:    calculateHandScore(cardsStr),
		})
	}

	return hands
}

func calculateHandScore(cardsStr string) uint64 {
	cardPowerMap := map[string]uint64{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	matches := []int{}
	for cardName := range cardPowerMap {
		count := strings.Count(cardsStr, cardName)
		if count > 1 {
			matches = append(matches, count)
			if count >= 4 {
				break
			}
		}
	}

	var handType uint64
	switch {
	case slices.Equal(matches, []int{5}):
		handType = HANDTYPE_FIVE_OF_A_KIND
	case slices.Equal(matches, []int{4}):
		handType = HANDTYPE_FOUR_OF_A_KIND
	case slices.Equal(matches, []int{3, 2}) || slices.Equal(matches, []int{2, 3}):
		handType = HANDTYPE_FULLHOUSE
	case slices.Equal(matches, []int{3}):
		handType = HANDTYPE_THREE_OF_A_KIND
	case slices.Equal(matches, []int{2, 2}):
		handType = HANDTYPE_TWOPAIR
	case slices.Equal(matches, []int{2}):
		handType = HANDTYPE_ONEPAIR
	default:
		handType = HANDTYPE_HIGHCARD
	}

	var score uint64 = handType * 10000000000
	for i := 0; i < 5; i++ {
		card := string(cardsStr[i])
		power := cardPowerMap[card]
		score += power * uint64(math.Pow10((4-i)*2))
	}

	return score
}
