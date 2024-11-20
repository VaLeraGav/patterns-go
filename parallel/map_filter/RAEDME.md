# Map and Filter

Паттерн `Map` в программировании — это один из распространенных паттернов обработки данных, который используется для применения функции к каждому элементу коллекции и сбора результатов. Этот паттерн часто используется в функциональных языках программирования, но также имеет широкое применение в других языках.

Пример элементарного `map`:

```go
// Map применяет функцию к каждому элементу входного канала и отправляет результат в выходной канал.
func Map(inputCh <-chan int, mapper func(int) int) <-chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			resultCh <- mapper(value) // Применяем функцию и отправляем результат
		}
	}()

	return resultCh
}

// Функция для удвоения числа
func double(x int) int {
	return x * 2
}

func main() {
	inputCh := make(chan int)

	// Генерируем числа от 1 до 5
	go func() {
		for i := 1; i <= 5; i++ {
			inputCh <- i
		}
		close(inputCh)
	}()

	// Применяем паттерн Map
	mappedCh := Map(inputCh, double)

	// Выводим результаты
	for result := range mappedCh {
		fmt.Println(result)
	}
}
```
