package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nrhvyc/checkers/internal/api"
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
	// ctx.GetState("board", &g.Board.Position)
	initGameUI()
}

// Render ...
func (s *Square) Render() app.UI {
	// s.style = color

	// position := boardState.GetPosition(s.position.Value)

	// if !s.hasChecker {
	// 	return app.Div().Class("Square", s.color)
	// }

	// if position.GetValue() == 37 {
	// 	console.Call("log", fmt.Sprintf("position 37: %v\n", position.isHighlighted))
	// }

	// return app.Div().Class(s.style).Text("Square")
	return app.Div().Class("Square", s.color).Body(
		// app.Text(position)
		app.If(s.hasChecker, s.Checker.Render()),
		app.If(
			UIGameState.PossiblePositions[s.location],
			app.Div().Class("possible_move").OnClick(s.onPossibleMoveClick)),
	)
}

// SetStyle ...
func (s *Square) SetStyle(style string) {
	s.style = style
}

// OnClick ...
func (s *Square) OnClick(ctx app.Context, e app.Event) {
	// ctx.JSSrc.Set("value", s.position.Value)
	// ctx.JSSrc.

	s.Update()
}

func (s *Square) onPossibleMoveClick(ctx app.Context, e app.Event) {
	makeMove(UIGameState.LastCheckerClicked, s.location)
}

func makeMove(from, to int) {
	checkerMoveRequest := api.CheckerMoveRequest{
		From: from,
		To:   to,
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

	fmt.Printf("PossibleMoves: %+v", checkerMoveResponse)

	UIGameState.PossiblePositions = make(map[int]bool)
	UIGameState.Board.State = checkerMoveResponse.GameState
	UIGameState.Board.calculatePositions()
}
