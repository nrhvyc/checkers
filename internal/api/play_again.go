package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type PlayAgainRequest struct{}
type PlayAgainResponse struct {
	GameState  string `json:"gameState"`
	PlayerTurn bool   `json:"playerTurn"` // false = black's turn; true = white's turn
	Winner     game.Winner
}

func PlayAgainHandler(w http.ResponseWriter, r *http.Request) {
	request := GameStateRequest{}
	fmt.Printf("PlayAgainHandler request: %+v\n", request)

	game.GameState = game.NewGame(game.GameState.GameMode)

	w.Header().Set("Content-Type", "application/json")

	resp := GameStateResponse{
		GameState:  game.GameState.StateToString(),
		PlayerTurn: game.GameState.PlayerTurn,
		Winner:     game.GameState.Winner,
	}

	json.NewEncoder(w).Encode(resp)
}
