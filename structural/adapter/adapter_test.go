package adapter

import (
	"testing"
)

func TestHouseBuilder(t *testing.T) {

	// Создаем старый принтер
	oldPrinter := &OldPrinter{}

	// Создаем адаптер для старого принтера
	adapter := &Adapter{oldPrinter: oldPrinter}

	// Используем адаптер как новый принтер
	PrintMessage(adapter, "Hello, Adapter Pattern!")
}

// Функция, принимающая новый принтер
func PrintMessage(printer NewPrinter, message string) {
	printer.Print(message)
}
