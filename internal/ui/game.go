package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nrhvyc/checkers/internal/api"
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

var winnerMessage = [3]string{"", "Player 1 Won!", "Player 2 Won!"}

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
		app.Raw(`<link href="https://fonts.cdnfonts.com/css/neue-haas-grotesk-text-pro" rel="stylesheet">`),
		app.Div().Body(
			g.Board.Render(),
		),
		app.Div().Class("menu").Body(
			app.If(winnerMsg != "",
				app.Div().Class("winner menu-center").Body(
					app.Div().Class("winner-text").Body(
						app.Text(winnerMsg),
					),
					app.Div().Class("play-again-container").Body(
						app.Div().Class("play-again btn-hover").Body(
							app.Text("Play Again"),
						).OnClick(g.onClickPlayAgain),
					),
				),
			),
		),
	)
}

func (g *Game) onClickPlayAgain(ctx app.Context, e app.Event) {
	resp, err := http.Get("http://localhost:7790/api/game/play-again")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("onClickPlayAgain() err: %s", err)
	}
	gameStateResponse := api.GameStateResponse{}
	json.Unmarshal(body, &gameStateResponse)

	UIGameState.Board.State = gameStateResponse.GameState
	UIGameState.Winner = gameStateResponse.Winner
	UIGameState.PlayerTurn = gameStateResponse.PlayerTurn
	fmt.Printf("Current Board State: %s\n", UIGameState.Board.State)
	UIGameState.PossibleMoves = make(map[int]*game.Move)
	UIGameState.Board.calculatePositions()
}
