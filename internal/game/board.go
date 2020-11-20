package game

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Board ...
type Board struct {
	app.Compo
	state     string
	Positions [64]*Position
}

// boardState holds global state
var boardState Board

// NewBoard creates a board with pieces in the starting state
func NewBoard() Board {
	b := Board{
		state: "_b_b_b_bb_b_b_b__b_b_b_b________________w_w_w_w__w_w_w_ww_w_w_w_",
		// state: "_b_b_b_b
		// 		b_b_b_b_
		// 		_b_b_b_b
		// 		________
		// 		________
		// 		w_w_w_w_
		// 		_w_w_w_w
		// 		w_w_w_w_",
	}

	b.calculatePositions()

	boardState = b
	return b
}

// Render ...
func (b *Board) Render() app.UI {
	var uiPositions []app.UI

	for i := 0; i < 8; i++ {
		row := []app.UI{} // reset row
		for j := 0; j < 8; j++ {
			row = append(row, b.Positions[8*i+j].Square.Render())
		}
		uiPositions = append(uiPositions, app.Div().Class("row").Body(row...))
	}

	return app.Div().Body(
		app.Div().
			Class("Board").
			// OnClick(b.logDebug).
			Body(uiPositions...),
	)
}

func (b *Board) calculatePositions() {
	// Creating the actual squares of the board
	for i, value := range b.state {
		b.Positions[i] = &Position{Value: i}
		b.Positions[i].Square = &Square{
			position: b.Positions[i],
		}

		// If no checker set the pointer to nil
		if string(value) != "b" && string(value) != "w" {
			b.Positions[i].Checker = nil
			continue
		}

		b.Positions[i].Checker = &Checker{
			Position: b.Positions[i],
			Value:    string(value),
		}

	}
	return
}

// GetPosition retrieves the respective Position from the game state
func (b *Board) GetPosition(val int) *Position {
	return b.Positions[val]
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

		if p.Checker != nil {
			p.Checker.Update()
		}
	}
}
