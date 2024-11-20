package mediator_train

type Train interface {
	arrive()
	depart()
	permitArrival()
}
