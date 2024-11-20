package faninandfanout

func Generator(doneCh chan struct{}, numbers []int) chan int {
	dataStream := make(chan int)

	go func() {
		defer close(dataStream)
		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case dataStream <- num:
			}
		}
	}()

	return dataStream
}
