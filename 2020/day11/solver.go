package day11

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strings"
)

func Solve() (result int) {
	arr := utility.StrArr(utility.Input2020Day11())
	runes := stringLists(arr)
	changed := true
	result, changed, runes = takenSeatsAndUpdate(runes)
	for changed {
		result, changed, runes = takenSeatsAndUpdate(runes)
	}
	return
}

func printMap(runes map[int][]string) {
	for i := 0; i < len(runes); i++ {
		line := runes[i]
		for _, cRune := range line {
			fmt.Print(cRune)
		}
		fmt.Print("\n")
	}
}

func stringLists(arr []string) (runes map[int][]string) {
	runes = make(map[int][]string, 0)
	for pos, line := range arr {
		runes[pos] = strings.SplitN(line, "", len(line))
	}
	return
}

func takenSeatsAndUpdate(runes map[int][]string) (result int, changed bool, stringMap map[int][]string) {
	originalRunes := make(map[int][]string, 0)
	for i := 0; i < len(runes); i++ {
		tempLine := make([]string, 0)
		for _, v := range runes[i] {
			tempLine = append(tempLine, v)
		}
		originalRunes[i] = tempLine
	}
	for linePos := 0; linePos < len(runes); linePos++ {
		line := runes[linePos]
		for runePos := range line {
			taken := 0
			if originalRunes[linePos][runePos] != "." {

				//Star1 solution
				//if linePos == 0 {
				//	taken += TakenSeatsDiagonals(runePos, originalRunes[linePos+1], false)
				//} else if linePos == len(originalRunes)-1 {
				//	taken += TakenSeatsDiagonals(runePos, originalRunes[linePos-1], false)
				//} else {
				//	taken += TakenSeatsDiagonals(runePos, originalRunes[linePos+1], false)
				//	taken += TakenSeatsDiagonals(runePos, originalRunes[linePos-1], false)
				//}
				//taken += TakenSeatsDiagonals(runePos, originalRunes[linePos], true)

				taken += TakenSeatsDiagonals(originalRunes, linePos, runePos)

				if line[runePos] == "#" {
					if taken >= 5 {
						line[runePos] = "L"
						changed = true
					} else {
						result++
					}
				} else if line[runePos] == "L" {
					if taken == 0 {
						line[runePos] = "#"
						changed = true
						result++
					}
				}
			}
		}
		runes[linePos] = line
	}

	stringMap = runes
	return
}
