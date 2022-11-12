package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type GameStateRequest struct{}
type GameStateResponse struct {
	Game game.Game `json:"game"`
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
		Game: *game.GameState,
	}

	json.NewEncoder(w).Encode(resp)
}

type PossibleMovesRequest struct {
	CheckerPosition int `json:"checkerPosition"` // 17
}
type PossibleMovesResponse struct {
	PossiblePositions []int `json:"possiblePositions"` // 24, 26
}

func PossibleMovesHandler(w http.ResponseWriter, r *http.Request) {
	request := PossibleMovesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleErr(w, fmt.Errorf("PossibleMovesHandler err: %s", err))
		return
	}

	fmt.Printf("PossibleMovesHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	resp := PossibleMovesResponse{
		PossiblePositions: []int{24, 26},
	}

	json.NewEncoder(w).Encode(resp)
}

func handleErr(w http.ResponseWriter, err error) {
	fmt.Printf("we got an error: %s\n", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
