package observer

import "fmt"

// Интерфейс субъекта
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

// Конкретный субъект
type ConcreteSubject struct {
	observers []Observer
	state     string
}

func NewSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Detach(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	fmt.Printf("Состояние субъекта изменено на: %s\n", state)
	s.Notify()
}
