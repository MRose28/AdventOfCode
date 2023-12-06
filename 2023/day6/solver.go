package day6

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Race struct {
	AllowedTime, Distance int
}

func Solve() (part1, part2 int) {

	input := utility.StrArr(utility.Input(2023, 6))

	part1 = P1(input)
	part2 = P2(input)

	return
}

func P2(input []string) int {
	race := parseRace(input)

	return waysToWin(race)
}

func parseRace(input []string) Race {
	t := parseInputP2(input[0])
	d := parseInputP2(input[1])

	return Race{
		AllowedTime: t,
		Distance:    d,
	}
}

func parseInputP2(s string) int {
	s = strings.ReplaceAll(strings.Split(s, ":")[1], " ", "")
	n, _ := strconv.Atoi(s)

	return n
}

func P1(input []string) int {
	races := parseRaces(input)
	ways := numOfWaysToWin(races)
	product := 0

	for _, w := range ways {
		if product == 0 {
			product = w
			continue
		}
		product *= w
	}

	return product
}

func numOfWaysToWin(races []Race) []int {
	ways := make([]int, 0)

	for _, race := range races {
		counterOfWays := waysToWin(race)
		ways = append(ways, counterOfWays)
	}

	return ways
}

func waysToWin(race Race) int {
	counterOfWays := 0
	var wasFaster bool
	for speed := 0; speed < race.AllowedTime; speed++ {
		if speed*(race.AllowedTime-speed) > race.Distance {
			if !wasFaster {
				wasFaster = true
			}
			counterOfWays++
		} else {
			if wasFaster {
				break
			}
		}

	}

	return counterOfWays
}

func parseRaces(input []string) []Race {
	races := make([]Race, 0)
	times := parseInput(input[0])
	distances := parseInput(input[1])

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			AllowedTime: times[i],
			Distance:    distances[i],
		})
	}

	return races
}

func parseInput(s string) []int {
	res := make([]int, 0)
	numbers := strings.Split(strings.Trim(strings.Split(s, ":")[1], " "), " ")

	for _, sNum := range numbers {
		n, _ := strconv.Atoi(sNum)
		if n > 0 {
			res = append(res, n)
		}
	}

	return res
}
