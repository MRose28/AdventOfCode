package day2

import (
	"mrose.de/aoc/utility"
	"strings"
)

func Solve() (result int) {
	input := utility.StrArrCustom(utility.Input(2022, 2), "\n")

	for _, s := range input {
		p := strings.Split(s, " ")[1]
		o := strings.Split(s, " ")[0]
		result += rps2(p, o)
	}

	return
}

// day1
func rps(p string, o string) (score int) {
	switch p {
	case "X":
		score += 1
		switch o {
		case "A":
			score += 3
		case "C":
			score += 6
		}
	case "Y":
		score += 2
		switch o {
		case "A":
			score += 6
		case "B":
			score += 3
		}
	case "Z":
		score += 3
		switch o {
		case "B":
			score += 6
		case "C":
			score += 3
		}

	}
	return
}

// X = loss		y = draw	z = win
// day2
func rps2(p string, o string) (score int) {
	switch p {
	case "X":
		switch o {
		case "A":
			score += 3
		case "B":
			score += 1
		case "C":
			score += 2
		}
	case "Y":
		score += 3
		switch o {
		case "A":
			score += 1
		case "B":
			score += 2
		case "C":
			score += 3
		}
	case "Z":
		score += 6
		switch o {
		case "A":
			score += 2
		case "B":
			score += 3
		case "C":
			score += 1
		}

	}
	return
}
