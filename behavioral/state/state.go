package state

// Интерфейс состояния
type OrderState interface {
	Pay(order *Order, amount float64)
	Ship(order *Order)
	Complete(order *Order)
}
