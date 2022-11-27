package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type CheckerMoveRequest struct {
	Move game.Move `json:"move"`
}
type CheckerMoveResponse struct {
	WasAllowed    bool
	GameState     string
	PlayerTurn    game.PlayerTurn // false = black's turn; true = white's turn
	FollowUpMoves []game.Move
	Winner        game.Winner
}

func CheckerMoveHandler(w http.ResponseWriter, r *http.Request) {
	request := CheckerMoveRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleErr(w, fmt.Errorf("CheckerMoveHandler err: %s", err))
		return
	}

	fmt.Printf("CheckerMoveHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	// game.GameState.Move(request.From, request.To)
	followUpMoves := game.GameState.Move(request.Move)

	resp := CheckerMoveResponse{
		WasAllowed: true,
		// GameState:  "_b_b_b_bb_b_b_b____b_b_bb_______________w_w_w_w__w_w_w_ww_w_w_w_",
		GameState:     game.GameState.StateToString(),
		PlayerTurn:    game.GameState.PlayerTurn,
		Winner:        game.GameState.Winner,
		FollowUpMoves: followUpMoves,
	}

	json.NewEncoder(w).Encode(resp)
}
