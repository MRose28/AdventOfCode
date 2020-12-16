package day16

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type field struct {
	name   string
	start  int
	end    int
	xStart int
	xEnd   int
}

var ruleMap map[int]string

func initializeRuleMap() map[int]string {
	return make(map[int]string, 0)
}

type ticket struct {
	fields map[int]string
	valid  bool
}

func Solve() (result int) {
	rules, _, otherTickets := prepare()
	result, otherTickets = validateTickets(rules, otherTickets)

	return
}

func validateTickets(rules []field, tickets []ticket) (result int, updatedTickets []ticket) {
	for _, ticket := range tickets {
		var tmp int
		tmp, ticket = validateTicket(ticket, rules)
		result += tmp
		if ticket.valid {
			updatedTickets = append(updatedTickets, ticket)
		}
	}
	return
}

func validateTicket(ticket ticket, rules []field) (sumInvalidNumbers int, updatedTicket ticket) {
	for k := range ticket.fields {
		ok, _ := fieldName(k, rules)
		if !ok {
			ticket.valid = false
			sumInvalidNumbers += k
			continue
		}
		ticket.valid = true
	}
	updatedTicket = ticket
	return
}

func fieldName(value int, rules []field) (ok bool, name string) {
	for _, rule := range rules {
		if (value >= rule.start && value < rule.xStart) || (value > rule.xEnd && value <= rule.end) {
			return true, rule.name
		}
	}

	return false, "not found"
}

func prepare() (rules []field, myTicket ticket, otherTickets []ticket) {
	rulesIn := utility.StrArr(utility.InputToString("2020/assets/day16/rules.txt"))
	myTicketIn := utility.InputToString("2020/assets/day16/myTicket.txt")
	otherTicketsIn := utility.InputToString("2020/assets/day16/otherTickets.txt")

	rules = parseRules(rulesIn)
	myTicket = parseTicket(myTicketIn)
	otherTickets = parseTickets(otherTicketsIn)
	return
}

func parseTickets(in string) (tickets []ticket) {
	for _, v := range strings.Split(in, "\n") {
		tickets = append(tickets, parseTicket(v))
	}
	return
}

func parseTicket(in string) (ticket ticket) {
	ticket.fields = make(map[int]string, 0)
	ticket.valid = false
	for _, v := range strings.Split(in, ",") {
		index, _ := strconv.Atoi(v)
		ticket.fields[index] = "unknown"
	}
	return
}

func parseRules(in []string) (rules []field) {
	for _, fieldRule := range in {
		start, end, xEnd, xStart := parseRuleValues(fieldRule)
		field := field{
			name:   strings.Split(fieldRule, ":")[0],
			start:  start,
			end:    end,
			xStart: xStart,
			xEnd:   xEnd,
		}
		rules = append(rules, field)
	}
	return
}

func parseRuleValues(rule string) (start int, end int, xEnd int, xStart int) {
	intPart := strings.Split(rule, ": ")[1]
	start, _ = strconv.Atoi(strings.Split(intPart, "-")[0])
	xStart, _ = strconv.Atoi(strings.Split(strings.Split(intPart, "-")[1], " ")[0])
	xStart++
	xEnd, _ = strconv.Atoi(strings.Split(strings.Split(intPart, " or ")[1], "-")[0])
	xEnd--
	arrSize := len(strings.Split(intPart, "-"))
	end, _ = strconv.Atoi(strings.Split(intPart, "-")[arrSize-1])
	return
}
