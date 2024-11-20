package faninandfanout

import "sync"

func FanIn(doneCh chan struct{}, channels ...chan int) chan int {
	finalStream := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		chCopy := ch
		wg.Add(1)

		go func() {
			defer wg.Done()
			for value := range chCopy {
				select {
				case <-doneCh:
					return
				case finalStream <- value:
				}
			}
		}()
	}

	// Удаление каналов
	go func() {
		wg.Wait()
		close(finalStream)
	}()

	return finalStream
}
