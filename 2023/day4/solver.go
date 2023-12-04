package day4

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Card struct {
	ID              int
	Winning, MyNums []int
	Instances       int
}

func (c Card) NumOfWinningNums() int {
	wNums := 0
	for _, num := range c.MyNums {
		if utility.ContainsInt(c.Winning, num) {
			wNums++
		}
	}

	return wNums
}

func Solve() (part1, part2 int) {

	input := utility.StrArr(utility.Input(2023, 4))

	cards := parseCards(input)
	part1 = getWinnings(cards)

	part2 = countCards(cards)

	return
}

func countCards(cards []Card) int {
	cardMap := make(map[int]int)
	for _, card := range cards {
		cardMap[card.ID] = 1
	}

	for _, card := range cards {
		for i := card.ID + 1; i <= card.ID+card.NumOfWinningNums(); i++ {
			if cardMap[i] == 0 {
				continue
			}
			cardMap[i] += cardMap[card.ID]
		}
	}
	return sumOfCards(cardMap)
}

func sumOfCards(cardMap map[int]int) int {
	sum := 0
	for _, v := range cardMap {
		sum += v
	}
	return sum
}

func getWinnings(cards []Card) int {
	overallWinnings := 0
	for _, card := range cards {
		overallWinnings += getWinningForCard(card.NumOfWinningNums())
	}

	return overallWinnings
}

func getWinningForCard(num int) int {
	v := 0
	if num == 0 {
		return v
	} else {
		v = 1
	}
	for i := 0; i < num-1; i++ {
		v *= 2
	}

	return v
}

func parseCards(input []string) []Card {
	cards := make([]Card, 0)

	for _, line := range input {
		line = strings.ReplaceAll(line, "   ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		wNums := make([]int, 0)
		mNums := make([]int, 0)
		id, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		wPart := strings.Split(strings.ReplaceAll(strings.Trim(strings.Split(strings.Split(line, ":")[1], "|")[0], " "), "  ", " "), " ")
		for _, s := range wPart {
			n, _ := strconv.Atoi(s)
			wNums = append(wNums, n)
		}
		lPart := strings.Split(strings.ReplaceAll(strings.Trim(strings.Split(strings.Split(line, ":")[1], "|")[1], " "), "  ", " "), " ")
		for _, s := range lPart {
			n, _ := strconv.Atoi(s)
			mNums = append(mNums, n)
		}
		cards = append(cards, Card{
			ID:        id,
			Winning:   wNums,
			MyNums:    mNums,
			Instances: 1,
		})
	}

	return cards
}
