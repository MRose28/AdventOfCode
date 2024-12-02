package day1

import (
	"strconv"
	"strings"

	"mrose.de/aoc/utility"
)

func Solve() (p1, p2 int) {
	reports := createReports(utility.InputAsStrArr(2024, 2, false))
	p1 = part1(reports)
	p2 = part2(reports)

	return
}

func part1(reports [][]int) int {
	safe := 0
	for _, report := range reports {
		if x, _ := calc(report); x {
			safe++
		}
	}

	return safe
}

func part2(reports [][]int) int {
	safe := 0
	for _, report := range reports {
		if x, i := calc(report); x {
			safe++
		} else {
			arr := make([]int, 0)
			arr = append(arr, report[:i]...)
			arr = append(arr, report[i+1:]...)
			if y, _ := calc(arr); y {
				safe++
			} else {
				arr := make([]int, 0)
				arr = append(arr, report[:i-1]...)
				arr = append(arr, report[i:]...)
				if y, _ := calc(arr); y {
					safe++
				}
			}
		}
	}

	return safe
}

func calc(report []int) (bool, int) {
	inc := true
	if report[0] > report[len(report)-1] {
		inc = false
	}

	if inc {
		previous := report[0] - 1
		for i, n := range report {
			if n <= previous || n-previous > 3 {
				return false, i
			}
			previous = n
		}
	} else {
		previous := report[0] + 1
		for k, n := range report {
			if n >= previous || previous-n > 3 {
				return false, k
			}
			previous = n
		}
	}

	return true, -1
}

func createReports(s []string) [][]int {
	reports := make([][]int, 0)
	for _, v := range s {
		currentReport := make([]int, 0)
		for _, level := range strings.Split(v, " ") {
			n, _ := strconv.Atoi(level)
			currentReport = append(currentReport, n)
		}
		reports = append(reports, currentReport)
	}

	return reports
}
