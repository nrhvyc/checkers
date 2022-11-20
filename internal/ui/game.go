package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nrhvyc/checkers/internal/game"
)

// Game ...
type Game struct {
	app.Compo

	Board Board
	// PossibleMoves  map[int]bool // map[location]bool locations highlighted for a possible move
	PossibleMoves map[int]*game.Move // map[location]bool locations highlighted for a possible move
	// Moves  map[int]bool // map[location]bool locations highlighted for a possible move
	LastCheckerClicked int  // location of the last checker clicked
	PlayerTurn         bool // false = black's turn; true = white's turn
	Winner             game.Winner
}

var winnerMessage = [3]string{"", "Player 1 Has Won!", "Player 2 Has Won!"}

// func (g *Game) OnPreRender(ctx app.Context) {}

// type GameStateResponse struct {
// 	GameState string `json:"gameState"`
// }

func (g *Game) OnMount(ctx app.Context) {
	initGameUI()
}

// Render ...
func (g *Game) Render() app.UI {
	g = &UIGameState
	winnerMsg := winnerMessage[UIGameState.Winner]

	// fmt.Printf("RenderState: %+v\n", g)

	return app.Div().Class("grid-container").Body(
		app.Div().Body(
			g.Board.Render(),
			// app.Text(fmt.Sprintf("%+v", g)),
			// app.Text(g.Board.HasUpdatedPositions),
			// app.Text(g.Board.Positions),
			// GameState,
		),
		app.If(winnerMsg != "",
			app.Div().Class("winner").Body(
				app.Div().Class("winner-text").Body(
					app.Text(winnerMsg),
				),
			),
		),
	)
}
