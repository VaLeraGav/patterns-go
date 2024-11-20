package visitor_city

// ----------------- SushiBar -----------------
type SushiBar struct {
}

// Accept implementation.
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi implementation.
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// ----------------- Pizzeria -----------------
type Pizzeria struct {
}

// Accept implementation.
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// BuyPizza implementation.
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// ----------------- BurgerBar -----------------
type BurgerBar struct {
}

// Accept implementation.
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger implementation.
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
