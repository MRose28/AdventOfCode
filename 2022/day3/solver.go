package day3

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strings"
)

func Solve() (result int) {
	rucksacks := utility.StrArr(utility.Input(2022, 3))
	result = prioritySum2(rucksacks)
	return
}

// star1
func prioritySum(rucksacks []string) (score int) {

	for _, r := range rucksacks {
		c1 := make([]rune, 0)
		c2 := make([]rune, 0)
		for i, t := range r {
			if i < len(r)/2 {
				c1 = append(c1, t)
			} else {
				c2 = append(c2, t)
			}
		}
		found := false
		for _, r := range c1 {
			for _, r2 := range c2 {
				newScore := scoreRunes(r, r2)
				if newScore > 0 {
					score += newScore
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	return
}

// star2
func prioritySum2(rucksacks []string) (score int) {

	for i, _ := range rucksacks {
		if (i+1)%3 == 0 {
			found := false
			for _, r1 := range rucksacks[i-2] {
				for _, r2 := range rucksacks[i-1] {
					for _, r3 := range rucksacks[i] {
						newScore := scoreRunes2(r1, r2, r3)
						if newScore > 0 {
							score += newScore
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	return
}

func scoreRunes(c1 rune, c2 rune) (score int) {
	alphabet := "-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if c1 == c2 {
		s := string(c1)
		score = strings.IndexAny(alphabet, s)
		fmt.Printf("String: %v - %d\n", s, score)
		return
	}
	return 0
}

func scoreRunes2(c1 rune, c2 rune, c3 rune) (score int) {
	alphabet := "-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if c1 == c2 && c2 == c3 {
		s := string(c1)
		score = strings.IndexAny(alphabet, s)
		fmt.Printf("String: %v - %d\n", s, score)
		return
	}
	return 0
}
