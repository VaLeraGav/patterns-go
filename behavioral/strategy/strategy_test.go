package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	basePrice := 100.0

	context := &Context{}

	// Используем стратегию без налога
	context.SetStrategy(&ConcreteStrategyA{})
	fmt.Printf("Цена без налога: %.2f\n", context.Calculate(basePrice))

	// Используем стратегию с налогом 10%
	context.SetStrategy(&ConcreteStrategyB{})
	fmt.Printf("Цена с налогом 10%%: %.2f\n", context.Calculate(basePrice))

	// Используем стратегию с налогом 20%
	context.SetStrategy(&ConcreteStrategyC{})
	fmt.Printf("Цена с налогом 20%%: %.2f\n", context.Calculate(basePrice))
}
