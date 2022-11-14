package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/game"
)

type CheckerMoveRequest struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type CheckerMoveResponse struct {
	WasAllowed bool
	GameState  string
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

	game.GameState.Move(request.From, request.To)

	resp := CheckerMoveResponse{
		WasAllowed: true,
		// GameState:  "_b_b_b_bb_b_b_b____b_b_bb_______________w_w_w_w__w_w_w_ww_w_w_w_",
		GameState: game.GameState.StateToString(),
	}

	json.NewEncoder(w).Encode(resp)
}
