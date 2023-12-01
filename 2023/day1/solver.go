package day1

import (
	"mrose.de/aoc/utility"
	"regexp"
	"strconv"
	"strings"
)

func Solve() (part1, part2 int) {
	i := utility.StrArr(utility.Input(2023, 1))

	part1 = findNumbers(i)
	// no extra solution for part 1
	part2 = 0
	return
}

func findNumbers(input []string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	found := make([]string, len(input))
	for i, n := range input {
		var temp string
		for j := 0; j < len(n); j++ {
			ok, _ := regexp.MatchString("[0-9]+", string(n[j]))
			if !ok {
				temp += string(n[j])
				for idx, s := range numbers {
					if strings.Contains(temp, s) {
						found[i] += strconv.Itoa(idx + 1)
						temp = ""
						j--
						break
					}
				}
				continue
			}

			found[i] += string(n[j])
			temp = ""

		}
	}

	var res int
	for _, n := range found {
		s := ""
		s += string(n[0])
		s += string(n[len(n)-1])
		num, _ := strconv.Atoi(s)
		res += num
	}

	return res
}
