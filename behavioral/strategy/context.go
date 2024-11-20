package strategy

// Context - контекст, использующий стратегию
type Context struct {
	strategy PriceStrategy
}

func (c *Context) SetStrategy(strategy PriceStrategy) {
	c.strategy = strategy
}

func (c *Context) Calculate(basePrice float64) float64 {
	return c.strategy.CalculatePrice(basePrice)
}
