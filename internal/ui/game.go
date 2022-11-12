package ui

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nrhvyc/checkers/internal/game"
)

// Game ...
type Game struct {
	app.Compo
	Board Board
}

// func (g Game) pageLoad() Game {
// 	// Check for ongoing game

// 	// Otherwise load new game
// 	return newGame()
// }

// func NewGame() *Game {
// 	return &Game{Board: Board{}.}
// }

func (g *Game) OnPreRender(ctx app.Context) {
	// fmt.Printf("Before: %+v\n", g)
	// g.Board.State = game.GameState.Board.State
	// fmt.Println()
	// fmt.Println()
	// fmt.Printf("After NewBoard(): %+v\n", g)
	// ctx.SetState("game", g)
}

type GameStateResponse struct {
	Game game.Game `json:"game"`
}

func (g *Game) OnMount(ctx app.Context) {
	initGameUI()
	fmt.Printf("game UI state: %+v\n", g)
	// fmt.Println("testing OnMount")
	// resp, err := http.Get("http://localhost:7790/api/game/state")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Game OnMount() err: %s", err)
	// }
	// gameStateResponse := GameStateResponse{}
	// json.Unmarshal(body, &gameStateResponse)

	// g.Board.State = gameStateResponse.Game.Board.State
	// g.Board.calculatePositions()
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