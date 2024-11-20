package visitor_city

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	city := &City{}

	// Добавляем места
	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	people := &People{}

	// Посещаем все места в городе
	result := city.Accept(people)
	fmt.Println(result)
	// Buy sushi...Buy pizza...Buy burger...
}
