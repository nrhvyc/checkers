package ui

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Board ...
type Board struct {
	app.Compo
	State     string
	Positions [64]Position
	// HasUpdatedPositions bool
}

// NewBoard creates a board with pieces in the starting state
// func NewBoard() Board {
// 	b := Board{
// 		State: "_b_b_b_bb_b_b_b__b_b_b_b________________w_w_w_w__w_w_w_ww_w_w_w_",
// 		// state: "_b_b_b_b
// 		// 		b_b_b_b_
// 		// 		_b_b_b_b
// 		// 		________
// 		// 		________
// 		// 		w_w_w_w_
// 		// 		_w_w_w_w
// 		// 		w_w_w_w_",
// 	}

// 	b.calculatePositions()

// 	return b
// }

func (b *Board) OnMount(ctx app.Context) {
	initGameUI()
	// ctx.ObserveState("board").Value(&g.Board)
	// fmt.Println("Board OnMount")
	// game := Game{}
	// ctx.GetState("game", &game)
	// b.State = game.Board.State
	// b.Positions = game.Board.Positions
	// b.HasUpdatedPositions = game.Board.HasUpdatedPositions
	// fmt.Printf("Board OnMount: %+v\n", b)
	// fmt.Printf("OnMount Game State: %+v\n", g)
}

// Render ...
func (b *Board) Render() app.UI {
	b = &UIGameState.Board
	var uiPositions []app.UI

	for i := 0; i < 8; i++ {
		row := []app.UI{} // reset row
		for j := 0; j < 8; j++ {
			// s := Square{}
			// row = append(row, s.Render())
			row = append(row, b.Positions[8*i+j].Render())
			// row = append(row, app.Div().Text("square"))
		}
		uiPositions = append(uiPositions, app.Div().Class("row").Body(row...))
	}

	return app.Div().Body(
		app.Div().
			// Text(b.Positions),
			Class("Board").
			// OnClick(b.logDebug).
			Body(uiPositions...),
	)
}

func (b *Board) calculatePositions() {
	darkMap := [64]bool{
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false}

	// b.HasUpdatedPositions = true

	// Creating the actual squares of the board
	for i, value := range strings.Split(b.State, "") {
		b.Positions[i] = Position{Value: i}
		b.Positions[i].Square = Square{
			Value: value,
		}

		if darkMap[i] {
			b.Positions[i].Square.color = "square_dark"
		} else {
			b.Positions[i].Square.color = "square_light"
		}

		// If no checker set the pointer to nil
		if string(value) == "b" || string(value) == "w" {
			b.Positions[i].Square.hasChecker = true
		} else {
			continue
		}

		b.Positions[i].Square.Checker = Checker{
			Value: string(value),
		}
	}
}

// GetPosition retrieves the respective Position from the game state
func (b *Board) GetPosition(val int) Position {
	return b.Positions[val]
}

// UpdatePosition ...
func (b *Board) UpdatePosition(val int, position Position) {
	b.Positions[val] = position
	b.Update()
}

// ClearHighlighted sets isHighlighted = false for all positions
func (b *Board) ClearHighlighted() {
	for _, p := range b.Positions {
		p.isHighlighted = false
	}
}

// UpdateAll the UI get rendered
func (b *Board) UpdateAll() {
	b.Update()
	for _, p := range b.Positions {
		p.Square.Update()

		// if p.Checker != nil {
		// 	p.Checker.Update()
		// }
	}
}
