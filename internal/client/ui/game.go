package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	clientAPI "github.com/nrhvyc/checkers/internal/api/client"
	"github.com/nrhvyc/checkers/internal/client"
	serverAPI "github.com/nrhvyc/checkers/internal/server/api"
	"github.com/nrhvyc/checkers/internal/server/game"
)

// Game ...
type Game struct {
	app.Compo

	Board              Board
	PossibleMoves      map[int]*game.Move // map[location]bool locations highlighted for a possible move
	LastCheckerClicked int                // location of the last checker clicked
	PlayerTurn         game.PlayerTurn    // false = black's turn; true = white's turn

	TurnCount int

	Winner       game.Winner
	GameMode     game.GameMode
	Players      [2]game.Player
	ClientPlayer game.PlayerTurn // which player this client is. Only used during single player for now
}

var winnerMessage = [3]string{"", "Player 1 Won!", "Player 2 Won!"}

// func (g *Game) OnPreRender(ctx app.Context) {}

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
			app.If(UIGameState.GameMode == game.NewGameMode,
				app.H1().Class("menu-title").Body(
					app.Text("Checkers"),
				),
				app.Div().Class("new-game-text").Body(
					app.Text("New Game Selection"),
				),
				app.Div().Class("menu-center").Body(
					app.Div().Class("new-game-container").Body(
						app.Div().Class("new-game btn-hover").Body(
							app.Text("Two Player"),
						).OnClick(g.onClickTwoPlayer),
					),
					app.Div().Class("new-game-container").Body(
						app.Div().Class("new-game btn-hover").Body(
							app.Text("Single Player"),
						).OnClick(g.onClickSinglePlayer),
					),
				),
			),
			app.If(UIGameState.GameMode != game.NewGameMode,
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
		),
	)
}

func (g *Game) onClickPlayAgain(ctx app.Context, e app.Event) {
	resp, err := http.Get("http://localhost:7790/server/api/game/play-again")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("onClickPlayAgain() err: %s", err)
	}
	gameStateResponse := serverAPI.GameStateResponse{}
	json.Unmarshal(body, &gameStateResponse)

	UIGameState.Board.State = gameStateResponse.GameState
	UIGameState.Winner = gameStateResponse.Winner
	UIGameState.PlayerTurn = gameStateResponse.PlayerTurn
	fmt.Printf("Current Board State: %s\n", UIGameState.Board.State)
	UIGameState.PossibleMoves = make(map[int]*game.Move)
	UIGameState.Board.calculatePositions()
}

func (g *Game) onClickTwoPlayer(ctx app.Context, e app.Event) {
	g.newGame(ctx, e, game.TwoPlayer)
}

func (g *Game) onClickSinglePlayer(ctx app.Context, e app.Event) {
	g.newGame(ctx, e, game.SinglePlayer)
}

func (g *Game) newGame(ctx app.Context, e app.Event, gameMode game.GameMode) {
	var newGameResponse serverAPI.NewGameResponse

	if gameMode == game.TwoPlayer {
		go clientAPI.ClientServer()
		client.RequestTwoPlayerMatch()
	} else {
		newLocalGame(ctx, gameMode)
	}

	UIGameState.GameMode = gameMode
	UIGameState.Board.State = newGameResponse.GameState
	UIGameState.Winner = game.NoWinner
	UIGameState.PlayerTurn = newGameResponse.PlayerTurn
	fmt.Printf("Current Board State: %s\n", UIGameState.Board.State)
	UIGameState.PossibleMoves = make(map[int]*game.Move)
	UIGameState.Board.calculatePositions()
}

/*
newLocalGame requests a new game with only one browser client connected with the
game running on the server
*/
func newLocalGame(ctx app.Context, gameMode game.GameMode) {
	newGameRequest := serverAPI.NewGameRequest{GameMode: game.SinglePlayer}
	req, err := json.Marshal(newGameRequest)
	if err != nil {
		fmt.Printf("error marshalling PossibleMovesRequest err: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:7790/server/api/game/new",
		bytes.NewBuffer(req))
	if err != nil {
		fmt.Printf("error creating request err: %s\n", err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("client.Do err: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Game OnMount() err: %s", err)
	}

	var newGameResponse serverAPI.NewGameResponse
	json.Unmarshal(body, &newGameResponse)
}
