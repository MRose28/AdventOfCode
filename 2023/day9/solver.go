package day9

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type TimeSeries struct {
	original    []int
	differences [][]int
}

type Direction string

const (
	Forwards  Direction = "forward"
	Backwards Direction = "backwards"
)

func (ts TimeSeries) extrapolate(direction Direction) TimeSeries {
	if direction == Forwards {
		return ts.forwards()
	} else if direction == Backwards {
		return ts.backwards()
	}

	return ts
}

func (ts TimeSeries) forwards() TimeSeries {
	for i := len(ts.differences) - 2; i >= 0; i-- {
		current := ts.differences[i]
		current = append(current, current[len(current)-1]+ts.differences[i+1][len(ts.differences[i+1])-1])
		ts.differences[i] = current
	}

	ts.original = append(ts.original, ts.original[len(ts.original)-1]+ts.differences[0][len(ts.differences[0])-1])

	return ts
}

func (ts TimeSeries) backwards() TimeSeries {
	for i := len(ts.differences) - 2; i >= 0; i-- {
		current := ts.differences[i]
		current = append([]int{current[0] - ts.differences[i+1][0]}, current...)
		ts.differences[i] = current
	}

	ts.original = append([]int{ts.original[0] - ts.differences[0][0]}, ts.original...)

	return ts
}

func differences(arr []int) []int {
	d := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		d = append(d, arr[i]-arr[i-1])
	}

	return d
}

func allDifferences(ts []int) [][]int {
	allDiffs := make([][]int, 0)
	current := ts
	for {
		currentDiffs := differences(current)
		allDiffs = append(allDiffs, currentDiffs)
		if allZero(currentDiffs) {
			break
		} else {
			current = currentDiffs
		}
	}

	return allDiffs
}

func allZero(arr []int) bool {
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}

func Solve() (p1, p2 int) {
	input := utility.InputAsStrArr(2023, 9, false)
	ts := parseTimeSeries(input)
	p1 = solveBoth(ts, Forwards)
	p2 = solveBoth(ts, Backwards)
	return
}

func solveBoth(ts []TimeSeries, direction Direction) int {
	for i, t := range ts {
		ts[i] = t.extrapolate(direction)
	}
	var sum int
	if direction == Forwards {
		for _, t := range ts {
			sum += t.original[len(t.original)-1]
		}
	} else if direction == Backwards {
		for _, t := range ts {
			sum += t.original[0]
		}
	}

	return sum
}

func parseTimeSeries(input []string) []TimeSeries {
	res := make([]TimeSeries, len(input))
	for i, line := range input {
		original := make([]int, 0)

		for _, s := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(s)
			original = append(original, n)
		}
		res[i] = TimeSeries{
			original:    original,
			differences: allDifferences(original),
		}
	}

	return res
}

/*
--- Day 9: Mirage Maintenance ---
You ride the camel through the sandstorm and stop where the ghost's maps told you to stop. The sandstorm subsequently subsides, somehow seeing you standing at an oasis!

The camel goes to get some water and you stretch your neck. As you look up, you discover what must be yet another giant floating island, this one made of metal! That must be where the parts to fix the sand machines come from.

There's even a hang glider partially buried in the sand here; once the sun rises and heats up the sand, you might be able to use the glider and the hot air to get all the way up to the metal island!

While you wait for the sun to rise, you admire the oasis hidden here in the middle of Desert Island. It must have a delicate ecosystem; you might as well take some ecological readings while you wait. Maybe you can report any environmental instabilities you find to someone so the oasis can be around for the next sandstorm-worn traveler.

You pull out your handy Oasis And Sand Instability Sensor and analyze your surroundings. The OASIS produces a report of many values and how they are changing over time (your puzzle input). Each line in the report contains the history of a single value. For example:

0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
To best protect the oasis, your environmental report should include a prediction of the next value in each history. To do this, start by making a new sequence from the difference at each step of your history. If that sequence is not all zeroes, repeat this process, using the sequence you just generated as the input sequence. Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.

In the above dataset, the first history is 0 3 6 9 12 15. Because the values increase by 3 each step, the first sequence of differences that you generate will be 3 3 3 3 3. Note that this sequence has one fewer value than the input sequence because at each step it considers two numbers from the input. Since these values aren't all zero, repeat the process: the values differ by 0 at each step, so the next sequence is 0 0 0 0. This means you have enough information to extrapolate the history! Visually, these sequences can be arranged like this:

0   3   6   9  12  15
  3   3   3   3   3
    0   0   0   0
To extrapolate, start by adding a new zero to the end of your list of zeroes; because the zeroes represent differences between the two values above them, this also means there is now a placeholder in every sequence above it:

0   3   6   9  12  15   B
  3   3   3   3   3   A
    0   0   0   0   0
You can then start filling in placeholders from the bottom up. A needs to be the result of increasing 3 (the value to its left) by 0 (the value below it); this means A must be 3:

0   3   6   9  12  15   B
  3   3   3   3   3   3
    0   0   0   0   0
Finally, you can fill in B, which needs to be the result of increasing 15 (the value to its left) by 3 (the value below it), or 18:

0   3   6   9  12  15  18
  3   3   3   3   3   3
    0   0   0   0   0
So, the next value of the first history is 18.

Finding all-zero differences for the second history requires an additional sequence:

1   3   6  10  15  21
  2   3   4   5   6
    1   1   1   1
      0   0   0
Then, following the same process as before, work out the next value in each sequence from the bottom up:

1   3   6  10  15  21  28
  2   3   4   5   6   7
    1   1   1   1   1
      0   0   0   0
So, the next value of the second history is 28.

The third history requires even more sequences, but its next value can be found the same way:

10  13  16  21  30  45  68
   3   3   5   9  15  23
     0   2   4   6   8
       2   2   2   2
         0   0   0
So, the next value of the third history is 68.

If you find the next value for each history in this example and add them together, you get 114.

Analyze your OASIS report and extrapolate the next value for each history. What is the sum of these extrapolated values?
*/
