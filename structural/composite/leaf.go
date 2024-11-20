package composite

import "fmt"

// Leaf represents a simple menu item.
type Leaf struct {
	name string
}

// NewLeaf creates a new Leaf instance.
func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

// Operation implements the Component interface for Leaf.
func (l *Leaf) Operation() string {
	return fmt.Sprintf("Leaf: %s", l.name)
}
