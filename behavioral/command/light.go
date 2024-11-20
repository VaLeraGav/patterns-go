package command

import "fmt"

// Получатель
type TV struct {
	isRunning bool
}

func (l *TV) On() {
	l.isRunning = true
	fmt.Println("Device включен.")
}

func (l *TV) Off() {
	l.isRunning = false
	fmt.Println("Device выключен.")
}
