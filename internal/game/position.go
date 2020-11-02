package game

// Position ...
type Position struct {
	Square  Square
	Checker *Checker

	Value int // TODO: make unexported
}

// HasChecker ...
func (p *Position) HasChecker() bool {
	if p.Checker == nil {
		return false
	}
	return true
}

// GetValue - since value will be immutable
func (p *Position) GetValue() int {
	return p.Value
}
