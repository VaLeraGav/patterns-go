package composite

import "fmt"

// Composite represents a complex menu item that can contain other components.
type Composite struct {
	name     string
	children []Component
}

// NewComposite creates a new Composite instance.
func NewComposite(name string) *Composite {
	return &Composite{name: name}
}

// Add adds a child component to the composite.
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

// Operation implements the Component interface for Composite.
func (c *Composite) Operation() string {
	result := fmt.Sprintf("Composite: %s\n", c.name)
	for _, child := range c.children {
		result += "  " + child.Operation() + "\n"
	}
	return result
}
