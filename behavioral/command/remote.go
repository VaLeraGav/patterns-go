package command

import "fmt"

// Пульт дистанционного управления
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) ExecuteCommand() {
	if r.command != nil {
		r.command.Execute()
	} else {
		fmt.Println("Команда не установлена.")
	}
}
