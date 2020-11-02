package game

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Checker ...
type Checker struct {
	app.Compo
	Position *Position

	Value string // b, w, or _
}

// Render ...
func (c *Checker) Render() app.UI {
	squareClasses := ""
	if string(c.Position.Checker.Value) == "b" {
		squareClasses += "checker-black checker"
	} else if string(c.Position.Checker.Value) == "w" {
		squareClasses += "checker-white checker"
	}

	return app.Div().
		Class(squareClasses)
	// .OnClick()
}

func (c *Checker) onClick(ctx app.Context, e app.Event) {
	for _, move := range c.PossibleMoves() {
		position := boardState.GetPosition(move)

		if !position.HasChecker() {
			position.Square.style = "possible_move"
		}
	}
	// fmt.Println("yep")
	// ctx.JSSrc.Set("value", c.Value)
	boardState.Update()
	return
}

// Move ...
func (c *Checker) Move(val int) {
	c.Position.Value = c.Position.Value + val
	return
}

// NewPosition ...
func (c *Checker) NewPosition(val int) (to Position) {
	to = *c.Position
	to.Value = c.Position.Value + val
	return
}

// PossibleMoves are the positions within the board relative the checker's position
// func (c *Checker) PossibleMoves() (validMoves []Position) {
func (c *Checker) PossibleMoves() (validMoves []int) {
	if c.Position.Value+7 < 63 {
		validMoves = append(validMoves, c.Position.Value+7)
	}
	if c.Position.Value+9 < 63 {
		validMoves = append(validMoves, c.Position.Value+9)
	}
	if c.Position.Value-7 > 0 {
		validMoves = append(validMoves, c.Position.Value-7)
	}
	if c.Position.Value-9 > 0 {
		validMoves = append(validMoves, c.Position.Value-9)
	}
	return
}
