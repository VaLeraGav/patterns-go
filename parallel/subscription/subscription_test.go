package subscription

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSubscription(t *testing.T) {
	subscribe()
	waiteSubscribe()
}

func subscribe() {
	pubsub := NewPubSub()

	// Подписчик 1
	pubsub.Subscribe("event1", func(event Event) {
		fmt.Printf("Подписчик 1 получил событие: %s с данными: %v\n", event.Name, event.Data)
	})

	// Подписчик 2
	pubsub.Subscribe("event1", func(event Event) {
		fmt.Printf("Подписчик 2 получил событие: %s с данными: %v\n", event.Name, event.Data)
	})

	// Publish: Отправляет событие всем подписчикам. Обработка событий происходит асинхронно с помощью горутин.
	pubsub.Publish(Event{Name: "event1", Data: "Некоторые данные"})

	// Подписчик 1 получил событие: event1 с данными: Некоторые данные
	// Подписчик 2 получил событие: event1 с данными: Некоторые данные
}

func waiteSubscribe() {
	pubsub := NewPubSub()

	var wg sync.WaitGroup
	wg.Add(2) // Ожидаем 2 подписчика

	// Подписчик 1 с ожиданием завершения
	pubsub.Subscribe("event1", func(event Event) {
		defer wg.Done()
		fmt.Printf("Подписчик 1 (с ожиданием) получил событие: %s с данными: %v\n", event.Name, event.Data)
	})

	// Подписчик 2 с ожиданием завершения
	pubsub.Subscribe("event1", func(event Event) {
		defer wg.Done()
		fmt.Printf("Подписчик 2 (с ожиданием) получил событие: %s с данными: %v\n", event.Name, event.Data)
	})

	// Запланируем публикацию события через 3 секунды
	go func() {
		time.Sleep(5 * time.Second) // Ждем 3 секунды
		pubsub.Publish(Event{Name: "event1", Data: "Дополнительные данные"})
	}()

	// // Ждем завершения всех горутин
	wg.Wait()

	// Подписчик 2 (с ожиданием) получил событие: event1 с данными: Дополнительные данные
	// Подписчик 1 (с ожиданием) получил событие: event1 с данными: Дополнительные данные
}
