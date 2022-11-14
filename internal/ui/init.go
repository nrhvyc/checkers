package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nrhvyc/checkers/internal/api"
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

	UIGameState.Board.State = gameStateResponse.GameState
	UIGameState.PossiblePositions = make(map[int]bool)
	UIGameState.Board.calculatePositions()
}
