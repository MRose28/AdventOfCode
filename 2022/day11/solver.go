package day11

import (
	"fmt"
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

var monkeyBusiness uint64

var lcm uint64 = 1

func Solve() (part1, part2 int) {

	input := utility.Input(2022, 11)
	monkeys := createMonkeys(input)
	getLCM(monkeys)
	play(monkeys, 10000)
	findMostActive(monkeys)

	return
}

func getLCM(monkeys []*monkey) {

	for _, m := range monkeys {
		lcm *= uint64(m.divisibleBy)
	}
}
func findMostActive(monkeys []*monkey) {
	var first, second int
	for _, m := range monkeys {
		if m.inspectionCount > first {
			second = first
			first = m.inspectionCount
			continue
		}
		if m.inspectionCount > second {
			second = m.inspectionCount
		}
	}

	monkeyBusiness = uint64(first * second)
	fmt.Println(monkeyBusiness)
}

func play(monkeys []*monkey, rounds int) {
	check := 3000
	for round := 0; round < rounds; round++ {
		if round == check {
			fmt.Printf("Round %v: \n", round+1)
		}
		for index, m := range monkeys {
			for {
				if len(m.items) == 0 {
					break
				}
				i := m.items[0]
				i.worryLevel = m.operation(i.worryLevel)
				i.worryLevel %= lcm
				m.inspectionCount++
				// part1 only
				//i.decreaseWorryLevel()
				m.throw(m.test(i))
			}
			if round == check {
				fmt.Printf("Monkey %d: inspected items %d times\n", index, m.inspectionCount)
			}
		}
		if round == check {
			fmt.Println()
		}
	}
}
func setTestResults(monkeys []*monkey) {
	for _, m := range monkeys {
		m.pos = monkeys[m.posId]
		m.neg = monkeys[m.negId]
	}
}

func createMultFunc(s string, s2 string, currentMonkey *monkey) func(old uint64) uint64 {
	return func(old uint64) (result uint64) {
		n1, err := strconv.Atoi(s)
		var b1, b2 uint64
		if err != nil {
			b1 = old
		} else {
			b1 = uint64(n1)
		}
		n2, err := strconv.Atoi(s2)
		if err != nil {
			b2 = old
		} else {
			b2 = uint64(n2)
		}
		result = b1 * b2

		return
	}
}

func createAddFunc(s string, s2 string, currentMonkey *monkey) func(old uint64) uint64 {

	return func(old uint64) (result uint64) {
		n1, err := strconv.Atoi(s)
		var b1, b2 uint64
		if err != nil {
			b1 = old
		} else {
			b1 = uint64(n1)
		}
		n2, err := strconv.Atoi(s2)
		if err != nil {
			b2 = old
		} else {
			b2 = uint64(n2)
		}
		result = b1 + b2

		return
	}
}

func createMonkeys(input string) (monkeys []*monkey) {
	strArr := utility.StrArr(input)
	monkeys = make([]*monkey, 0)
	currentMonkey := newMonkey()
	for _, s := range strArr {
		s = strings.Trim(s, " ")
		arr := strings.Split(s, ":")
		start := arr[0]
		if strings.Contains(start, "Monkey") {
			currentMonkey = newMonkey()
			continue
		}
		if strings.Contains(start, "Starting") {
			worryLevels := strings.Split(strings.Trim(arr[1], " "), ", ")
			items := make([]*item, 0)
			for _, level := range worryLevels {
				level, err := strconv.Atoi(level)
				if err != nil {
					log.Fatalf("could not parse level: %v", level)
				}
				item := &item{worryLevel: uint64(level)}
				items = append(items, item)
			}
			currentMonkey.addItems(items)
			continue
		}

		if strings.Contains(arr[0], "Operation") {
			comps := strings.Split(strings.Replace(arr[1], " new = ", "", 1), " ")
			operator := strings.Trim(comps[1], " ")
			var operation func(old uint64) uint64
			switch operator {
			case "+":
				operation = createAddFunc(comps[0], comps[2], currentMonkey)
			case "*":
				operation = createMultFunc(comps[0], comps[2], currentMonkey)
			}
			currentMonkey.operation = operation
			continue
		}

		if strings.Contains(arr[0], "Test") {
			nr, err := strconv.Atoi(strings.Split(strings.Trim(arr[1], " "), " ")[2])
			if err != nil {
				log.Fatal("could not parse test")
			}
			currentMonkey.divisibleBy = nr
			continue
		}

		if strings.Contains(arr[0], "true") {
			id, err := strconv.Atoi(strings.Split(strings.Trim(arr[1], " "), " ")[3])
			if err != nil {
				log.Fatal("could not parse monkey id")
			}
			currentMonkey.posId = id
			continue
		}
		if strings.Contains(arr[0], "false") {
			id, err := strconv.Atoi(strings.Split(strings.Trim(arr[1], " "), " ")[3])
			if err != nil {
				log.Fatal("could not parse monkey id")
			}
			currentMonkey.negId = id
			monkeys = append(monkeys, currentMonkey)
			continue
		}
	}
	setTestResults(monkeys)

	return
}
