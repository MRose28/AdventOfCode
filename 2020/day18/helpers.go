package day18

import (
	"strings"
)

type InputLine struct {
	input string
}

func (l *InputLine) getResult() (result int) {
	//for {
	//	term, pTerm := NextTerm(l.input)
	//	if term == "-1" {
	//		break
	//	}
	//	for i, v := range term {
	//		if num, err := strconv.Atoi(v); err == nil {
	//
	//		}
	//	}
	//}
	return
}

func NextTerm(s string) (term, pTerm string) {
	if !strings.Contains(s, "(") {
		return "-1", "-1"
	}
	start := strings.LastIndex(s, "(")
	end := strings.Index(s, ")") + 1
	pTerm = s[start:end]
	term = s[start+1 : end-1]
	print(term)
	return
}
