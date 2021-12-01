// RULES
// '()' => take precedence
// Evaluation from left to right. '+' & '*' take the same precedence.

package day18

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strings"
)

func Solve() (result int) {
	input := utility.Input2020Day18()
	input = strings.ReplaceAll(input, " ", "")
	lines := make([]InputLine, 0)
	for _, v := range utility.StrArr(input) {
		lines = append(lines, InputLine{
			input: v,
		})
	}

	for i, v := range lines {
		print(fmt.Sprintf("Line %v:", i),v.getResult(), "\n")
	}

	return
}
