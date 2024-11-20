package state

import (
	"testing"
)

func TestState(t *testing.T) {
	order := &Order{state: &NewOrderState{}}

	order.Pay(50)
	order.Pay(100)
	order.Ship()
	order.Complete()

	// Нельзя оформить заказ, так как сумма меньше 100.
	// Заказ оплачен.
	// Заказ отправлен.
	// Заказ завершен.

	orderNew := &Order{state: &NewOrderState{}}
	orderNew.Pay(100)
	orderNew.Complete()

	// Заказ оплачен.
	// Нельзя завершить неоплаченный заказ.
}
