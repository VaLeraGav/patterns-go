package pizza

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builderOne := NewConcretePizzaBuilder()
	pizzaOne := builderOne.SetSize("Large").
		SetDough("Thin").
		AddTopping("Cheese").
		AddTopping("Tomato").
		AddTopping("Olives").
		Build()

	fmt.Println(pizzaOne) // Pizza [Size: Large, Dough: Thin, Toppings: [Cheese Tomato Olives]]

	builderTwo := NewConcretePizzaBuilder()
	pizzaTwo := builderTwo.SetSize("Large").
		SetDough("Thin").
		AddTopping("Olives").
		Build()

	fmt.Println(pizzaTwo) // Pizza [Size: Large, Dough: Thin, Toppings: [Olives]]
}
