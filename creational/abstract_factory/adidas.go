package abstract_factory

type Adidas struct {
}

func (a *Adidas) MakeShoe() AbstractShoe {
	return &AdidasShoe{
		logo: "adidas",
		size: 16,
	}
}

func (a *Adidas) MakeShirt() AbstractShirt {
	return &AdidasShirt{
		size: 16,
	}
}
