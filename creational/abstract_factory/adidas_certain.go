package abstract_factory

// ----------------------------------- AdidasShirt --------------------------------------

type AdidasShirt struct {
	size int
}

func (s *AdidasShirt) setSize(size int) {
	s.size = size
}

func (s *AdidasShirt) GetSize() int {
	return s.size
}

// ----------------------------------- AdidasShoe --------------------------------------

type AdidasShoe struct {
	logo string
	size int
}

func (s *AdidasShoe) setLogo(logo string) {
	s.logo = logo
}

func (s *AdidasShoe) GetLogo() string {
	return s.logo
}

func (s *AdidasShoe) setSize(size int) {
	s.size = size
}

func (s *AdidasShoe) GetSize() int {
	return s.size
}
