package adapter

// NewPrinter - новый интерфейс
type NewPrinter interface {
	Print(message string)
}

type Adapter struct {
	oldPrinter *OldPrinter
}

// Метод Print для адаптера
func (a *Adapter) Print(message string) {
	a.oldPrinter.PrintOldFormat(message)
}
