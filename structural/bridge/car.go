package bridge

// Carer provides car interface.
type Carer interface {
	Rase() string
}

// Car implementation.
type Car struct {
	engine Enginer
}

// NewCar is the Car constructor.
func NewCar(engine Enginer) Carer {
	return &Car{
		engine: engine,
	}
}

// Rase implementation.
func (c *Car) Rase() string {
	return c.engine.GetSound()
}
