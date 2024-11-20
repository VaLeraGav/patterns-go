package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	// Создаем объект света
	Device := &TV{}

	// Создаем команды
	deviceOnCommand := NewDeviceOnCommand(Device)
	deviceOffCommand := NewDeviceOffCommand(Device)

	// Создаем пульт дистанционного управления
	remote := &RemoteControl{}

	// Устанавливаем команды
	remote.SetCommand(deviceOnCommand)
	remote.ExecuteCommand() // Device включен.

	remote.SetCommand(deviceOffCommand)
	remote.ExecuteCommand() // Device выключен.
}
