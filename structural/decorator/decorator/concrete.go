package decorator

// ConcreteComponent is the base implementation of the Component interface.
type ConcreteComponent struct {
	message string
}

// NewConcreteComponent creates a new ConcreteComponent instance.
func NewConcreteComponent(message string) *ConcreteComponent {
	return &ConcreteComponent{message: message}
}

// Operation returns the message.
func (c *ConcreteComponent) Operation() string {
	return c.message
}
