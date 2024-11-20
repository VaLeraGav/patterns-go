package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	subject := NewSubject()

	observer1 := NewObserver("Наблюдатель 1")
	observer2 := NewObserver("Наблюдатель 2")

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.SetState("Состояние 1")
	subject.SetState("Состояние 2")

	subject.Detach(observer1)

	subject.SetState("Состояние 3")

	// 	Состояние субъекта изменено на: Состояние 1
	// Наблюдатель 1 обновлен с состоянием: Состояние 1
	// Наблюдатель 2 обновлен с состоянием: Состояние 1
	// Состояние субъекта изменено на: Состояние 2
	// Наблюдатель 1 обновлен с состоянием: Состояние 2
	// Наблюдатель 2 обновлен с состоянием: Состояние 2
	// Состояние субъекта изменено на: Состояние 3
	// Наблюдатель 2 обновлен с состоянием: Состояние 3
}
