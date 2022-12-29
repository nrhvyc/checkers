package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	serverAPI "github.com/nrhvyc/checkers/internal/server/api"
	"github.com/nrhvyc/checkers/internal/server/game"
	"github.com/nrhvyc/checkers/internal/server/matchmaker"

	"github.com/pion/stun"
	"github.com/pion/webrtc/v3"
)

func RequestTwoPlayerMatch() {
	go AnswerServer()

	addPlayerToMatchQueue()
}

func addPlayerToMatchQueue() serverAPI.AddToMatchQueueResponse {
	clientAddress, err := ClientAddress()
	if err != nil {
		log.Fatalf("unable to retrieve ClientAddress err: %s", clientAddress)
	} else {
		log.Printf("clientAddress: %s", clientAddress.IP) // TODO: make this debug only for prod release
	}

	addPlayerRequest := serverAPI.AddToMatchQueueRequest{
		GameMode: game.TwoPlayer,
		ClientInfo: matchmaker.ClientInfo{
			MappedAddress: clientAddress,
		},
	}
	req, err := json.Marshal(addPlayerRequest)
	if err != nil {
		fmt.Printf("error marshalling PossibleMovesRequest err: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:7790/server/api/match/add",
		bytes.NewBuffer(req))
	if err != nil {
		fmt.Printf("error creating request err: %s\n", err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("client.Do err: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Game OnMount() err: %s", err)
	}

	var addToMatchQueueResponse serverAPI.AddToMatchQueueResponse
	json.Unmarshal(body, &addToMatchQueueResponse)

	return addToMatchQueueResponse
}

func signalCandidate(addr string, c *webrtc.ICECandidate) error {
	payload := []byte(c.ToJSON().Candidate)
	resp, err := http.Post(fmt.Sprintf("http://%s/candidate", addr),
		"application/json; charset=utf-8", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	if closeErr := resp.Body.Close(); closeErr != nil {
		return closeErr
	}

	return nil
}

func ClientAddress() (stun.XORMappedAddress, error) {
	// Creating a "connection" to STUN server.
	conn, err := stun.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		return stun.XORMappedAddress{}, err
	}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	var xorMappedAddress stun.XORMappedAddress

	// Sending request to STUN server, waiting for response message.
	if err := conn.Do(message, func(res stun.Event) {
		if res.Error != nil {
			log.Fatal(res.Error)
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			log.Fatal(err)
		}
		fmt.Println("your IP is", xorAddr.IP)

		xorMappedAddress = xorAddr

	}); err != nil {
		return stun.XORMappedAddress{}, err
	}

	return xorMappedAddress, nil
}
