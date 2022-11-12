package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Checker ...
type Checker struct {
	app.Compo
	// Position *Position

	Value string // b, w, or _

	HTMLClasses string
}

func (c *Checker) OnMount(ctx app.Context) {
	initGameUI()
}

// Render ...
func (c *Checker) Render() app.UI {
	squareClasses := ""
	if c.Value == "b" {
		squareClasses += "checker-black checker"
	} else if c.Value == "w" {
		squareClasses += "checker-white checker"
	}
	if c.HTMLClasses != "" {
		squareClasses = c.HTMLClasses
	} else if c.Value == "b" {
		squareClasses += "checker-black checker"
	} else if c.Value == "w" {
		squareClasses += "checker-white checker"
	}

	// if position.Value == 37 {
	// 	console.Call("log", fmt.Sprintf("rendering Checker @ Position: %v\n", position.Value))
	// }

	return app.Div().
		// OnClick(c.onClick).
		Class("Checker", squareClasses)
}

// func (c *Checker) onClick(ctx app.Context, e app.Event) {
// 	// boardState.ClearHighlighted() // Clear existing highlights

// 	// for _, move := range c.PossibleMoves() {
// 	// 	position := boardState.GetPosition(move)

// 	// 	if !position.HasChecker() {
// 	// 		position.ToggleHighlight()

// 	// 		console.Call("log", fmt.Sprintf("position: %v\n", position))
// 	// 		console.Call("log", fmt.Sprintf("HasChecker: %v\n", position.HasChecker()))
// 	// 		console.Call("log",
// 	// 			fmt.Sprintf("isHighlighted: %v\n", position.isHighlighted))
// 	// 	}
// 	// 	position.Square.Update()
// 	// }

// 	// console.Call("log", fmt.Sprintf("c.Value: %v\n", c.Value))

// 	boardState := Board{}
// 	ctx.GetState("board-state", &boardState)

// 	position := boardState.GetPosition(c.Position.Value)
// 	position.Checker.Value = "b"

// 	boardState.UpdatePosition(c.Position.Value, position)

// 	// c.Position = position

// 	// console.Call("log", fmt.Sprintf("position: %v\n", boardState.Positions[37]))
// 	// console.Call("log", fmt.Sprintf("position: %v\n", boardState.Positions[c.Position.Value]))
// 	// boardState.Update()
// 	// boardState.UpdateAll()
// 	c.HTMLClasses = "checker-black checker"
// 	c.Update()
// 	// position.Checker.Update()
// 	return
// }

// // Move ...
// func (c *Checker) Move(val int) {
// 	c.Position.Value = c.Position.Value + val
// }

// // NewPosition ...
// func (c *Checker) NewPosition(val int) (to Position) {
// 	to = *c.Position
// 	to.Value = c.Position.Value + val
// 	return
// }

// // PossibleMoves are the positions within the board relative the checker's position
// // func (c *Checker) PossibleMoves() (validMoves []Position) {
// func (c *Checker) PossibleMoves() (validMoves []int) {
// 	if c.Position.Value+7 < 63 {
// 		validMoves = append(validMoves, c.Position.Value+7)
// 	}
// 	if c.Position.Value+9 < 63 {
// 		validMoves = append(validMoves, c.Position.Value+9)
// 	}
// 	if c.Position.Value-7 > 0 {
// 		validMoves = append(validMoves, c.Position.Value-7)
// 	}
// 	if c.Position.Value-9 > 0 {
// 		validMoves = append(validMoves, c.Position.Value-9)
// 	}
// 	return
// }
