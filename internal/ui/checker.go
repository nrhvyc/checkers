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

// Checker ...
type Checker struct {
	app.Compo

	location int // index in UIGameState.Game.Board.Position[location]

	Value string // b, w, or _

	HTMLClasses string
}

func (c *Checker) OnMount(ctx app.Context) {
	initGameUI()
}

// Render ...
func (c *Checker) Render() app.UI {
	squareClasses := ""
	if c.Value == "b" {
		squareClasses += "checker-black checker"
	} else if c.Value == "w" {
		squareClasses += "checker-white checker"
	}
	if c.HTMLClasses != "" {
		squareClasses = c.HTMLClasses
	} else if c.Value == "b" {
		squareClasses += "checker-black checker"
	} else if c.Value == "w" {
		squareClasses += "checker-white checker"
	}

	// if position.Value == 37 {
	// 	console.Call("log", fmt.Sprintf("rendering Checker @ Position: %v\n", position.Value))
	// }

	return app.Div().
		OnClick(c.onClick).
		Class("Checker", squareClasses)
}

func (c *Checker) onClick(ctx app.Context, e app.Event) {
	updatePossiblePositions(c.location)
}

func updatePossiblePositions(checkerPosition int) {
	possibleMovesRequest := api.PossibleMovesRequest{CheckerPosition: checkerPosition}
	req, err := json.Marshal(possibleMovesRequest)
	if err != nil {
		fmt.Printf("error marshalling PossibleMovesRequest err: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:7790/api/checker/possible-moves",
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

	possibleMovesResponse := api.PossibleMovesResponse{}
	json.Unmarshal(body, &possibleMovesResponse)

	fmt.Printf("PossibleMoves: %+v", possibleMovesResponse)

	UIGameState.LastCheckerClicked = checkerPosition
	UIGameState.PossiblePositions = make(map[int]bool)
	for _, possibleMove := range possibleMovesResponse.Moves {
		UIGameState.PossiblePositions[possibleMove.ToLocation] = true
	}
}

// Move ...
// func (c *Checker) Move(val int) {
// 	c.Position.Value = c.Position.Value + val
// }

// // NewPosition ...
// func (c *Checker) NewPosition(val int) (to Position) {
// 	to = *c.Position
// 	to.Value = c.Position.Value + val
// 	return
// }

// // PossibleMoves are the positions within the board relative the checker's position
// // func (c *Checker) PossibleMoves() (validMoves []Position) {
// func (c *Checker) PossibleMoves() (validMoves []int) {
// 	if c.Position.Value+7 < 63 {
// 		validMoves = append(validMoves, c.Position.Value+7)
// 	}
// 	if c.Position.Value+9 < 63 {
// 		validMoves = append(validMoves, c.Position.Value+9)
// 	}
// 	if c.Position.Value-7 > 0 {
// 		validMoves = append(validMoves, c.Position.Value-7)
// 	}
// 	if c.Position.Value-9 > 0 {
// 		validMoves = append(validMoves, c.Position.Value-9)
// 	}
// 	return
// }
