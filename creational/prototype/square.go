package prototype

import "fmt"

// Square структура для квадрата
type Square struct {
	side float64
}

func (s *Square) Clone() Shape {
	return &Square{side: s.side}
}

func (s *Square) Draw() string {
	return fmt.Sprintf("Square with side: %.2f", s.side)
}
