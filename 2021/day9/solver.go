package day9

import (
	"log"
	"mrose.de/aoc/utility"
	"strconv"
)

var maxX int
var maxY int

type Coordinate struct {
	X, Y, Value int
}

func NewCoordinate(x, y, value int) *Coordinate {
	return &Coordinate{
		X:     x,
		Y:     y,
		Value: value,
	}
}

func Solve() (result int) {
	inputArr := utility.StrArr(utility.Input(2021, 9))
	maxY = len(inputArr)
	maxX = len(inputArr[0])
	coordinates := createCoordinateSystem(inputArr)
	lowPoints := findLowPoints(coordinates)
	basins := getBasins(lowPoints)
	largestBasins := getTHREELargestBasins(basins)
	//result = getPartOneResult(lowPoints)
	result = getPartTwoResult(largestBasins)
	return
}

func getPartTwoResult(basins interface{}) int {
	return 0
}

func (c *Coordinate) getAdjacents() (t, b, l, r *Coordinate) {
	return nil, nil, nil, nil
}

func getTHREELargestBasins(basins interface{}) interface{} {
	return nil
}

func getBasins(points []*Coordinate) interface{} {
	return nil
}

func getPartOneResult(points []*Coordinate) (sum int) {
	for _, point := range points {
		sum += point.Value + 1
	}
	return
}

func findLowPoints(coordinates [][]int) (lowPoints []*Coordinate) {
	lowPoints = make([]*Coordinate, 0)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			currentValue := coordinates[y][x]
			if x < maxX-1 {
				if currentValue >= coordinates[y][x+1] {
					continue
				}
			}
			if x > 0 {
				if currentValue >= coordinates[y][x-1] {
					continue
				}
			}
			if y < maxY-1 {
				if currentValue >= coordinates[y+1][x] {
					continue
				}
			}
			if y > 0 {
				if currentValue >= coordinates[y-1][x] {
					continue
				}
			}
			lowPoints = append(lowPoints, NewCoordinate(x, y, currentValue))
		}
	}
	return
}

func createCoordinateSystem(arr []string) (coordinates [][]int) {
	coordinates = make([][]int, maxY)
	for y, v := range arr {
		coordinates = append(coordinates, make([]int, maxX))
		for x := 0; x < maxX; x++ {
			value, err := strconv.Atoi(v[x : x+1])
			if err != nil {
				log.Fatalf("error parsing number\nerr: %e", err)
			}
			coordinates[y] = append(coordinates[y], value)
		}
	}
	return
}
