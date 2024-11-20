package pipeline

// Add добавляет 2 к каждому элементу входного канала.
func Add(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			result := value + 2

			select {
			case <-doneCh: // если нужно завершить горутину
				return
			case resultCh <- result: // отправляем результат
			}
		}
	}()

	return resultCh
}

// FilterEven фильтрует только четные числа.
func FilterEven(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			// Проверяем, является ли число четным
			if value%2 == 0 {
				select {
				case <-doneCh:
					return
				case resultCh <- value: // Отправляем только четные числа
				}
			}
		}
	}()

	return resultCh
}

// Generator отправляет данные в канал.
func Generator(doneCh chan struct{}, numbers []int) chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)

		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case outputCh <- num:
			}
		}
	}()

	return outputCh
}

// НЕ РАЗОБРАЛСЯ: после него вылетают ошибки
// Pipeline запускает цепочку обработки данных и останавливает все горутины, если встречается число 6.
func CheckSix(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			if value == 6 {
				close(doneCh) // Останавливаем все горутины
				return
			}
			resultCh <- value
		}
	}()

	return resultCh
}
