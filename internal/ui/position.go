package ui

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Position ...
type Position struct {
	Square  Square
	Checker Checker

	isHighlighted bool // Where a move could be made

	Value int // TODO: make unexported
}

func (p *Position) Render() app.UI {
	return app.Div().Class("Position").Body(
		p.Square.Render(),
	)
}

// HasChecker ...
// func (p *Position) HasChecker() bool {
// 	return p.Checker != nil
// }

// GetValue - since value will be immutable
// func (p *Position) GetValue() int {
// 	return p.Value
// }

// // ToggleHighlight is for toggling highlighting for a move
// func (p *Position) ToggleHighlight() {
// 	p.isHighlighted = !p.isHighlighted
// }
