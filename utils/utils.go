package utils

import "strings"

func CalculateLineLength(input string) int {
	for i := 0; true; i++ {
		if input[i:i+2] == "\r\n" {
			return i + 2
		}
	}

	panic("Reached unreachable code")
}

// calculate the indeces of characters adjacent to number, not in order
func CalculateAdjacentSpaces(numberIndex []int, lineLength int) []int {
	spaces := []int{}

	// top and bottom
	for i := numberIndex[0]; i < numberIndex[1]; i++ {
		spaces = append(spaces, i-lineLength, i+lineLength)
	}

	// corners
	spaces = append(spaces,
		numberIndex[0]-1+lineLength,
		numberIndex[0]-1,
		numberIndex[0]-1-lineLength,
		numberIndex[1]+lineLength,
		numberIndex[1],
		numberIndex[1]-lineLength,
	)

	return spaces
}

func ByteIsDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func Lines(s string) []string {
	return strings.Split(s, "\r\n")
}
