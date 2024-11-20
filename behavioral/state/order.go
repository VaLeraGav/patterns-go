package state

// Структура заказа
type Order struct {
	state   OrderState
	payment float64
}

func (o *Order) SetState(state OrderState) {
	o.state = state
}

func (o *Order) Pay(amount float64) {
	o.state.Pay(o, amount)
}

func (o *Order) Ship() {
	o.state.Ship(o)
}

func (o *Order) Complete() {
	o.state.Complete(o)
}
