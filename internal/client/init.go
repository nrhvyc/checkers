package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nrhvyc/checkers/internal/api"
	"github.com/pion/webrtc/v3"
)

func RequestTwoPlayerMatch() {
	go answerServer()

	addPlayerToMatchQueue()
}

func addPlayerToMatchQueue() {
	resp, err := http.Get("http://localhost:7790/api/game/state")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Game OnMount() err: %s", err)
	}
	gameStateResponse := api.GameStateResponse{}
	json.Unmarshal(body, &gameStateResponse)

	fmt.Printf("gameStateResponse.GameMode: %d\n", gameStateResponse.GameMode)
}

// "github.com/pion/webrtc/v3/examples/internal/signal"

// func NewPeer() {
// 	config := webrtc.Configuration{
// 		ICEServers: []webrtc.ICEServer{
// 			{
// 				URLs: []string{"stun:stun.l.google.com:19302"},
// 			},
// 		},
// 	}

// 	// Create a new RTCPeerConnection
// 	peerConnection, err := webrtc.NewPeerConnection(config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if cErr := peerConnection.Close(); cErr != nil {
// 			fmt.Printf("cannot close peerConnection: %v\n", cErr)
// 		}
// 	}()

// 	// Set the handler for Peer connection state
// 	// This will notify you when the peer has connected/disconnected
// 	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
// 		fmt.Printf("Peer Connection State has changed: %s\n", s.String())

// 		if s == webrtc.PeerConnectionStateFailed {
// 			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
// 			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
// 			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
// 			fmt.Println("Peer Connection has gone to failed exiting")
// 			os.Exit(0)
// 		}
// 	})

// 	// Register data channel creation handling
// 	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
// 		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

// 		// Register channel opening handling
// 		d.OnOpen(func() {
// 			fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", d.Label(), d.ID())

// 			for range time.NewTicker(5 * time.Second).C {
// 				// message := signal.RandSeq(15)
// 				message := "test message"
// 				fmt.Printf("Sending '%s'\n", message)

// 				// Send the message as text
// 				sendErr := d.SendText(message)
// 				if sendErr != nil {
// 					panic(sendErr)
// 				}
// 			}
// 		})

// 		// Register text message handling
// 		d.OnMessage(func(msg webrtc.DataChannelMessage) {
// 			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
// 		})
// 	})

// 	// Wait for the offer to be pasted
// 	// offer := webrtc.SessionDescription{}
// 	// signal.Decode(signal.MustReadStdin(), &offer)

// 	// Set the remote SessionDescription
// 	// err = peerConnection.SetRemoteDescription(offer)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// Create an answer
// 	answer, err := peerConnection.CreateAnswer(nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create channel that is blocked until ICE Gathering is complete
// 	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

// 	// Sets the LocalDescription, and starts our UDP listeners
// 	err = peerConnection.SetLocalDescription(answer)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Block until ICE Gathering is complete, disabling trickle ICE
// 	// we do this because we only can exchange one signaling message
// 	// in a production application you should exchange ICE Candidates via OnICECandidate
// 	<-gatherComplete

// 	// Output the answer in base64 so we can paste it in browser
// 	// fmt.Println(signal.Encode(*peerConnection.LocalDescription()))

// 	// Block forever
// 	select {}
// }

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
