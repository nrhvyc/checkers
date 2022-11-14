package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Game ...
type Game struct {
	app.Compo

	Board              Board
	PossiblePositions  map[int]bool // map[location]bool locations highlighted for a possible move
	LastCheckerClicked int          // location of the last checker clicked
}

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
	// fmt.Printf("RenderState: %+v\n", g)
	return app.Div().Body(
		app.Main().Body(
			g.Board.Render(),
			// app.Text(fmt.Sprintf("%+v", g)),
			// app.Text(g.Board.HasUpdatedPositions),
			// app.Text(g.Board.Positions),
			// GameState,
		),
	)
}
