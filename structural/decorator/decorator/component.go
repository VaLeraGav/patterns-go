package decorator

// Component is the interface that defines the method to be decorated.
type Component interface {
	Operation() string
}
