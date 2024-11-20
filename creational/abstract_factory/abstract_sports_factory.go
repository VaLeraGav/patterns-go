package abstract_factory

type AbstractSportsFactory interface {
	MakeShoe() AbstractShoe
	MakeShirt() AbstractShirt
}

func CreateAdidasFactory() AbstractSportsFactory {
	return &Adidas{}
}
func CreateNikeFactory() AbstractSportsFactory {
	return &Nike{}
}
