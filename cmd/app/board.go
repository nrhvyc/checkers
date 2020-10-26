package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type Board struct {
	app.Compo
	state string
}

func newBoard() Board {
	return Board{
		state: "_b_b_b_bb_b_b_b__b_b_b_b________________w_w_w_w__w_w_w_ww_w_w_w_",
	}
}

// Render ...
func (b *Board) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("Board").Body(
			b.calculateBoard()...,
		),
	)
}

func (b *Board) calculateBoard() (boardUI []app.UI) {
	// board state will be recieved as single string
	board := b.state

	row := []app.UI{}

	// Creating the actual squares of the board
	for i, square := range board {
		if i%8 == 0 {
			boardUI = append(boardUI, app.Div().Class("row").Body(row...))
			row = []app.UI{} // reset row
		}

		row = append(row, SquareDiv(i, string(square)))

		if len(board)-1 == i {
			boardUI = append(boardUI, app.Div().Class("row").Body(row...))
		}
	}

	return boardUI
}

// SquareDiv ...
func SquareDiv(position int, value string) app.HTMLDiv {
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
	if darkMap[position] {
		color = "square_dark"
	} else {
		color = "square_light"
	}

	squareClasses := color + " "

	if string(value) == "b" {
		squareClasses += "checker-black checker"
	} else if string(value) == "w" {
		squareClasses += "checker-white checker"
	}

	return app.Div().Class(color).Body(
		app.Div().Class(squareClasses),
	)
}
