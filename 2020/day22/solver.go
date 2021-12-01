package day22

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type card struct {
	prev  *card
	next  *card
	value int
}

type player struct {
	top    *card
	bottom *card
	count  int
}

func Solve() (result int) {
	p1 := player{
		top:    nil,
		bottom: nil,
		count:  0,
	}
	p2 := player{
		top:    nil,
		bottom: nil,
		count:  0,
	}
	inputArr := strings.Split(utility.Input2020Day22(), "\n\n")
	valuesP1 := utility.IntArr(inputArr[0], "\n")
	valuesP2 := utility.IntArr(inputArr[1], "\n")
	initializeHandCards(valuesP1, &p1)
	initializeHandCards(valuesP2, &p2)
	winner := combat(&p1, &p2)
	result = score(winner)
	return
}

func combat(p1 *player, p2 *player) (winner player) {
	history := make(map[string]struct{}, 0)
	for p1.count > 0 && p2.count > 0 {
		if happenedBefore(playerCardsToString(*p1), playerCardsToString(*p2), history) {
			cardToWinner(p1, p2)
			continue
		}
		history[playerCardsToString(*p1)+"-"+playerCardsToString(*p2)] = struct{}{}
		if p1.top.value <= p1.count-1 && p2.top.value <= p2.count-1 {
			newP1 := subPlayer(*p1)
			newP2 := subPlayer(*p2)
			combat(&newP1, &newP2)
			if newP2.count == 0 {
				cardToWinner(p1, p2)
				continue
			}
			if newP1.count == 0 {
				cardToWinner(p2, p1)
				continue
			}
		} else {
			if p1.top.value > p2.top.value {
				cardToWinner(p1, p2)
			} else {
				cardToWinner(p2, p1)
			}
		}
	}
	if p1.count == 0 {
		winner = *p2
	} else {
		winner = *p1
	}
	return
}

func happenedBefore(u string, u2 string, history map[string]struct{}) bool {
	if containsHistory(history, u+"-"+u2) {
		return true
	}
	return false
}

func containsHistory(history map[string]struct{}, s string) bool {
	if _, exists := history[s]; exists {
		return true
	}
	return false
}

func subPlayer(p player) player {
	sPlayer := player{
		top:    nil,
		bottom: nil,
		count:  0,
	}
	cArr := subDeckValues(p)
	initializeHandCards(cArr, &sPlayer)
	return sPlayer
}

func subDeckValues(p player) (result []int) {
	card := p.top.next
	for i := 0; i < p.top.value; i++ {
		result = append(result, card.value)
		card = card.next
	}
	return
}

func cardToWinner(winner *player, loser *player) {
	topToBottom(winner)
	addCard(winner, loser.top)
	removeTopFromLoser(loser)
}

func removeTopFromLoser(loser *player) {
	if loser.count > 1 {
		loser.top = loser.top.next
		loser.top.prev = nil
	} else {
		loser.top = nil
		loser.bottom = nil
	}
	loser.count--
}

func initializeHandCards(values []int, p *player) {
	for _, v := range values {
		c := card{
			prev:  nil,
			next:  nil,
			value: v,
		}
		addCard(p, &c)
	}
}

func score(player player) (result int) {
	multiplier := player.count
	card := player.top
	for i := 0; i < player.count; i++ {
		result += multiplier * card.value
		multiplier--
		card = card.next
	}
	return
}
func addCard(p *player, c *card) {
	if p.count == 0 {
		p.top = c
		p.bottom = c
	} else {
		p.bottom.next = c
		c.prev = p.bottom
		p.bottom = c
	}
	p.count++
}

func topToBottom(p *player) {
	p.bottom.next = p.top
	p.bottom = p.top
	p.top = p.top.next
	p.top.prev = nil
	p.bottom.next = nil
}

func playerCardsToString(p player) string {
	s := ""
	c := p.top
	for i := 0; i < p.count; i++ {
		s += strconv.Itoa(c.value)
		c = c.next
	}
	return s
}
