package mediator_colleague

import (
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := NewMediator()

	colleague1 := NewColleague1(mediator)
	colleague2 := NewColleague2(mediator)

	mediator.SetColleague1(colleague1)
	mediator.SetColleague2(colleague2)

	colleague1.Send("Привет от Коллеги 1")
	colleague2.Send("Привет от Коллеги 2")

	// Коллега 1 отправляет: Привет от Коллеги 1
	// Коллега 2 получил: Привет от Коллеги 1
	// Коллега 2 отправляет: Привет от Коллеги 2
	// Коллега 1 получил: Привет от Коллеги 2
}
