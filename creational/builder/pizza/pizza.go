package pizza

import "fmt"

// Pizza представляет собой сложный объект
type Pizza struct {
	Size     string
	Dough    string
	Toppings []string
}

// Метод для отображения информации о пицце
func (p *Pizza) String() string {
	return fmt.Sprintf("Pizza [Size: %s, Dough: %s, Toppings: %v]", p.Size, p.Dough, p.Toppings)
}
