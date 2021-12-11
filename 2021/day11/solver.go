package day11

import (
	"fmt"
	"mrose.de/aoc/utility"
)

/* 10x10 Grid of Octopuses
O-Energy-Level 0-9
One Step ->
	* All O energy++
	* O at 9+ Energy Flash
	* Adjacent get energy++ (also diagonally)
		* If this increases Energy of another O to 9+ it flashes
		* Whenever a new Octopus flashes continue steps else stop.
	* One FLash/Octopus/Step
	* After flashing E = 0
*/

var flashCounter = 0
var stepCounter = 0

type Octopus struct {
	Energy            int
	Flash, hasFlashed bool
	x, y              int
}

func NewOctopus(x, y, energy int) *Octopus {
	return &Octopus{
		Energy: energy,
		Flash:  false,
		x:      x,
		y:      y,
	}
}

func (o *Octopus) increaseEnergy() {
	if !o.Flash {
		o.Energy++

		if o.Energy > 9 {
			o.Flash = true
			o.Energy = 0
		}
	}
}

func (o *Octopus) activateFlash(og *OctopusGrid) {
	flashCounter++
	o.hasFlashed = true
	neighbours := og.findNeighbours(o.x, o.y)
	for _, neighbour := range neighbours {
		neighbour.increaseEnergy()
	}
	for _, neighbour := range neighbours {
		if neighbour.Flash && !neighbour.hasFlashed {
			neighbour.activateFlash(og)
		}
	}
	o.Flash = false
}

type OctopusGrid [][]*Octopus

func (og *OctopusGrid) getOctopus(x, y int) *Octopus {
	if (x < 0 || y < 0) || x >= len((*og)[0]) || y >= len(*og) {
		return nil
	}
	return (*og)[y][x]
}

func (og *OctopusGrid) findNeighbours(x, y int) []*Octopus {
	oArr := make([]*Octopus, 0)

	o := og.getOctopus(x+1, y)
	if o != nil {
		oArr = append(oArr, o)
	}

	o = og.getOctopus(x-1, y)
	if o != nil {
		oArr = append(oArr, o)
	}
	o = og.getOctopus(x, y+1)
	if o != nil {
		oArr = append(oArr, o)
	}
	o = og.getOctopus(x, y-1)
	if o != nil {
		oArr = append(oArr, o)
	}
	o = og.getOctopus(x+1, y+1)
	if o != nil {
		oArr = append(oArr, o)
	}
	o = og.getOctopus(x+1, y-1)
	if o != nil {
		oArr = append(oArr, o)
	}
	o = og.getOctopus(x-1, y-1)
	if o != nil {
		oArr = append(oArr, o)
	}
	o = og.getOctopus(x-1, y+1)
	if o != nil {
		oArr = append(oArr, o)
	}

	return oArr
}

func (og *OctopusGrid) increaseEnergyOfAll() {
	for _, octopuses := range *og {
		for _, octopus := range octopuses {
			octopus.increaseEnergy()
		}
	}
}

func Solve() (result int) {
	input := utility.StrArr(utility.Input2021Day11())
	octopuses := parseInput(input)

	return octopuses.simulate()
}

func (og *OctopusGrid) simulate() int {
	for {
		stepCounter++
		fmt.Printf("Step %v\n", stepCounter)
		if stepCounter == 142 {
			fmt.Printf("here")
		}
		og.increaseEnergyOfAll()
		og.flashOctopuses()
		if og.allFlashed() {
			break
		}
		og.resetAllFlashers()

	}

	return stepCounter
}

func (og *OctopusGrid) flashOctopuses() {
	for _, octopuses := range *og {
		for _, octopus := range octopuses {
			if octopus.Flash {
				octopus.activateFlash(og)
			}
		}
	}
}

func (og *OctopusGrid) resetAllFlashers() {
	for _, octopuses := range *og {
		for _, octopus := range octopuses {
			if octopus.hasFlashed {
				octopus.Energy = 0
				octopus.hasFlashed = false
				octopus.Flash = false
			}
		}
	}
}

func (og *OctopusGrid) allFlashed() bool {
	for _, octopuses := range *og {
		for _, octopus := range octopuses {
			if octopus.hasFlashed == false {
				return false
			}
		}
	}
	return true
}

func parseInput(input []string) OctopusGrid {
	oGrid := make([][]*Octopus, 0)
	for vi, line := range input {
		os := make([]*Octopus, 0)
		for hi, v := range utility.IntArr(line, "") {
			os = append(os, NewOctopus(hi, vi, v))
		}
		oGrid = append(oGrid, os)
	}
	return oGrid
}
