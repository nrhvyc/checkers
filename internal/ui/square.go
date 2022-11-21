package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nrhvyc/checkers/internal/api"
	"github.com/nrhvyc/checkers/internal/game"
)

// Square ...
type Square struct {
	app.Compo

	Checker    Checker
	hasChecker bool
	Value      string // b, w, or _
	color      string

	location int // index in UIGameState.Game.Board.Position[location]

	style string
}

func (s *Square) OnMount(ctx app.Context) {
	if UIGameState.GameMode != NewGame {
		initGameUI()
	}
}

// Render ...
func (s *Square) Render() app.UI {
	return app.Div().Class("Square", s.color).Body(
		app.If(s.hasChecker, s.Checker.Render()),
		app.If(
			UIGameState.PossibleMoves[s.location] != nil,
			app.Div().Class("possible_move").OnClick(s.onPossibleMoveClick)),
	)
}

// SetStyle ...
func (s *Square) SetStyle(style string) {
	s.style = style
}

// OnClick ...
func (s *Square) OnClick(ctx app.Context, e app.Event) {
	s.Update()
}

func (s *Square) onPossibleMoveClick(ctx app.Context, e app.Event) {
	makeMove(UIGameState.LastCheckerClicked, s.location)
}

func makeMove(from, to int) {
	fmt.Printf("square makeMove() from: %d to: %d", from, to)
	checkerMoveRequest := api.CheckerMoveRequest{
		Move: *UIGameState.PossibleMoves[to],
	}
	req, err := json.Marshal(checkerMoveRequest)
	if err != nil {
		fmt.Printf("error marshalling MoveRequest err: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:7790/api/checker/move",
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

	checkerMoveResponse := api.CheckerMoveResponse{}
	json.Unmarshal(body, &checkerMoveResponse)

	fmt.Printf("makeMove() PossibleMoves: %+v", checkerMoveResponse)

	UIGameState.PossibleMoves = make(map[int]*game.Move)
	for _, possibleMove := range checkerMoveResponse.FollowUpMoves {
		possible := possibleMove
		UIGameState.PossibleMoves[possibleMove.Path[len(possibleMove.Path)-1]] = &possible
		fmt.Printf("UIGameState.PossibleMoves: [%d]%+v\n", possibleMove.Path[len(possibleMove.Path)-1], possibleMove)
	}
	UIGameState.Board.State = checkerMoveResponse.GameState
	UIGameState.PlayerTurn = checkerMoveResponse.PlayerTurn
	UIGameState.Winner = checkerMoveResponse.Winner
	UIGameState.Board.calculatePositions()
}
