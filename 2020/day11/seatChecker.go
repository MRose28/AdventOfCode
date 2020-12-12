package day11

func TakenSeatsInLine(runePos int, line []string, isCurrent bool) (taken int) {
	taken = 0
	positions := position(runePos, len(line), isCurrent)
	for _, pos := range positions {
		if line[pos] == "#" {
			taken++
		}
	}

	return
}

func TakenSeatsDiagonals(lineMap map[int][]string, linePos int, runePos int) (result int) {
	result = 0
	hSize := len(lineMap[0])
	vSize := len(lineMap)

	result += horizontal(lineMap, linePos, runePos, hSize)
	result += vertical(lineMap, linePos, runePos, vSize)
	result += diagonal(lineMap, linePos, runePos, hSize, vSize)
	return
}

func vertical(letters map[int][]string, linePos int, runePos int, size int) (result int) {
	if linePos < size-1 {
		for i := linePos + 1; i < size; i++ {
			currentRune := letters[i][runePos]
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				break
			}
		}
	}
	if linePos > 0 {
		for i := linePos - 1; i >= 0; i-- {
			currentRune := letters[i][runePos]
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				break
			}
		}
	}
	return
}

func horizontal(letters map[int][]string, linePos int, runePos int, size int) (result int) {
	//right
	if runePos < size-1 {
		for i := runePos + 1; i < size; i++ {
			currentRune := letters[linePos][i]
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				break
			}
		}
	}

	//left
	if runePos > 0 {
		for i := runePos - 1; i >= 0; i-- {
			currentRune := letters[linePos][i]
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				break
			}
		}
	}
	return
}

func diagonal(letters map[int][]string, linePos int, runePos int, hSize int, vSize int) (result int) {
	if linePos > 0 && runePos > 0 {
		ok := true
		counter := 1
		for ok {
			currentRune := letters[linePos-counter][runePos-counter]
			counter++
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				ok = false
			}
			if linePos-counter < 0 || runePos-counter < 0 {
				ok = false
			}
		}
	}
	if linePos < vSize-1 && runePos < hSize-1 {
		ok := true
		counter := 1
		for ok {
			currentRune := letters[linePos+counter][runePos+counter]
			counter++
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				ok = false
			}
			if linePos+counter >= vSize || runePos+counter >= hSize {
				ok = false
			}
		}
	}
	if linePos > 0 && runePos < hSize-1 {
		ok := true
		counter := 1
		for ok {
			currentRune := letters[linePos-counter][runePos+counter]
			counter++
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				ok = false
			}
			if linePos-counter < 0 || runePos+counter >= hSize {
				ok = false
			}
		}
	}
	if linePos < vSize-1 && runePos > 0 {
		ok := true
		counter := 1
		for ok {
			currentRune := letters[linePos+counter][runePos-counter]
			counter++
			if currentRune != "." {
				if currentRune == "#" {
					result++
				}
				ok = false
			}
			if linePos+counter >= vSize || runePos-counter < 0 {
				ok = false
			}
		}
	}
	return
}

func position(runePos int, runeArrLen int, isCurrent bool) (positionsToCheck []int) {
	positionsToCheck = make([]int, 0)
	if runePos == 0 {
		positionsToCheck = append(positionsToCheck, runePos+1)
	} else if runePos == runeArrLen-1 {
		positionsToCheck = append(positionsToCheck, runePos-1)
	} else {
		positionsToCheck = append(positionsToCheck, runePos+1)
		positionsToCheck = append(positionsToCheck, runePos-1)
	}
	if !isCurrent {
		positionsToCheck = append(positionsToCheck, runePos)
	}
	return
}
