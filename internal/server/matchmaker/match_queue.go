package matchmaker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/pion/webrtc/v3"
)

// type Match struct {
// 	GameState *game.Game
// 	Player1
// }

// type Peers struct {
// 	ListLock    sync.RWMutex
// 	Connections []PeerConnectionState
// }

// type PeerConnectionState struct {
// 	PeerConnection *webrtc.PeerConnection
// 	Websocket      *ThreadSafeWriter
// }

// type ThreadSafeWriter struct {
// 	Conn  *websocket.Conn
// 	Mutex sync.Mutex
// }

type ClientInfo struct {
	// MappedAddress stun.XORMappedAddress
	// Peers [2]Peer
	SessionDescription webrtc.SessionDescription
}

type Match struct {
	sync.Locker
	Player1, Player2 *ClientInfo
}

var (
	matchWaitingQueue chan ClientInfo
	// matches           = []Match{}
)

func AddToMatchQueue(clientInfo ClientInfo) {
	matchWaitingQueue <- clientInfo
}

func RunMatchMaker() {
	const maxClientWaiting = 2
	matchWaitingQueue = make(chan ClientInfo, maxClientWaiting)

	var match *Match

	for {
		player := <-matchWaitingQueue

		if match == nil {
			match.Unlock()
			match = &Match{
				Player1: &ClientInfo{
					SessionDescription: player.SessionDescription,
				},
			}
			match.Lock()
		} else if match.Player1 == nil {
			match.Unlock()
			match.Player1 = &ClientInfo{
				SessionDescription: player.SessionDescription,
			}
			match.Lock()
		} else if match.Player2 == nil {
			match.Unlock()
			match.Player2 = &ClientInfo{
				SessionDescription: player.SessionDescription,
			}
			match.Lock()
		}
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
