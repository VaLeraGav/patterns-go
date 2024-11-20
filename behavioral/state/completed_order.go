package state

import "fmt"

type CompletedOrderState struct{}

func (s *CompletedOrderState) Pay(order *Order, amount float64) {
	fmt.Println("Заказ уже завершен.")
}

func (s *CompletedOrderState) Ship(order *Order) {
	fmt.Println("Заказ уже завершен.")
}

func (s *CompletedOrderState) Complete(order *Order) {
	fmt.Println("Заказ уже завершен.")
}
