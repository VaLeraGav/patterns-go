package prototype

type Shape interface {
	Clone() Shape
	Draw() string
}
