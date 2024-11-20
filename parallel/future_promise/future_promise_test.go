package future_promise

import (
	"fmt"
	"testing"
	"time"
)

func TestFanInAndFanOut(t *testing.T) {

	// Задача, которая возвращает ошибку
	// taskWithError := func() (int, error) {
	// 	time.Sleep(1 * time.Second)
	// 	return 0, errors.New("что-то пошло не так")
	// }

	// Запускаем задачу через Promise
	future := Promise(func() (int, error) {
		time.Sleep(1 * time.Second)
		return 0, nil
	})

	fmt.Println("Задача запущена, можно делать что-то еще...")

	// Ожидаем результат
	result := <-future
	if result.err != nil {
		fmt.Println("Ошибка:", result.err)
	} else {
		fmt.Println("Результат:", result.value)
	}

}
