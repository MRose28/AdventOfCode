package day11

type item struct {
	worryLevel uint64
}

func (i *item) decreaseWorryLevel() *item {
	i.worryLevel /= 3
	return i
}
