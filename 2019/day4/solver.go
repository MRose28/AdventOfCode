package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve() (p1, p2 int) {
	p1 = getP1()
	p2 = getP2()
	return
}

func getP1() int {
	n := 0
	for i := lower; i <= upper; i++ {
		if isSixDigit(i) && isWithinBounds(i) && hasAdjacentIdenticalNumbers(i) && neverDecreasingLeftRight(i) {
			n++
			fmt.Println(i)
		}
	}
	return n
}

func getP2() int {
	n := 0
	for i := lower; i <= upper; i++ {
		if isSixDigit(i) {
			if isWithinBounds(i) {
				if hasAdjacentIdenticalNumbers(i) {
					if neverDecreasingLeftRight(i) {
						n++
					}
				} else {
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}

	return n
}

var upper = 657474
var lower = 183564

func isSixDigit(number int) bool {
	n := strconv.Itoa(number)
	for strings.HasPrefix(n, "0") {
		n = n[1:]
	}
	return len(n) == 6
}

func isWithinBounds(number int) bool {
	return number >= lower && number <= upper
}

func hasAdjacentIdenticalNumbers(number int) bool {
	s := strconv.Itoa(number)
	current := -1
	amount := 0

	for i, r := range s {
		if i == 0 {
			current, _ = strconv.Atoi(string(r))
			amount++
			continue
		}
		if n, _ := strconv.Atoi(string(r)); n == current {
			amount++
			continue
		} else {
			if amount == 2 {
				return true
			} else {
				current, _ = strconv.Atoi(string(r))
				amount = 1
			}
		}
	}

	return amount == 2
}

func neverDecreasingLeftRight(number int) bool {
	s := strconv.Itoa(number)
	for i, r := range s {
		if i == 0 {
			continue
		}
		if r < rune(s[i-1]) {
			return false
		}
	}
	return true
}

// --- Day 4: Secure Container ---
//You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.
//
//However, they do remember a few key facts about the password:
//
//It is a six-digit number.
//The value is within the range given in your puzzle input.
//Two adjacent digits are the same (like 22 in 122345).
//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
//Other than the range rule, the following are true:
//
//111111 meets these criteria (double 11, never decreases).
//223450 does not meet these criteria (decreasing pair of digits 50).
//123789 does not meet these criteria (no double).
//How many different passwords within the range given in your puzzle input meet these criteria?
//
//Your puzzle input is 183564-657474.
