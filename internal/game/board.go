package game

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Board ...
type Board struct {
	app.Compo
	state     string
	Positions [64]Position
}

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

	// b.calculateBoard()

	return b
}

// Render ...
func (b *Board) Render() app.UI {
	b.calculateBoard()
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
			Body(uiPositions...),
	)
}

func (b *Board) calculateBoard() {
	// board := b.state
	// board state will be recieved as single string

	// row := []app.UI{}

	// Creating the actual squares of the board
	for i, value := range b.state {
		// if i%8 == 0 {
		// 	boardUI = append(boardUI, app.Div().Class("row").Body(row...))
		// 	row = []app.UI{} // reset row
		// }

		b.Positions[i] = Position{Value: i}
		b.Positions[i].Square = Square{
			Position: &b.Positions[i],
			// Value:    string(value),
		}
		b.Positions[i].Checker = &Checker{
			Position: &b.Positions[i],
			Value:    string(value),
		}

		// row = append(row, square.Render())

		// if len(board)-1 == i {
		// 	boardUI = append(boardUI, app.Div().Class("row").Body(row...))
		// }
	}

	return
}

// func GetPosition(value int) *Position {
// 	return &board.Positions[value]
// }
