package faninandfanout

import "time"

// add — добавляет 1 к каждому значению
func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)
		for num := range inputCh {
			// Имитация более затратной работы
			time.Sleep(time.Second)
			result := num + 1

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}
	}()

	return resultStream
}

// fanOut — создает несколько горутин add для параллельной обработки данных
func FanOut(doneCh chan struct{}, inputCh chan int, workers int) []chan int {
	resultChannels := make([]chan int, workers)

	for i := 0; i < workers; i++ {
		resultChannels[i] = add(doneCh, inputCh)
	}

	return resultChannels
}
