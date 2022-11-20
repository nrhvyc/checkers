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
