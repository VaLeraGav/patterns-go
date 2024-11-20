package bridge

// Enginer provides engine interface.
type Enginer interface {
	GetSound() string
}

// --------------------  EngineSuzuki --------------------

// EngineSuzuki implements Suzuki engine.
type EngineSuzuki struct {
}

// GetSound returns sound of the engine.
func (e *EngineSuzuki) GetSound() string {
	return "SssuuuuZzzuuuuKkiiiii"
}

// --------------------  EngineHonda --------------------

// EngineHonda implements Honda engine.
type EngineHonda struct {
}

// GetSound returns sound of the engine.
func (e *EngineHonda) GetSound() string {
	return "HhoooNnnnnnnnnDddaaaaaaa"
}

// --------------------  EngineLada --------------------

// EngineLada implements Lada engine.
type EngineLada struct {
}

// GetSound returns sound of the engine.
func (e *EngineLada) GetSound() string {
	return "PhhhhPhhhhPhPhPhPhPh"
}
