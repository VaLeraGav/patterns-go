package state

import "fmt"

type NewOrderState struct{}

func (s *NewOrderState) Pay(order *Order, amount float64) {
	if amount <= 100 {
		fmt.Println("Сумма оплаты должна быть больше 100 рублей.")
		return
	}
	fmt.Println("Заказ оплачен.")
	order.payment = amount
	order.SetState(&PaidOrderState{})
}

func (s *NewOrderState) Ship(order *Order) {
	fmt.Println("Нельзя отправить неоплаченный заказ.")
}

func (s *NewOrderState) Complete(order *Order) {
	fmt.Println("Нельзя завершить неоплаченный заказ.")
}
