package mediator_colleague

import "fmt"

// Интерфейс коллеги
type Colleague interface {
	Send(message string)
	Receive(message string)
}

// ------------ Конкретный коллега 1 ------------
type Colleague1 struct {
	mediator Mediator
}

func NewColleague1(mediator Mediator) *Colleague1 {
	return &Colleague1{mediator: mediator}
}

func (c *Colleague1) Send(message string) {
	fmt.Println("Коллега 1 отправляет:", message)
	c.mediator.Send(message, c)
}

func (c *Colleague1) Receive(message string) {
	fmt.Println("Коллега 1 получил:", message)
}

// ------------ Конкретный коллега 2 ------------
type Colleague2 struct {
	mediator Mediator
}

func NewColleague2(mediator Mediator) *Colleague2 {
	return &Colleague2{mediator: mediator}
}

func (c *Colleague2) Send(message string) {
	fmt.Println("Коллега 2 отправляет:", message)
	c.mediator.Send(message, c)
}

func (c *Colleague2) Receive(message string) {
	fmt.Println("Коллега 2 получил:", message)
}
