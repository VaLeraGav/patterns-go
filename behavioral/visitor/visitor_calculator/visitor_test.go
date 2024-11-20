package visitor_calculator

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)

	// Calculating area for square
	// Calculating area for circle
	// Calculating area for rectangle
	// Calculating middle point coordinates for square
	// Calculating middle point coordinates for circle
	// Calculating middle point coordinates for rectangle
}
