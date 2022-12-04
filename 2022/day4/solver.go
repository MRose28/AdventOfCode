package day4

import (
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type elf struct {
	start, end int
}

type group []elf

func Solve() (result int) {
	input := utility.StrArr(utility.Input(2022, 4))
	groups := groupUp(input)
	result = overlapping(groups)
	return
}

// count all overlapping groups / star 2
func overlapping(groups []group) (score int) {
	for _, g := range groups {
		containing := false
		if g[1].start <= g[0].end && g[1].end >= g[0].end {
			containing = true
		}
		if g[0].start <= g[1].end && g[0].end >= g[1].end {
			containing = true
		}

		if containing {
			score++
		}
	}
	return
}

// count all occurrences of one group fully containing another / star1
func countContainingGroups(groups []group) (score int) {
	for _, g := range groups {
		containing := false
		if g[1].start >= g[0].start && g[1].end <= g[0].end {
			containing = true
		}
		if g[0].start >= g[1].start && g[0].end <= g[1].end {
			containing = true
		}
		if containing {
			score++
		}
	}

	return
}

// list groups
func groupUp(input []string) []group {
	groups := make([]group, 0)

	for _, pairRaw := range input {
		elves := make([]elf, 0)
		elfRaw := strings.Split(pairRaw, ",")
		for i := 0; i < 2; i++ {
			start, err := strconv.Atoi(strings.Split(elfRaw[i], "-")[0])
			if err != nil {
				log.Fatalf("could not parse start date %v\nerr: %e\n", strings.Split(elfRaw[i], "-")[0], err)
			}
			end, err := strconv.Atoi(strings.Split(elfRaw[i], "-")[1])
			if err != nil {
				log.Fatalf("could not parse start date %v\nerr: %e\n", strings.Split(elfRaw[i], "-")[1], err)
			}
			elves = append(elves, elf{
				start: start,
				end:   end,
			})
		}
		if len(elves) > 0 {
			groups = append(groups, elves)
		}
	}
	return groups
}
