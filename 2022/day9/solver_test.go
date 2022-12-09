package day9

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 13
	result, _ := Solve()
	if want != result {
		t.Fatalf(`day9.Solve() == %d, want match for %d`, result, want)
	}
}

func TestPart2(t *testing.T) {

}
