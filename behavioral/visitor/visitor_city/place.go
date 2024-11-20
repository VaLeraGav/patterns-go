package visitor_city

// Visitor provides a visitor interface.
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Place provides an interface for place that the visitor should visit.
type Place interface {
	Accept(v Visitor) string
}
