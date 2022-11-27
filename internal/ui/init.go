package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nrhvyc/checkers/internal/api"
	"github.com/nrhvyc/checkers/internal/game"
)

// var console = js.Global().Get("console")

// UIGameState holds global state
var UIGameState Game

var hasLoaded bool

func initGameUI() {
	if hasLoaded {
		return
	} else {
		hasLoaded = true
	}
	resp, err := http.Get("http://localhost:7790/api/game/state")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Game OnMount() err: %s", err)
	}
	gameStateResponse := api.GameStateResponse{}
	json.Unmarshal(body, &gameStateResponse)

	fmt.Printf("gameStateResponse.GameMode: %d\n", gameStateResponse.GameMode)

	UIGameState.GameMode = gameStateResponse.GameMode

	if UIGameState.GameMode == game.NewGameMode {
		return
	}

	UIGameState.Board.State = gameStateResponse.GameState
	UIGameState.Winner = gameStateResponse.Winner
	fmt.Printf("Current Board State: %s\n", UIGameState.Board.State)
	UIGameState.PossibleMoves = make(map[int]*game.Move)

	UIGameState.Board.calculatePositions()
}
