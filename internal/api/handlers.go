package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type GameStateRequest struct{}
type GameStateResponse struct {
	GameMode   game.GameMode   `json:"gameMode"`
	GameState  string          `json:"gameState"`
	PlayerTurn game.PlayerTurn `json:"playerTurn"`
	Winner     game.Winner
	Players    [2]game.Player `json:"players"`
}

func GameStateHandler(w http.ResponseWriter, r *http.Request) {
	request := GameStateRequest{}
	// err := json.NewDecoder(r.Body).Decode(&request)
	// if err != nil {
	// 	handleErr(w, fmt.Errorf("GameStateHandler err: %s", err))
	// 	return
	// }

	fmt.Printf("GameStateHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	var resp GameStateResponse
	if game.GameState == nil || game.GameState.GameMode == game.NewGameMode {
		resp = GameStateResponse{
			GameMode: game.NewGameMode,
		}
	} else {
		resp = GameStateResponse{
			GameMode:   game.GameState.GameMode,
			GameState:  game.GameState.StateToString(),
			PlayerTurn: game.GameState.PlayerTurn,
			Winner:     game.GameState.Winner,
			Players:    game.GameState.Players,
		}
	}

	fmt.Printf("GameStateHandler resp: %+v\n", resp)
	json.NewEncoder(w).Encode(resp)
}

func handleErr(w http.ResponseWriter, err error) {
	fmt.Printf("we got an error: %s\n", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
