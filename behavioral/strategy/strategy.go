package strategy

// PriceStrategy интерфейс для различных стратегий расчета цены
type PriceStrategy interface {
	CalculatePrice(basePrice float64) float64
}
