package day16

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type rule struct {
	name   string
	start  int
	end    int
	xStart int
	xEnd   int
}

type ticket struct {
	fields []int
	valid  bool
}

type field struct {
	name   string
	minMax minMax
}

type minMax struct {
	min int
	max int
}

var fieldList []field

//Min and max value for 1 rule have to fit one category
func initializeFieldMap(tickets []ticket, rules []rule) {
	fieldList = make([]field, 0)
	createFieldList(tickets, rules)
	//setFieldList(minMaxMap, rules)
}

func createFieldList(tickets []ticket, rules []rule) {
	valid := true
	var name string
	var ok bool
	for i := 0; i < 3; i++ {
		valid = true
		for _, ticket := range tickets {
			if ticket.fields[i] > 1000 {
				fmt.Printf("%d", ticket.fields[i])
			}
			ok, name = fieldName(ticket.fields[i], rules)
			if !ok {
				valid = false
			}
		}
		if valid {
			fieldList[i] = field{
				name:   name,
				minMax: minMax{},
			}
		}
	}
}

func setFieldList(minMaxMap map[int]minMax, rules []rule) {
	possibleFields := make([]rule, 0)
	runAgain := make(map[int]minMax, 0)

	for i := 0; i < len(minMaxMap); i++ {
		possibleFields = findField(minMaxMap[i], rules)
		if len(possibleFields) < 1 {
			runAgain[i] = minMaxMap[i]
		} else if len(possibleFields) > 1 {
			runAgain[i] = minMaxMap[i]
		} else {
			if len(possibleFields) == 1 {
				fmt.Println(possibleFields[0])
			}
			fieldList = append(fieldList, field{
				name:   possibleFields[0].name,
				minMax: minMaxMap[i],
			})
		}
		possibleFields = make([]rule, 0)
	}
	if len(runAgain) != 0 {
		minMaxMap = runAgain
		setFieldList(minMaxMap, rules)
	}
}

func findField(minMax minMax, rules []rule) (possibleFields []rule) {
	var possibleField rule
	for _, field := range rules {
		if fieldListContainsField(field.name) {
			continue
		}
		if (minMax.min >= field.start && minMax.min < field.xStart) ||
			(minMax.min > field.xEnd && minMax.min <= field.end) {
			possibleField = field
		}
		if (minMax.max >= field.start && minMax.max < field.xStart) ||
			(minMax.max > field.xEnd && minMax.max <= field.end) {
			if possibleField == field {
				possibleFields = append(possibleFields, possibleField)
			}
		}
	}
	return
}

func fieldListContainsField(name string) bool {
	for _, v := range fieldList {
		if v.name == name {
			return true
		}
	}
	return false
}

func minMaxList(tickets []ticket) (minMaxMap map[int]minMax) {
	minMaxMap = make(map[int]minMax, 0)
	counter := 0
	newMinMax := minMax{
		min: 0,
		max: 0,
	}
	for _, ticket := range tickets {
		counter = 0
		for _, v := range ticket.fields {
			_, ok := minMaxMap[counter]
			if !ok {
				minMaxMap[counter] = minMax{
					min: v,
					max: v,
				}
			} else {
				if minMaxMap[counter].max < v {
					newMinMax.max = v
				} else {
					newMinMax.max = minMaxMap[counter].max
				}
				if minMaxMap[counter].min > v {
					newMinMax.min = v
				} else {
					newMinMax.min = minMaxMap[counter].min
				}

				minMaxMap[counter] = newMinMax
			}
			counter++
		}
	}
	return
}

func Solve() (result int) {
	rules, myTicket, otherTickets := prepare()
	_, otherTickets = validateTickets(rules, otherTickets)
	initializeFieldMap(append(otherTickets, myTicket), rules)

	return
}

func validateTickets(rules []rule, tickets []ticket) (result int, updatedTickets []ticket) {
	updatedTickets = make([]ticket, 0)
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

func validateTicket(currentTicket ticket, rules []rule) (sumInvalidNumbers int, updatedTicket ticket) {
	updatedTicket = ticket{
		fields: make([]int, 0),
		valid:  false,
	}
	for _, k := range currentTicket.fields {
		ok, _ := fieldName(k, rules)
		if !ok {
			currentTicket.valid = false
			sumInvalidNumbers += k
			continue
		}
		currentTicket.valid = true
	}
	updatedTicket = currentTicket
	return
}

func fieldName(value int, rules []rule) (ok bool, name string) {
	if value == 991 {
		print("whats this???")
	}
	possibleRules := make([]rule, 0)
	for _, rule := range rules {
		if fieldListContainsField(rule.name) {
			continue
		}
		if (value >= rule.start && value < rule.xStart) || (value > rule.xEnd && value <= rule.end) {
			possibleRules = append(possibleRules, rule)
		} else {
			fmt.Printf("Value: %d, not fitting %s\n\n", value, rule.name)
		}
	}
	if len(possibleRules) == 0 {
		fmt.Printf("0 fields for value %d\n\n", value)
	}

	if len(possibleRules) == 1 {
		return true, possibleRules[0].name
	}
	return false, "not found"
}

func prepare() (rules []rule, myTicket ticket, otherTickets []ticket) {
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
	ticket.fields = make([]int, 0)
	ticket.valid = false
	for _, v := range strings.Split(in, ",") {
		value, _ := strconv.Atoi(v)
		ticket.fields = append(ticket.fields, value)
	}
	return
}

func parseRules(in []string) (rules []rule) {
	for _, fieldRule := range in {
		start, end, xEnd, xStart := parseRuleValues(fieldRule)
		field := rule{
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

/*
--- Day 16: Ticket Translation ---
As you're walking to yet another connecting flight, you realize that one of the legs of your re-routed trip coming up is on a high-speed train. However, the train ticket you were given is in a language you don't understand. You should probably figure out what it says before you get to the train station after the next flight.

Unfortunately, you can't actually read the words on the ticket. You can, however, read the numbers, and so you figure out the fields these tickets must have and the valid ranges for values in those fields.

You collect the rules for ticket fields, the numbers on your ticket, and the numbers on other nearby tickets for the same train service (via the airport security cameras) together into a single document you can reference (your puzzle input).

The rules for ticket fields specify a list of fields that exist somewhere on the ticket and the valid ranges of values for each field. For example, a rule like class: 1-3 or 5-7 means that one of the fields in every ticket is named class and can be any value in the ranges 1-3 or 5-7 (inclusive, such that 3 and 5 are both valid in this field, but 4 is not).

Each ticket is represented by a single line of comma-separated values. The values are the numbers on the ticket in the order they appear; every ticket has the same format. For example, consider this ticket:

.--------------------------------------------------------.
| ????: 101    ?????: 102   ??????????: 103     ???: 104 |
|                                                        |
| ??: 301  ??: 302             ???????: 303      ??????? |
| ??: 401  ??: 402           ???? ????: 403    ????????? |
'--------------------------------------------------------'
Here, ? represents text in a language you don't understand. This ticket might be represented as 101,102,103,104,301,302,303,401,402,403; of course, the actual train tickets you're looking at are much more complicated. In any case, you've extracted just the numbers in such a way that the first number is always the same specific field, the second number is always a different specific field, and so on - you just don't know what each position actually means!

Start by determining which tickets are completely invalid; these are tickets that contain values which aren't valid for any field. Ignore your ticket for now.

For example, suppose you have the following notes:

class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
It doesn't matter which position corresponds to which field; you can identify invalid nearby tickets by considering only whether tickets contain values that are not valid for any field. In this example, the values on the first nearby ticket are all valid for at least one field. This is not true of the other three nearby tickets: the values 4, 55, and 12 are are not valid for any field. Adding together all of the invalid values produces your ticket scanning error rate: 4 + 55 + 12 = 71.

Consider the validity of the nearby tickets you scanned. What is your ticket scanning error rate?

Your puzzle answer was 27802.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Now that you've identified which tickets contain invalid values, discard those tickets entirely. Use the remaining valid tickets to determine which field is which.

Using the valid ranges for each field, determine what order the fields appear on the tickets. The order is consistent between all tickets: if seat is the third field, it is the third field on every ticket, including your ticket.

For example, suppose you have the following notes:

class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
Based on the nearby tickets in the above example, the first position must be row, the second position must be class, and the third position must be seat; you can conclude that in your ticket, class is 12, row is 11, and seat is 13.

Once you work out which field is which, look for the six fields on your ticket that start with the word departure. What do you get if you multiply those six values together?
*/
