package day2

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Color string

const (
	Red   Color = "red"
	Blue  Color = "blue"
	Green Color = "green"
)

type Reveals struct {
	ColorMap map[Color]int
}

type Game struct {
	ID      int
	Valid   bool
	Reveals []Reveals
}

func Solve() (part1, part2 int) {

	input := utility.StrArr(utility.Input(2023, 2))
	games := parseInput(input)

	part1 = findPossibleGames(games)
	part2 = findPowers(games)

	return
}

func findPowers(games []Game) int {
	var power int
	for _, game := range games {
		power += findPower(game)
	}

	return power
}

func findPower(game Game) int {
	// max values
	var r, g, b int
	for _, reveal := range game.Reveals {
		if reveal.ColorMap[Red] > r {
			r = reveal.ColorMap[Red]
		}
		if reveal.ColorMap[Blue] > b {
			b = reveal.ColorMap[Blue]
		}
		if reveal.ColorMap[Green] > g {
			g = reveal.ColorMap[Green]
		}
	}

	return r * g * b
}

// return added IDs of possible games
func findPossibleGames(games []Game) int {
	var res int

	for _, g := range games {
		if g.Valid {
			res += g.ID
		}
	}

	return res
}

func parseInput(input []string) []Game {
	games := make([]Game, len(input))
	for i := 0; i < len(input); i++ {
		line := input[i]
		games[i] = parseLine(line)
	}

	return games
}

func parseLine(line string) Game {
	id := parseId(strings.Split(line, ":")[0])
	bags := parseBags(strings.Split(strings.Split(line, ":")[1], ";"))

	v := validateGame(bags)

	return Game{
		ID:      id,
		Valid:   v,
		Reveals: bags,
	}

}

func validateGame(bags []Reveals) bool {
	for _, bag := range bags {
		if bag.ColorMap[Red] > 12 {
			return false
		}
		if bag.ColorMap[Green] > 13 {
			return false
		}
		if bag.ColorMap[Blue] > 14 {
			return false
		}
	}

	return true
}

func parseBags(input []string) []Reveals {
	reveals := make([]Reveals, 0)
	for _, s := range input {
		reveals = append(reveals, parseReveal(s))
	}

	return reveals
}

func parseReveal(s string) Reveals {
	colorMap := make(map[Color]int, 3)
	colorSections := strings.Split(s, ",")

	for _, section := range colorSections {
		section = strings.Trim(section, " ")
		parts := strings.Split(section, " ")
		count, _ := strconv.Atoi(parts[0])
		switch parts[1] {
		case string(Red):
			colorMap[Red] = count
		case string(Blue):
			colorMap[Blue] = count
		case string(Green):
			colorMap[Green] = count
		}
	}

	return Reveals{ColorMap: colorMap}
}

func parseId(s string) int {
	id, _ := strconv.Atoi(strings.Split(s, "Game ")[1])

	return id
}
