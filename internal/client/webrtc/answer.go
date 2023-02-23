package webrtc

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"log"

	"github.com/pion/webrtc/v3"
)

var (
	offerAddr  = "localhost:50000"
	answerAddr = "localhost:60000"
)

type answerPeer struct {
	sync.Mutex
	connection        *webrtc.PeerConnection
	pendingCandidates []*webrtc.ICECandidate
}

func AnswerServer() {
	log.Printf("AnswerServer()")
	flag.Parse()

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := peerConnection.Close(); err != nil {
			fmt.Printf("cannot close peerConnection: %v\n", err)
		}
	}()
	answerPeer := answerPeer{
		connection: peerConnection,
	}
	fmt.Println("answer peer connection created")
	fmt.Println("registering answer peer connection handlers...")

	local := answerPeer.connection.LocalDescription()
	fmt.Printf("answer peer pending local desciption: %+v\n", local)

	// When an ICE candidate is available to send to the other Pion instance
	// the other Pion instance will add this candidate by calling AddICECandidate
	peerConnection.OnICECandidate(answerPeer.handleICECandidate)

	// A HTTP handler that allows the other Pion instance to send us ICE candidates
	// This allows us to add ICE candidates faster, we don't have to wait for STUN or TURN
	// candidates which may be slower
	http.HandleFunc("/candidate", answerPeer.handleCandidateReceived)

	// A HTTP handler that processes a SessionDescription given to us from the other Pion process
	http.HandleFunc("/sdp", answerPeer.handleSessionDescriptionReceived)

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(answerPeer.handleStateChange)

	// Register data channel creation handling
	peerConnection.OnDataChannel(answerPeer.handleDataChannelMessage)

	fmt.Println("handlers registered.")

	// Start HTTP server that accepts requests from the offer process to exchange SDP and Candidates
	panic(http.ListenAndServe(answerAddr, nil))
}

func (answer *answerPeer) handleICECandidate(c *webrtc.ICECandidate) {
	if c == nil {
		return
	}

	answer.Lock()
	defer answer.Unlock()

	desc := answer.connection.RemoteDescription()
	if desc == nil {
		answer.pendingCandidates = append(answer.pendingCandidates, c)
	} else if onICECandidateErr := signalCandidate(offerAddr, c); onICECandidateErr != nil {
		panic(onICECandidateErr)
	}
}

func (peer *answerPeer) handleSessionDescriptionReceived(w http.ResponseWriter, r *http.Request) {
	sdp := webrtc.SessionDescription{}
	if err := json.NewDecoder(r.Body).Decode(&sdp); err != nil {
		panic(err)
	}

	if err := peer.connection.SetRemoteDescription(sdp); err != nil {
		panic(err)
	}

	// Create an peer to send to the other process
	answer, err := peer.connection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	// Send our answer to the HTTP server listening in the other process
	payload, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/sdp", offerAddr), "application/json; charset=utf-8", bytes.NewReader(payload))
	if err != nil {
		panic(err)
	} else if closeErr := resp.Body.Close(); closeErr != nil {
		panic(closeErr)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	err = peer.connection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	peer.Lock()
	for _, c := range peer.pendingCandidates {
		onICECandidateErr := signalCandidate(offerAddr, c)
		if onICECandidateErr != nil {
			panic(onICECandidateErr)
		}
	}
	peer.Unlock()
}

func (peer *answerPeer) handleStateChange(s webrtc.PeerConnectionState) {
	fmt.Printf("Peer Connection State has changed: %s\n", s.String())

	if s == webrtc.PeerConnectionStateFailed {
		// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
		// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
		// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
		fmt.Println("Peer Connection has gone to failed exiting")
		os.Exit(0)
	}
}

func (peer *answerPeer) handleCandidateReceived(w http.ResponseWriter, r *http.Request) {
	candidate, candidateErr := io.ReadAll(r.Body)
	if candidateErr != nil {
		panic(candidateErr)
	}

	iceCandidate := webrtc.ICECandidateInit{Candidate: string(candidate)}
	if candidateErr := peer.connection.AddICECandidate(iceCandidate); candidateErr != nil {
		panic(candidateErr)
	}
}

func (peer *answerPeer) handleDataChannelMessage(d *webrtc.DataChannel) {
	fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

	// Register channel opening handling
	d.OnOpen(func() {
		fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n",
			d.Label(), d.ID())

		for range time.NewTicker(5 * time.Second).C {
			rand.Seed(time.Now().Unix())
			randInt := rand.Intn(99999999)
			message := fmt.Sprintf("message %d", randInt)
			fmt.Printf("Sending '%s'\n", message)

			// Send the message as text
			sendTextErr := d.SendText(message)
			if sendTextErr != nil {
				panic(sendTextErr)
			}
		}
	})

	// Register text message handling
	d.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
	})
}
