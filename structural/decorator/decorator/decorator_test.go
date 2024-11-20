package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	component := NewConcreteComponent("  Hello, Decorator Pattern!  ")

	// Оборачиваем в декораторы
	trimDecorator := NewTrimDecorator(component)
	upperCaseDecorator := NewUpperCaseDecorator(component)

	// Выполняем операцию
	fmt.Println(trimDecorator.Operation())      //  Hello, Decorator Pattern!
	fmt.Println(upperCaseDecorator.Operation()) //    HELLO, DECORATOR PATTERN!
}
