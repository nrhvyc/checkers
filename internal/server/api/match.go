package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrhvyc/checkers/internal/server/game"
	"github.com/nrhvyc/checkers/internal/server/matchmaker"
)

type AddToMatchQueueRequest struct {
	GameMode   game.GameMode
	ClientInfo matchmaker.ClientInfo
}

type AddToMatchQueueResponse struct {
	Status string
}

func AddToMatchQueueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddToMatchQueueHandler")

	request := PossibleMovesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleErr(w, fmt.Errorf("AddToMatchQueueHandler err: %s", err))
		return
	}

	fmt.Printf("AddToMatchQueueHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	matchmaker.AddToMatchQueue(matchmaker.ClientInfo{}) // TODO: send this from the frontend code

	resp := AddToMatchQueueResponse{Status: "test status"}

	fmt.Printf("AddToMatchQueueHandler response: %+v\n", resp)

	json.NewEncoder(w).Encode(resp)
}
