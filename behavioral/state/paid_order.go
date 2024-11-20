package state

import "fmt"

type PaidOrderState struct{}

func (s *PaidOrderState) Pay(order *Order, amount float64) {
	fmt.Println("Заказ уже оплачен.")
}

func (s *PaidOrderState) Ship(order *Order) {
	fmt.Println("Заказ отправлен.")
	order.SetState(&ShippedOrderState{})
}

func (s *PaidOrderState) Complete(order *Order) {
	fmt.Println("Нельзя завершить неоплаченный заказ.")
}
