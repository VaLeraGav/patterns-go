package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	// данные, которые будем обрабатывать
	numbers := []int{1, 2, 3, 4, 5, 34, 35, 78}

	// канал для остановки работы горутин
	doneCh := make(chan struct{})
	defer close(doneCh)

	// запускаем генератор, который отправляет числа
	inputCh := Generator(doneCh, numbers)

	// этапы конвейера: сначала Add, потом FilterEven
	addCh := Add(doneCh, inputCh)
	filterEvenCh := FilterEven(doneCh, addCh)

	// выводим результаты
	for res := range filterEvenCh {
		fmt.Println(res)
	}

	// 4
	// 6
	// 36
	// 80
}
