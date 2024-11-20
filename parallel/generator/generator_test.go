package generator

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	// Данные, которые будут отправляться в канал
	items := []int{10, 20, 30, 40, 50}

	// Получаем канал с данными из генератора
	dataChannel := Generator(items)

	for item := range dataChannel {
		fmt.Println(item)
	}
}
