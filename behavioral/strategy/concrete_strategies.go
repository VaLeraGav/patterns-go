package strategy

// ------------ ConcreteStrategyA - стратегия без налога ------------
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) CalculatePrice(basePrice float64) float64 {
	return basePrice
}

// ------------ ConcreteStrategyB - стратегия с налогом 10% ------------
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) CalculatePrice(basePrice float64) float64 {
	return basePrice * 1.10
}

// ------------ ConcreteStrategyC - стратегия с налогом 20% ------------
type ConcreteStrategyC struct{}

func (s *ConcreteStrategyC) CalculatePrice(basePrice float64) float64 {
	return basePrice * 1.20
}
