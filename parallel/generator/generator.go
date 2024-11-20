package generator

func Generator(items []int) chan int {
	ch := make(chan int)

	go func() {
		// Закрываем канал после завершения отправки данных
		defer close(ch)

		// Перебираем элементы и отправляем их в канал
		for _, item := range items {
			ch <- item
		}
	}()

	return ch
}
