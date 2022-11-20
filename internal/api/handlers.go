package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type GameStateRequest struct{}
type GameStateResponse struct {
	GameState  string `json:"gameState"`
	PlayerTurn bool   `json:"playerTurn"` // false = black's turn; true = white's turn
	Winner     game.Winner
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

	resp := GameStateResponse{
		GameState:  game.GameState.StateToString(),
		PlayerTurn: game.GameState.PlayerTurn,
		Winner:     game.GameState.Winner,
	}

	json.NewEncoder(w).Encode(resp)
}

func handleErr(w http.ResponseWriter, err error) {
	fmt.Printf("we got an error: %s\n", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
