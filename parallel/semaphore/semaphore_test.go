package semaphore

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	var wg sync.WaitGroup

	// создаем семафор, который позволит работать только двум горутинам одновременно
	const maxConcurrentWorkers = 2
	semaphore := NewSemaphore(maxConcurrentWorkers)

	// запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(taskID int) {
			// резервируем место в семафоре перед началом работы
			semaphore.Acquire()

			// когда горутина завершает работу, освобождаем место и уменьшаем счетчик WaitGroup
			defer wg.Done()
			defer semaphore.Release()

			// симулируем работу горутины
			log.Printf("Запущен рабочий %d", taskID)
			time.Sleep(1 * time.Second)
		}(i)
	}

	// ждем завершения всех горутин
	wg.Wait()
}

func TestSemaphoreLight(t *testing.T) {
	const maxConcurrentWorkers = 2

	sem := make(chan struct{}, maxConcurrentWorkers) // Создаем семафор с заданным количеством ресурсов
	var wg sync.WaitGroup

	// запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Worker(i, sem, &wg) // Запускаем рабочую горутину
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Все рабочие горутины завершены.")
}

func Worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Запрашиваем доступ к ресурсу
	sem <- struct{}{}        // Захватываем семафор
	defer func() { <-sem }() // Освобождаем семафор

	// Имитация работы
	fmt.Printf("Worker %d начал работу\n", id)
	time.Sleep(2 * time.Second) // Имитация длительной работы
	fmt.Printf("Worker %d закончил работу\n", id)
}
