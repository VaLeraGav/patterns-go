# Стратегия (Strategy)

Паттерн `Strategy` относится к поведенческим паттернам уровня объекта.

Паттерн `Strategy` определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный класс и делает их подменяемыми. Паттерн `Strategy` позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.

Требуется для реализации:

- Класс Context, представляющий собой контекст выполнения той или иной стратегии

```go
type Context struct {
	strategy PriceStrategy
}

func (c *Context) SetStrategy(strategy PriceStrategy) {
	c.strategy = strategy
}

func (c *Context) Calculate(basePrice float64) float64 {
	return c.strategy.CalculatePrice(basePrice)
}
```

- Абстрактный класс `Strategy`, определяющий интерфейс различных стратегий

```go
type PriceStrategy interface {
	CalculatePrice(basePrice float64) float64
}
```

- Класс `ConcreteStrategyB`, `ConcreteStrategyA` реализует одно из стратегий представляющую собой алгоритмы, направленные на достижение определенной цели.

```go
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) CalculatePrice(basePrice float64) float64 {
	return basePrice * 1.10
}

type ConcreteStrategyC struct{}

func (s *ConcreteStrategyC) CalculatePrice(basePrice float64) float64 {
	return basePrice * 1.20
}
```
