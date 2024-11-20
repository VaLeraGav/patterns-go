package map_filter

import "fmt"

// Map применяет функцию к каждому элементу входного канала и отправляет результат в выходной канал.
func Map(doneCh chan struct{}, inputCh <-chan int, mapper func(int) int) <-chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for {
			select {
			case <-doneCh:
				fmt.Println("Получен сигнал завершения, выходим из Map")
				return
			case value, ok := <-inputCh:
				if !ok { // Проверяем, закрыт ли канал
					fmt.Println("inputCh закрыт, выходим из Map")
					return
				}
				resultCh <- mapper(value) // Применяем функцию и отправляем результат
			}
		}
	}()

	return resultCh
}

// Filter применяет функцию к каждому элементу входного канала и отправляет элементы, удовлетворяющие условию, в выходной канал.
func Filter(doneCh chan struct{}, inputCh <-chan int, predicate func(int) bool) <-chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for {
			select {
			case <-doneCh:
				fmt.Println("Получен сигнал завершения, выходим из Filter")
				return
			case value, ok := <-inputCh:
				if !ok { // Проверяем, закрыт ли канал
					fmt.Println("inputCh закрыт, выходим из Filter")
					return
				}
				if predicate(value) { // Проверяем условие
					resultCh <- value // Отправляем элемент, если он удовлетворяет условию
				}
			}
		}
	}()

	return resultCh
}

// Generator создает последовательность целых чисел.
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
