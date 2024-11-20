package observer

import "fmt"

// Интерфейс наблюдателя
type Observer interface {
	Update(state string)
}

// Конкретный наблюдатель
type ConcreteObserver struct {
	name string
}

func NewObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (o *ConcreteObserver) Update(state string) {
	fmt.Printf("%s обновлен с состоянием: %s\n", o.name, state)
}
