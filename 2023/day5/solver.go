package day5

import (
	"math"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type SeedNumbers []int
type SeedRange struct {
	Start, Range int
}

type LocationNum struct {
	Value              int
	PossibleSeedValues []int
}
type ConversionRule struct {
	Src, Dest int
	Range     int
}
type ConversionMap struct {
	Name  string
	Rules []ConversionRule
}

func Solve() (part1, part2 int) {
	inputBlocks := strings.Split(utility.Input(2023, 5), "\n\n")
	seedNumbers := parseSeedNumbers(inputBlocks[0])
	maps := parseMaps(inputBlocks[1:])

	part1 = getLowestLocationNumber(seedNumbers, maps)
	ranges := parseSeedRanges(seedNumbers)
	locNumbers := locNums(maps)
	part2 = findLowestPossibleLocNumber(locNumbers, ranges, seedNumbers)
	return
}

func findLowestPossibleLocNumber(numbers []LocationNum, ranges []SeedRange, seedNumbers SeedNumbers) int {
	low := math.MaxInt
	for _, n := range numbers {
		for _, r := range ranges {
			for _, value := range n.PossibleSeedValues {
				if value >= r.Start && value < r.Start+r.Range {
					return n.Value
				}
			}
		}
	}
	return low
}

func locNums(maps []ConversionMap) []LocationNum {
	locNumbers := addPossibleSeedValues(maps)
	return locNumbers
}

var history = make(map[string]map[int]int)

func addPossibleSeedValues(maps []ConversionMap) []LocationNum {
	lNums := make([]LocationNum, 0)
	for _, m := range maps {
		for _, rule := range m.Rules {
			for n := rule.Dest; n < rule.Dest+rule.Range; n++ {
				values := findPossibleSeedValues(n, maps)

				lNums = append(lNums, LocationNum{
					Value:              n,
					PossibleSeedValues: values,
				})
			}
		}
	}

	return lNums
}

func findPossibleSeedValues(n int, maps []ConversionMap) []int {
	current := n
	res := make([]int, 0)
	for i := len(maps) - 1; i >= 0; i-- {
		m := maps[i]
		if history[m.Name][n] != 0 {
			current = history[m.Name][n]
			continue
		}
		if history[m.Name] == nil {
			history[m.Name] = make(map[int]int)
		}
		for k, rule := range m.Rules {
			if current >= rule.Dest && current < rule.Dest+rule.Range {
				current = rule.Src + (current - rule.Dest)
				history[m.Name][n] = current
				break
			}
			if k == len(m.Rules) {
				history[m.Name][n] = current
				current = n
			}
		}
	}
	res = append(res, n)
	return res
}

func parseSeedRanges(numbers SeedNumbers) []SeedRange {
	ranges := make([]SeedRange, 0)
	for i := 0; i < len(numbers); i += 2 {
		ranges = append(ranges, SeedRange{
			Start: numbers[i],
			Range: numbers[i+1],
		})
	}

	return ranges
}

func getLowestLocationNumber(numbers SeedNumbers, maps []ConversionMap) int {
	low := math.MaxInt
	for _, n := range numbers {
		current := n
		for _, m := range maps {
			for _, rule := range m.Rules {
				var changed bool
				current, changed = convertedNumber(current, rule)
				if changed {
					break
				}
			}
		}
		if current < low {
			low = current
		}
	}

	return low
}

func convertedNumber(n int, m ConversionRule) (int, bool) {

	if n >= m.Src && n <= m.Src+m.Range-1 {
		return m.Dest + (n - m.Src), true
	}
	return n, false
}

func parseMaps(blocks []string) []ConversionMap {

	maps := make([]ConversionMap, 0)

	for _, block := range blocks {
		name := strings.Split(block, ":\n")[0]
		rules := parseRules(utility.StrArr(strings.Split(block, ":\n")[1]))

		maps = append(maps, ConversionMap{
			Name:  name,
			Rules: rules,
		})
	}

	return maps
}

func parseRules(block []string) []ConversionRule {
	cRules := make([]ConversionRule, 0)

	for _, line := range block {
		numsAsStr := strings.Split(line, " ")
		nums := make([]int, 0)
		for _, s := range numsAsStr {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		cRules = append(cRules, ConversionRule{
			Src:   nums[1],
			Dest:  nums[0],
			Range: nums[2],
		})
	}

	return cRules
}

func parseSeedNumbers(seedBlock string) SeedNumbers {
	nums := make(SeedNumbers, 0)

	s := strings.Split(strings.Trim(strings.Split(seedBlock, ": ")[1], " "), " ")

	for _, str := range s {
		n, _ := strconv.Atoi(str)
		nums = append(nums, n)
	}

	return nums
}
