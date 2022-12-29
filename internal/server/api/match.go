package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/server/game"
)

type AddToMatchQueueRequest struct {
	CheckerPosition int `json:"checkerPosition"` // 17
}
type AddToMatchQueueResponse struct {
	Moves []game.Move `json:"moves"` // 24, 26
}

func AddToMatchQueueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("AddToMatchQueueHandler")

	request := PossibleMovesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleErr(w, fmt.Errorf("PossibleMovesHandler err: %s", err))
		return
	}

	fmt.Printf("PossibleMovesHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	moves, captureMoves := game.GameState.Board.PossibleMoves(request.CheckerPosition)
	moves = append(moves, captureMoves...)
	resp := PossibleMovesResponse{
		Moves: moves,
	}

	fmt.Printf("PossibleMovesHandler response: %+v\n", resp)

	json.NewEncoder(w).Encode(resp)
}
