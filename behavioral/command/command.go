package command

// Интерфейс команды
type Command interface {
	Execute()
}

type Device interface {
	On()
	Off()
}

// ------------- Команда для включения Device -------------
type DeviceOnCommand struct {
	device Device
}

func NewDeviceOnCommand(device Device) *DeviceOnCommand {
	return &DeviceOnCommand{device: device}
}

func (c *DeviceOnCommand) Execute() {
	c.device.On()
}

// ------------- Команда для выключения Device -------------
type DeviceOffCommand struct {
	device Device
}

func NewDeviceOffCommand(device Device) *DeviceOffCommand {
	return &DeviceOffCommand{device: device}
}

func (c *DeviceOffCommand) Execute() {
	c.device.Off()
}
