package chain_of_responsibility

import (
	"fmt"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {

	// Создаем обработчиков
	low := &LowPriorityHandler{}
	medium := &MediumPriorityHandler{}
	high := &HighPriorityHandler{}

	// Устанавливаем цепочку
	low.SetNext(medium).SetNext(high)

	// Тестируем цепочку
	requests := []string{"low", "medium", "high", "unknown"}

	for _, req := range requests {
		fmt.Printf("Client: What do you want to do with '%s'?\n", req)
		response := low.Handle(req)
		if response != "" {
			fmt.Println(response)
		} else {
			fmt.Println("No handler found for this request.")
		}
	}

	// Client: What do you want to do with 'low'?
	// LowPriorityHandler: Handling low priority request.
	//
	// Client: What do you want to do with 'medium'?
	// MediumPriorityHandler: Handling medium priority request.
	//
	// Client: What do you want to do with 'high'?
	// HighPriorityHandler: Handling high priority request.
	//
	// Client: What do you want to do with 'unknown'?
	// No handler found for this request.
}
