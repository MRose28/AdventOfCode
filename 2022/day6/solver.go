package day6

import (
	"mrose.de/aoc/utility"
	"strings"
)

func Solve() (result int) {
	input := utility.Input(2022, 6)
	result = findMarker2(input)

	return
}

// star 1
func findMarker(input string) int {
	lastFour := ""

	for i, r := range input {
		if !strings.Contains(lastFour, string(r)) {
			lastFour += string(r)
		} else {
			lastFour = ""
		}
		if len(lastFour) == 4 {
			return i
		}
	}
	return -1
}

// star 2
func findMarker2(input string) int {
	lastFourteen := ""

	for i, r := range input {
		if !strings.Contains(lastFourteen, string(r)) {
			lastFourteen += string(r)
		} else {
			idx := strings.Index(lastFourteen, string(r))
			lastFourteen = lastFourteen[idx+1:]
			lastFourteen += string(r)
		}
		if len(lastFourteen) == 14 {
			return i + 1
		}
	}
	return -1
}
