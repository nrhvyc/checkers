package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nrhvyc/checkers/internal/api"
	"github.com/nrhvyc/checkers/internal/game"
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
	squareClasses := []string{}
	if strings.ToLower(c.Value) == "b" {
		squareClasses = append(squareClasses, []string{"checker-black", "checker"}...)
	} else if c.Value == "w" {
		squareClasses = append(squareClasses, []string{"checker-white", "checker"}...)
	}
	if c.HTMLClasses != "" {
		squareClasses = append(squareClasses, c.HTMLClasses)
	} else if strings.ToLower(c.Value) == "b" {
		squareClasses = append(squareClasses, []string{"checker-black", "checker"}...)
	} else if strings.ToLower(c.Value) == "w" {
		squareClasses = append(squareClasses, []string{"checker-white", "checker"}...)
	}

	if c.Value == "B" || c.Value == "W" {
		squareClasses = append(squareClasses, "king")
	}

	isClientTurn := false
	if strings.ToLower(c.Value) == "b" && UIGameState.PlayerTurn == game.Player1 {
		isClientTurn = true
	} else if strings.ToLower(c.Value) == "w" && UIGameState.PlayerTurn == game.Player2 {
		isClientTurn = true
	}

	if isClientTurn {
		squareClasses = append(squareClasses, " clickable")
		return app.Div().
			OnClick(c.onClick).
			Class("Checker", strings.Join(squareClasses, " "))
	} else {
		return app.Div().
			Class("Checker", strings.Join(squareClasses, " "))
	}
}

func (c *Checker) onClick(ctx app.Context, e app.Event) {
	updatePossibleMoves(c.location)
}

func updatePossibleMoves(checkerPosition int) {
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

	fmt.Printf("updatePossibleMoves() PossibleMoves: %+v\n", possibleMovesResponse)

	UIGameState.LastCheckerClicked = checkerPosition
	UIGameState.PossibleMoves = make(map[int]*game.Move)
	for _, possibleMove := range possibleMovesResponse.Moves {
		possible := possibleMove
		UIGameState.PossibleMoves[possibleMove.Path[len(possibleMove.Path)-1]] = &possible
		fmt.Printf("UIGameState.PossibleMoves: [%d]%+v\n", possibleMove.Path[len(possibleMove.Path)-1], possibleMove)
	}
	for k, v := range UIGameState.PossibleMoves {
		fmt.Printf("UIGameState.PossibleMoves[%d]: %+v\n", k, v)
	}
}
