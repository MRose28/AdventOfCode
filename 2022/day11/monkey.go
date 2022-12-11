package day11

type monkey struct {
	items           []*item
	divisibleBy     int
	operation       func(old uint64) uint64
	posId, negId    int
	pos, neg        *monkey
	inspectionCount int
}

func newMonkey() *monkey {
	return &monkey{}
}
func (m *monkey) addItems(items []*item) *monkey {
	m.items = items

	return m
}

func (m *monkey) test(i *item) (result bool) {
	result = i.worryLevel%uint64(m.divisibleBy) == 0

	return
}

func (m *monkey) throw(b bool) {
	if b {
		m.pos.items = append(m.pos.items, m.items[0])
	} else {
		m.neg.items = append(m.neg.items, m.items[0])
	}
	m.items = m.items[1:len(m.items)]
}
