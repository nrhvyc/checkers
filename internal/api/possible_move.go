package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

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
		PossiblePositions: game.GameState.PossibleMoves(request.CheckerPosition),
	}

	json.NewEncoder(w).Encode(resp)
}
