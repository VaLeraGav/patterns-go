package facade

import (
	"strings"
)

// NewMan creates man.
func NewMan(house *House, tree *Tree) *Man {
	return &Man{
		house: house,
		tree:  tree,
	}
}

// Man implements man and facade.
type Man struct {
	house *House
	tree  *Tree
}

// Todo returns that man must do.
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
	}
	return strings.Join(result, "\n")
}

// ------------ House ------------
type House struct {
}

func (h *House) Build() string {
	return "Build house"
}

// ------------ Tree ------------
type Tree struct {
}

func (t *Tree) Grow() string {
	return "Tree grow"
}
