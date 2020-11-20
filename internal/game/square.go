package game

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Square ...
type Square struct {
	app.Compo
	position *Position
	// Value    string // b, w, or _

	style string
}

// Render ...
func (s *Square) Render() app.UI {
	darkMap := [64]bool{
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false}

	color := ""
	if darkMap[s.position.Value] {
		color = "square_dark"
	} else {
		color = "square_light"
	}

	s.style += color

	position := boardState.GetPosition(s.position.Value)

	if position.Checker == nil {
		return app.Div().Class(s.style)
	}

	if position.GetValue() == 37 {
		console.Call("log", fmt.Sprintf("position 37: %v\n", position.isHighlighted))
	}

	return app.Div().Class(s.style).Body(
		position.Checker.Render(),
		app.If(position.isHighlighted, app.Div().Class("possible_move")),
	)
}

// SetStyle ...
func (s *Square) SetStyle(style string) {
	s.style = style
}

// OnClick ...
func (s *Square) OnClick(ctx app.Context, e app.Event) {
	// ctx.JSSrc.Set("value", s.position.Value)
	// ctx.JSSrc.

	s.Update()
}
