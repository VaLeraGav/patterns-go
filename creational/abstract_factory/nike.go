package abstract_factory

type Nike struct {
}

func (n *Nike) MakeShoe() AbstractShoe {
	return &NikeShoe{
		logo: "nike",
		size: 14,
	}
}

func (n *Nike) MakeShirt() AbstractShirt {
	return &NikeShirt{
		size: 14,
	}
}
