package day12

type node struct {
	up, down, left, right *node
	height                int
	target                bool
	visited               bool
	steps                 int
}

func (n *node) neighbours() (neighbours []*node) {
	neighbours = []*node{n.left, n.right, n.up, n.down}

	return
}
func (n *node) u() *node {

	return n.up
}

func (n *node) d() *node {

	return n.down
}

func (n *node) l() *node {

	return n.left
}

func (n *node) r() *node {

	return n.right
}
