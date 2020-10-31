package game

// Position ...
type Position struct {
	Square  Square
	Checker *Checker
	Value   int
}

// HasChecker ...
func (p *Position) HasChecker() bool {
	if p.Checker == nil {
		return false
	}
	return true
}
