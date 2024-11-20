package state

import "fmt"

type ShippedOrderState struct{}

func (s *ShippedOrderState) Pay(order *Order, amount float64) {
	fmt.Println("Заказ уже оплачен.")
}

func (s *ShippedOrderState) Ship(order *Order) {
	fmt.Println("Заказ уже отправлен.")
}

func (s *ShippedOrderState) Complete(order *Order) {
	fmt.Println("Заказ завершен.")
	order.SetState(&CompletedOrderState{})
}
