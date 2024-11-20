package subscription

import "sync"

// Event представляет событие, на которое могут подписываться подписчики.
type Event struct {
	Name string
	Data interface{}
}

// Subscriber представляет подписчика, который получает уведомления о событиях.
type Subscriber func(event Event)

// PubSub реализует паттерн подписки.
type PubSub struct {
	subscribers map[string][]Subscriber
	mu          sync.RWMutex
}

// NewPubSub создает новый экземпляр PubSub.
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]Subscriber),
	}
}

// Subscribe добавляет нового подписчика на событие.
func (ps *PubSub) Subscribe(eventName string, subscriber Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.subscribers[eventName] = append(ps.subscribers[eventName], subscriber)
}

// Publish отправляет событие всем подписчикам.
func (ps *PubSub) Publish(event Event) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if subs, ok := ps.subscribers[event.Name]; ok {
		for _, sub := range subs {
			go sub(event) // Обрабатываем событие асинхронно
		}
	}
}
