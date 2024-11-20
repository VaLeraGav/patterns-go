package map_filter

import (
	"fmt"
	"testing"
)

// Функция для удвоения числа
func double(x int) int {
	return x * 2
}

// Условие для фильтрации: оставляем только четные числа
func isEven(x int) bool {
	return x%2 == 0
}

func TestMapFilter(t *testing.T) {
	numbers := []int{1, 2, 4, 5, 34, 35, 78}

	doneCh := make(chan struct{})

	gen := Generator(doneCh, numbers)
	mapped := Map(doneCh, gen, double)
	resultCh := Filter(doneCh, mapped, isEven)

	fmt.Println("Результаты после применения функции удвоения:")

	for num := range resultCh {
		fmt.Println(num)
	}

	close(doneCh)
	// Проверка, что doneCh закрыт
	select {
	case <-doneCh:
		fmt.Println("Горутины завершены.")
	default:
		fmt.Println("doneCh не закрыт.")
	}

	// 	Результаты после применения функции удвоения:
	// 2
	// 4
	// 8
	// 10
	// 68
	// 70
	// inputCh закрыт, выходим из Map
	// inputCh закрыт, выходим из Filter
	// 156
	// Горутины завершены.
}
