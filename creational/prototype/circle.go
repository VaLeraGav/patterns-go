package prototype

import "fmt"

// Circle структура для круга
type Circle struct {
	radius float64
}

func (c *Circle) Clone() Shape {
	return &Circle{radius: c.radius}
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Circle with radius: %.2f", c.radius)
}
