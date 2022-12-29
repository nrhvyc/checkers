package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	serverAPI "github.com/nrhvyc/checkers/internal/api/server"
	"github.com/nrhvyc/checkers/internal/game"
)

// UIGameState holds global client state
var UIGameState Game

var hasLoaded bool

func initGameUI() {
	if hasLoaded {
		return
	} else {
		hasLoaded = true
	}
	resp, err := http.Get("http://localhost:7790/server/api/game/state")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Game OnMount() err: %s", err)
	}
	gameStateResponse := serverAPI.GameStateResponse{}
	json.Unmarshal(body, &gameStateResponse)

	fmt.Printf("gameStateResponse.GameMode: %d\n", gameStateResponse.GameMode)

	UIGameState.GameMode = gameStateResponse.GameMode

	if UIGameState.GameMode == game.NewGameMode {
		return
	}

	if UIGameState.GameMode == game.SinglePlayer {
		for pos, player := range gameStateResponse.Players {
			if player.Type == game.HumanPlayer {
				UIGameState.ClientPlayer = game.PlayerTurn(pos)
				break
			}
		}
	}

	UIGameState.Board.State = gameStateResponse.GameState
	UIGameState.Winner = gameStateResponse.Winner
	fmt.Printf("Current Board State: %s\n", UIGameState.Board.State)
	UIGameState.PossibleMoves = make(map[int]*game.Move)

	UIGameState.Board.calculatePositions()
}
