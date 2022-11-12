package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Square ...
type Square struct {
	app.Compo

	Checker    Checker
	hasChecker bool
	Value      string // b, w, or _
	color      string

	style string
}

func (s *Square) OnMount(ctx app.Context) {
	// ctx.GetState("board", &g.Board.Position)
	initGameUI()
}

// Render ...
func (s *Square) Render() app.UI {
	// s.style = color

	// position := boardState.GetPosition(s.position.Value)

	// if !s.hasChecker {
	// 	return app.Div().Class("Square", s.color)
	// }

	// if position.GetValue() == 37 {
	// 	console.Call("log", fmt.Sprintf("position 37: %v\n", position.isHighlighted))
	// }

	// return app.Div().Class(s.style).Text("Square")
	return app.Div().Class("Square", s.color).Body(
		// app.Text(position)
		app.If(s.hasChecker, s.Checker.Render()),
		// app.If(s.Position.isHighlighted, app.Div().Class("possible_move")),
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
