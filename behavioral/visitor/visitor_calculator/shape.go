package visitor_calculator

type Shape interface {
	getType() string
	accept(Visitor)
}
