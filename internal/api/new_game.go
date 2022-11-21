package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type NewGameRequest struct{}
type NewGameResponse struct {
	GameState  string `json:"gameState"`
	PlayerTurn bool   `json:"playerTurn"` // false = black's turn; true = white's turn
}

func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	request := NewGameRequest{}
	// err := json.NewDecoder(r.Body).Decode(&request)
	// if err != nil {
	// 	handleErr(w, fmt.Errorf("GameStateHandler err: %s", err))
	// 	return
	// }

	fmt.Printf("GameStateHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	resp := NewGameResponse{
		GameState:  game.GameState.StateToString(),
		PlayerTurn: game.GameState.PlayerTurn,
	}

	json.NewEncoder(w).Encode(resp)
}
