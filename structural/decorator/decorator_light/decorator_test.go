package decorator_light

import (
	"fmt"
	"testing"
)

func TestDecoratorLight(t *testing.T) {
	decorator := &ConcreteDecorator{&ConcreteComponent{}}

	result := decorator.Operation()

	fmt.Println(result)
	// <strong>I am component!</strong>
}
