// RULES
// '()' => take precedence
// Evaluation from left to right. '+' & '*' take the same precedence.

package day18

import (
	"fmt"
	"strings"

	"mrose.de/aoc/utility"
)

func Solve() (p1, p2 int) {
	input := utility.InputAsStrArr(2020, 18, true)

	for _, v := range input {
		p1 += calc(v)
	}

	return
}

func calc(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)

	return 2
}
