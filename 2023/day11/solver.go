package day11

import (
	"github.com/google/uuid"
	"mrose.de/aoc/utility"
	"slices"
)

const (
	empty  = "."
	galaxy = "#"
)

type Point struct {
	x, y  int
	value rune
}

type GMap struct {
	coordinates []Point
}

type TestStr struct {
	id *uuid.UUID
	s  string
}

func Solve() (p1, p2 int) {
	input := utility.InputAsStrArr(2023, 11, true)

	m := gmap(input)
	m = expandEmptyRowsCols(m)

	p1 = len(input)

	return
}

func expandEmptyRowsCols(points []*Point) []*Point {
	// true if taken
	takenX := make(map[int]bool)
	takenY := make(map[int]bool)
	biggestX := 0
	biggestY := 0
	var insertionsX, insertionY int

	for _, point := range points {
		var taken bool

		if point.x > biggestX {
			biggestX = point.x
		}
		if point.y > biggestY {
			biggestY = point.y
		}

		if point.value == '#' {
			taken = true
		}
		if !takenX[point.x] {
			takenX[point.x] = taken
		}
		if !takenY[point.y] {
			takenY[point.y] = taken
		}
	}

	for _, t := range takenX {
		if t {
			insertionsX++
		}
	}
	for _, t := range takenY {
		if t {
			insertionY++
		}
	}

	for x, taken := range takenX {
		if !taken {
			for _, point := range points {
				if point.x >= x {
					point.x += 1
				}
			}
			for i := 0; i < biggestY+insertionY; i++ {
				points = append(points, &Point{
					x:     x,
					y:     i,
					value: '.',
				})
			}
		}
	}

	for y, taken := range takenY {
		if !taken {
			for _, point := range points {
				if point.y >= y {
					point.y += 1
				}
			}
			for i := 0; i < biggestX+insertionsX; i++ {
				points = append(points, &Point{
					x:     i,
					y:     y,
					value: '.',
				})
			}
		}
	}

	slices.SortFunc(points, func(a, b *Point) int {
		if a.y < b.y && a.x < b.x {
			return -1
		}
		if a.y > b.y {
			return 1
		}
		if a.y == b.y && a.x < b.x {
			return -1
		}
		if a.y == b.y && a.x > b.x {
			return 1
		}
		return 0
	})

	return points
}

func gmap(input []string) []*Point {
	points := make([]*Point, 0)

	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				points = append(points, &Point{
					x:     x,
					y:     y,
					value: '#',
				})
			}
			if r == '.' {
				points = append(points, &Point{
					x:     x,
					y:     y,
					value: '.',
				})
			}
		}
	}

	return points
}

/*
--- Day 11: Cosmic Expansion ---
You continue following signs for "Hot Springs" and eventually come across an observatory. The Elf within turns out to be a researcher studying cosmic expansion using the giant telescope here.

He doesn't know anything about the missing machine parts; he's only visiting for this research project. However, he confirms that the hot springs are the next-closest area likely to have people; he'll even take you straight there once he's done with today's observation analysis.

Maybe you can help him with the analysis to speed things up?

The researcher has collected a bunch of data and compiled the data into a single giant image (your puzzle input). The image includes empty space (.) and galaxies (#). For example:

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
The researcher is trying to figure out the sum of the lengths of the shortest path between every pair of galaxies. However, there's a catch: the universe expanded in the time it took the light from those galaxies to reach the observatory.

Due to something involving gravitational effects, only some space expands. In fact, the result is that any rows or columns that contain no galaxies should all actually be twice as big.

In the above example, three columns and two rows contain no galaxies:

   v  v  v
 ...#......
 .......#..
 #.........
>..........<
 ......#...
 .#........
 .........#
>..........<
 .......#..
 #...#.....
   ^  ^  ^
These rows and columns need to be twice as big; the result of cosmic expansion therefore looks like this:

....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
Equipped with this expanded universe, the shortest path between every pair of galaxies can be found. It can help to assign every galaxy a unique number:

....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......
In these 9 galaxies, there are 36 pairs. Only count each pair once; order within the pair doesn't matter. For each pair, find any shortest path between the two galaxies using only steps that move up, down, left, or right exactly one . or # at a time. (The shortest path between two galaxies is allowed to pass through another galaxy.)

For example, here is one of the shortest paths between galaxies 5 and 9:

....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......
This path has length 9 because it takes a minimum of nine steps to get from galaxy 5 to galaxy 9 (the eight locations marked # plus the step onto galaxy 9 itself). Here are some other example shortest path lengths:

Between galaxy 1 and galaxy 7: 15
Between galaxy 3 and galaxy 6: 17
Between galaxy 8 and galaxy 9: 5
In this example, after expanding the universe, the sum of the shortest path between all 36 pairs of galaxies is 374.

Expand the universe, then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?
*/
