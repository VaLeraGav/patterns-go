package mediator_train

import (
	"testing"
)

// Два поезда никогда не конфликтуют между собой за наличие свободной платформы,
// вместо этого менеджер станции stationManager выступает посредником, который
// оглашает платформу свободной лишь для одного поезда, а остальные содержит в очереди.

func TestMediator(t *testing.T) {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()

	// 	PassengerTrain: Arrived
	// FreightTrain: Arrival blocked, waiting
	// PassengerTrain: Leaving
	// FreightTrain: Arrival permitted
	// FreightTrain: Arrived
}
