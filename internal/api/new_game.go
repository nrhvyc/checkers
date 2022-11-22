package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type NewGameRequest struct {
	GameMode game.GameMode
}
type NewGameResponse struct {
	GameState  string `json:"gameState"`
	PlayerTurn bool   `json:"playerTurn"` // false = black's turn; true = white's turn
}

func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	request := NewGameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleErr(w, fmt.Errorf("NewGameHandler err: %s", err))
		return
	}

	fmt.Printf("NewGameHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	game.GameState = game.NewGame(request.GameMode)

	resp := NewGameResponse{
		GameState:  game.GameState.StateToString(),
		PlayerTurn: game.GameState.PlayerTurn,
	}

	fmt.Printf("NewGameHandler resp: %+v\n", resp)
	json.NewEncoder(w).Encode(resp)
}
