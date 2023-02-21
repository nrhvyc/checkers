package webrtc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	client "github.com/nrhvyc/checkers/internal/client/webrtc"
)

type EstablishPeerConnectionRequest struct {
	AnswerAddress string `json:"answerAddress"`
}
type EstablishPeerConnectionResponse struct {
}

func EstablishPeerConnectionHandler(w http.ResponseWriter, r *http.Request) {
	request := EstablishPeerConnectionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleErr(w, fmt.Errorf("EstablishPeerConnectionHandler err: %s", err))
		return
	}

	fmt.Printf("EstablishPeerConnectionHandler request: %+v\n", request)

	w.Header().Set("Content-Type", "application/json")

	// client.EstablishPeerConnection(request.AnswerAddress)
	mappedAddress, err := client.ClientAddress()
	if err != nil {
		log.Fatal(err)
	}
	client.EstablishPeerConnection(mappedAddress)

	resp := EstablishPeerConnectionResponse{}

	fmt.Printf("EstablishPeerConnectionHandler response: %+v\n", resp)

	json.NewEncoder(w).Encode(resp)
}

func handleErr(w http.ResponseWriter, err error) {
	fmt.Printf("we got an error: %s\n", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
