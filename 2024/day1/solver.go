package day1

import (
	"slices"
	"strconv"
	"strings"

	"mrose.de/aoc/utility"
)

func Solve() (p1, p2 int) {
	i := utility.InputAsStrArr(2024, 1, false)

	p1 = part1(i)
	p2 = part2(i)

	return
}

func part1(i []string) int {
	a1, a2 := seperateArrays(i)

	slices.Sort(a1)
	slices.Sort(a2)

	sum := 0
	for i := 0; i < len(a1); i++ {
		n1 := a1[i]
		n2 := a2[i]

		if n1 > n2 {
			sum += n1 - n2
			continue
		}

		sum += n2 - n1
	}

	return sum
}

func part2(i []string) int {
	a1, a2 := seperateArrays(i)
	sum := 0

	for _, n1 := range a1 {
		count := 0
		for _, n2 := range a2 {
			if n1 == n2 {
				count++
			}
		}
		sum += n1 * count
	}

	return sum
}

func seperateArrays(i []string) ([]int, []int) {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for _, v := range i {
		parts := strings.Split(v, "   ")
		n, _ := strconv.Atoi(parts[0])
		arr1 = append(arr1, n)
		n, _ = strconv.Atoi(parts[1])
		arr2 = append(arr2, n)
	}
	return arr1, arr2
}
