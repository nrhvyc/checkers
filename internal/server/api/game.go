package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/server/game"
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

type NewGameRequest struct {
	GameMode       game.GameMode
	ClientAnswerIP string
}
type NewGameResponse struct {
	GameState  string          `json:"gameState"`
	PlayerTurn game.PlayerTurn `json:"playerTurn"` // false = black's turn; true = white's turn
	Players    [2]game.Player  `json:"players"`
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
		Players:    game.GameState.Players,
	}

	fmt.Printf("NewGameHandler resp: %+v\n", resp)
	json.NewEncoder(w).Encode(resp)
}

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
