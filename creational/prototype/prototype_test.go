package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {

	// Создаем прототипы
	circlePrototype := &Circle{radius: 5.0}
	squarePrototype := &Square{side: 4.0}

	// Клонируем прототипы
	circleClone := circlePrototype.Clone()
	squareClone := squarePrototype.Clone()

	// Изменяем радиус клона круга
	circleClone.(*Circle).radius = 10.0

	// Выводим результаты
	fmt.Println(circlePrototype.Draw()) // Circle with radius: 5.00
	fmt.Println(circleClone.Draw())     // Circle with radius: 10.00
	fmt.Println(squarePrototype.Draw()) // Square with side: 4.00
	fmt.Println(squareClone.Draw())     // Square with side: 4.00
}
