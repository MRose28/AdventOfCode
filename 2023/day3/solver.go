package day3

import (
	"mrose.de/aoc/utility"
	"regexp"
	"strconv"
)

type Num struct {
	Value        int
	fxPos, lxPos int
	y            int
}

type Gear struct {
	Adjacent []Num
}

func (g Gear) gearRatio() int {
	if len(g.Adjacent) == 2 {
		return g.Adjacent[0].Value * g.Adjacent[1].Value
	}
	return 0
}

func (n Num) isAdjacent(symbols []Symbol) bool {
	for _, symbol := range symbols {
		if symbol.Y > n.y+1 || symbol.Y < n.y-1 {
			continue
		}

		if symbol.X >= n.fxPos-1 && symbol.X <= n.lxPos+1 {
			return true
		}
	}
	return false
}

type Symbol struct {
	Value rune
	X, Y  int
}

func Solve() (part1, part2 int) {
	input := utility.StrArr(utility.Input(2023, 3))

	symbols, nums := parse(input)

	part1 = Part1(symbols, nums)
	part2 = Part2(symbols, nums)

	return
}

func Part2(symbols []Symbol, nums []Num) int {
	sum := 0
	gears := getGears(symbols, nums)

	for _, gear := range gears {
		sum += gear.gearRatio()
	}

	return sum
}

func getGears(symbols []Symbol, nums []Num) []Gear {
	gears := make([]Gear, 0)
	for _, symbol := range symbols {
		if symbol.Value == '*' {
			adjacent := make([]Num, 0)
			for _, num := range nums {
				if symbol.Y < num.y-1 || symbol.Y > num.y+1 {
					continue
				}
				if symbol.X >= num.fxPos-1 && symbol.X <= num.lxPos+1 {
					adjacent = append(adjacent, num)
				}
			}
			if len(adjacent) == 2 {
				gears = append(gears, Gear{Adjacent: adjacent})
			}
		}
	}
	return gears
}

func Part1(symbols []Symbol, nums []Num) int {

	sum := 0
	for _, num := range nums {
		if num.isAdjacent(symbols) {
			sum += num.Value
		}
	}

	return sum
}

func parse(input []string) ([]Symbol, []Num) {
	symbols := make([]Symbol, 0)
	nums := make([]Num, 0)
	for y, line := range input {
		currentNumAsStr := ""
		for x, r := range line {
			if r == '.' {
				continue
			}
			foundSymbol := false

			if match, _ := regexp.MatchString("[^0-9.]", string(r)); match {
				symbols = append(symbols, Symbol{
					Value: r,
					X:     x,
					Y:     y,
				})
				foundSymbol = true
			}

			if foundSymbol {
				continue
			}

			currentNumAsStr += string(r)
			var match bool
			if x != len(line)-1 {
				match, _ = regexp.MatchString("[0-9]+", string(line[x+1]))
				if match {
					continue
				}
			}

			n, _ := strconv.Atoi(currentNumAsStr)
			nums = append(nums, Num{
				Value: n,
				fxPos: x - (len(currentNumAsStr) - 1),
				lxPos: x,
				y:     y,
			})
			currentNumAsStr = ""
		}
	}

	return symbols, nums
}

/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

Your puzzle answer was 546563.

--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.
*/
