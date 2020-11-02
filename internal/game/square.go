package game

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// Square ...
type Square struct {
	app.Compo
	position *Position
	// Value    string // b, w, or _

	style string
}

// func newSquare() {}

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

	s.style += " " + color

	checker := Checker{Position: s.position}

	return app.Div().Class(color).Body(
		checker.Render(),
	)
}

// TogglePossibleMoveHighlight ...
func (s *Square) TogglePossibleMoveHighlight() {
	// s.JSSrc.Get("isHighlighted")
}

// OnClick ...
func (s *Square) OnClick(ctx app.Context, e app.Event) {
	ctx.JSSrc.Set("value", s.position.Value)
	s.Update()
}
