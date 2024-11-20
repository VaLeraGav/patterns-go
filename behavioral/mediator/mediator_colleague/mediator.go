package mediator_colleague

// Интерфейс посредника
type Mediator interface {
	Send(message string, colleague Colleague)
}

// Конкретный посредник
type ConcreteMediator struct {
	colleague1 Colleague
	colleague2 Colleague
}

func NewMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

func (m *ConcreteMediator) SetColleague1(c Colleague) {
	m.colleague1 = c
}

func (m *ConcreteMediator) SetColleague2(c Colleague) {
	m.colleague2 = c
}

func (m *ConcreteMediator) Send(message string, colleague Colleague) {
	if colleague == m.colleague1 {
		m.colleague2.Receive(message)
	} else {
		m.colleague1.Receive(message)
	}
}
