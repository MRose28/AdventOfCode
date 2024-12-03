package day1

import (
	"regexp"
	"strconv"
	"strings"

	"mrose.de/aoc/utility"
)

type Operator int

const (
	Mul Operator = iota
)

type Expression struct {
	N1, N2   int
	Operator Operator
	Pos      int
}

func (e Expression) solve() int {
	return e.N1 * e.N2
}

func NewExpression(n1, n2 int, op Operator, pos int) Expression {
	return Expression{
		n1, n2, op, pos,
	}
}

// p2 solved in part1() function
func Solve() (p1, p2 int) {
	instructions := utility.InputAsStrArr(2024, 3, false)
	instructionString := ""

	for _, instruction := range instructions {
		instructionString += instruction
	}
	p1 = part1(instructionString)
	p2 = part2(instructionString)

	return
}

func part1(instructions string) int {
	expressions := findExpressions(instructions)
	var result int

	for _, e := range expressions {
		result += e.solve()
	}

	return result
}

func findExpressions(instructions string) []Expression {
	// The regex pattern to match
	pattern := `mul\(\d{1,3},\d{1,3}\)`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(instructions, -1)
	positions := re.FindAllStringIndex(instructions, -1)

	expressions := make([]Expression, 0)

	dodontlist := getdodontlist(instructions)

	for i, m := range matches {
		n1, _ := strconv.Atoi(strings.Split(strings.Split(m, "(")[1], ",")[0])
		n2, _ := strconv.Atoi(strings.Split(strings.Split(m, ",")[1], ")")[0])
		op := Mul

		expressions = append(expressions, NewExpression(n1, n2, op, positions[i][0]))
	}

	return executableExpressions(expressions, dodontlist)
}

func executableExpressions(expressions []Expression, dodontlist []int) []Expression {
	var lastPos int
	enabled := true
	var result = []Expression{}
	dodontlist = append(dodontlist, 9999999999)

	for _, nextChange := range dodontlist {
		for _, exp := range expressions {
			if exp.Pos >= nextChange {
				enabled = !enabled
				lastPos = nextChange
				break
			} else {
				if enabled && exp.Pos > lastPos {
					result = append(result, exp)
				}
			}
		}
	}

	return result
}

func getdodontlist(instructions string) []int {
	l := []int{}

	pattern := `do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(instructions, -1)
	positions := re.FindAllStringIndex(instructions, -1)
	enabled := true

	for i := 0; i < len(matches); i++ {
		if enabled {
			if matches[i] == "don't()" {
				enabled = false
				l = append(l, positions[i][0])
			}
		} else {
			if matches[i] == "do()" {
				enabled = true
				l = append(l, positions[i][0])
			}
		}
	}

	return l
}

func part2(instructions string) int {

	return 666
}
