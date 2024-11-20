package mediator_train

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
