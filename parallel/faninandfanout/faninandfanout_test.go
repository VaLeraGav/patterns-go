package faninandfanout

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFanInAndFanOut(t *testing.T) {

	// Создаем несколько каналов
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := Generator(doneCh, numbers)

	// создаем 10 горутин add с помощью fanOut
	channels := FanOut(doneCh, inputCh, 10)

	// объединяем результаты из всех каналов
	addResultCh := FanIn(doneCh, channels...)

	// // передаем результат в следующий этап multiply
	// resultCh := multiply(doneCh, addResultCh)

	// выводим результаты
	for result := range addResultCh {
		fmt.Println(result)
	}
	// 10
	// 7
	// 6
	// 11
	// 5
	// 8
	// 9
	// 2
	// 3
	// 4
}

// Генерирует случайные числа и отправляет их в канал.
func generateNumbers(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)                                        // Генерируем случайное число
		ch <- num                                                    // Отправляем число в канал
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Имитация работы
	}
}

func TestFanIn(t *testing.T) {
	// Создаем несколько каналов
	ch1 := make(chan int)
	ch2 := make(chan int)
	doneCh := make(chan struct{})

	var wg sync.WaitGroup

	// Запускаем несколько горутин для генерации чисел
	wg.Add(2)
	go generateNumbers(1, ch1, &wg)
	go generateNumbers(2, ch2, &wg)

	// Объединяем каналы с помощью Fan-In
	out := FanIn(doneCh, ch1, ch2)

	// Читаем из выходного канала
	go func() {
		for num := range out {
			fmt.Println("Получено число:", num)

			// Получено число: 10
			// Получено число: 70
			// Получено число: 61
		}
	}()

	// Ждем завершения генерации
	wg.Wait()
	close(ch1)
	close(ch2)

	// Закрываем doneCh для завершения горутин в fanIn
	close(doneCh)

	// Ждем завершения чтения
	time.Sleep(time.Second)
}
