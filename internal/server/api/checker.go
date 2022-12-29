package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/server/game"
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

	followUpMoves := game.GameState.Move(request.Move)

	if game.GameState.Players[game.GameState.PlayerTurn].Type == game.AIPlayer {
		game.GameState.AIMove(game.GameState.PlayerTurn)
		followUpMoves = []game.Move{}
	}

	resp := CheckerMoveResponse{
		WasAllowed:    true,
		GameState:     game.GameState.StateToString(),
		PlayerTurn:    game.GameState.PlayerTurn,
		Winner:        game.GameState.Winner,
		FollowUpMoves: followUpMoves,
	}

	json.NewEncoder(w).Encode(resp)
}

type PossibleMovesRequest struct {
	CheckerPosition int `json:"checkerPosition"` // 17
}
type PossibleMovesResponse struct {
	Moves []game.Move `json:"moves"` // 24, 26
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

	moves, captureMoves := game.GameState.Board.PossibleMoves(request.CheckerPosition)
	moves = append(moves, captureMoves...)
	resp := PossibleMovesResponse{
		Moves: moves,
	}

	fmt.Printf("PossibleMovesHandler response: %+v\n", resp)

	json.NewEncoder(w).Encode(resp)
}
