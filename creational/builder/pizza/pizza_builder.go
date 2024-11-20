package pizza

// PizzaBuilder определяет интерфейс для построителей пиццы
type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetDough(dough string) PizzaBuilder
	AddTopping(topping string) PizzaBuilder
	Build() *Pizza
}

// ConcretePizzaBuilder реализует интерфейс PizzaBuilder
type ConcretePizzaBuilder struct {
	pizza *Pizza
}

// Конструктор для ConcretePizzaBuilder
func NewConcretePizzaBuilder() *ConcretePizzaBuilder {
	return &ConcretePizzaBuilder{&Pizza{}}
}

// Устанавливает размер пиццы
func (b *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

// Устанавливает тесто пиццы
func (b *ConcretePizzaBuilder) SetDough(dough string) PizzaBuilder {
	b.pizza.Dough = dough
	return b
}

// Добавляет топпинг к пицце
func (b *ConcretePizzaBuilder) AddTopping(topping string) PizzaBuilder {
	b.pizza.Toppings = append(b.pizza.Toppings, topping)
	return b
}

// Строит пиццу
func (b *ConcretePizzaBuilder) Build() *Pizza {
	return b.pizza
}
