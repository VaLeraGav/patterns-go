package abstract_factory

// ----------------------------------- NikeShirt --------------------------------------

type NikeShirt struct {
	size int
}

func (s *NikeShirt) setSize(size int) {
	s.size = size
}

func (s *NikeShirt) GetSize() int {
	return s.size
}

// ----------------------------------- NikeShoe --------------------------------------

type NikeShoe struct {
	logo string
	size int
}

func (s *NikeShoe) setLogo(logo string) {
	s.logo = logo
}

func (s *NikeShoe) GetLogo() string {
	return s.logo
}

func (s *NikeShoe) setSize(size int) {
	s.size = size
}

func (s *NikeShoe) GetSize() int {
	return s.size
}
