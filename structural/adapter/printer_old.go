package adapter

import (
	"fmt"
)

// OldPrinter - старая структура
type OldPrinter struct{}

// Метод Print для старого принтера
func (p *OldPrinter) PrintOldFormat(message string) {
	fmt.Println("Old Printer:", message)
}
