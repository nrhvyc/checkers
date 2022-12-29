package matchmaker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pion/stun"
)

// type Match struct {
// 	GameState *game.Game
// 	Player1
// }

type ClientInfo struct {
	MappedAddress stun.XORMappedAddress
}

var matchWaitingQueue chan ClientInfo

func AddToMatchQueue(clientInfo ClientInfo) {
	matchWaitingQueue <- clientInfo
}

func RunMatchMaker() {
	const maxClientWaiting = 2
	matchWaitingQueue = make(chan ClientInfo, maxClientWaiting)

	// matchedPlayers := make([]ClientInfo, 2)

	for {
		// clientInfo := <-matchWaitingQueue // Blocks until player requests match

		// fmt.Printf("clientInfo: %+v", clientInfo)

		// matchedPlayers = append(matchedPlayers, clientInfo)

		// if len(matchedPlayers) > 1 {
		// 	// Blocks until two players requests match
		// callClientEstablishPeerConnection(<-matchWaitingQueue, <-matchWaitingQueue)
		clientOne := <-matchWaitingQueue
		fmt.Printf("clientOne: %+v", clientOne)
		// callClientEstablishPeerConnection(<-matchWaitingQueue, ClientInfo{})
		callClientEstablishPeerConnection(clientOne, ClientInfo{})
		// }
	}
}

type EstablishPeerConnectionRequest struct {
	AnswerAddress string `json:"answerAddress"`
}
type EstablishPeerConnectionResponse struct {
}

func callClientEstablishPeerConnection(clientOne, clientTwo ClientInfo) {
	establishPeerConnectionRequest := EstablishPeerConnectionRequest{
		AnswerAddress: "localhost:60000", // TODO: don't hard code this
	}
	req, err := json.Marshal(establishPeerConnectionRequest)
	if err != nil {
		fmt.Printf("error marshalling EstablishPeerConnectionRequest err: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:7791/client/match/start",
		bytes.NewBuffer(req))
	if err != nil {
		fmt.Printf("error creating request err: %s\n", err)
	}

	request.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		fmt.Printf("httpClient.Do err: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Game OnMount() err: %s", err)
	}

	establishPeerConnectionHandlerResponse := EstablishPeerConnectionResponse{}
	json.Unmarshal(body, &establishPeerConnectionHandlerResponse)
}
