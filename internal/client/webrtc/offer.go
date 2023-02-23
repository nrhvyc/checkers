package webrtc

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/pion/webrtc/v3"
)

type offerPeer struct {
	sync.Mutex
	connection        *webrtc.PeerConnection
	pendingCandidates []*webrtc.ICECandidate
}

func OfferServer(offerChan chan webrtc.SessionDescription) {
	log.Printf("OfferServer()")
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
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	offerPeer := offerPeer{
		connection: peerConnection,
	}
	fmt.Println("offer peer connection created")
	fmt.Println("registering offer peer connection handlers...")

	// When an ICE candidate is available send to the other Pion instance
	// the other Pion instance will add this candidate by calling AddICECandidate
	peerConnection.OnICECandidate(offerPeer.handleICECandidate)

	// A HTTP handler that allows the other Pion instance to send us ICE candidates
	// This allows us to add ICE candidates faster, we don't have to wait for STUN or TURN
	// candidates which may be slower
	http.HandleFunc("/candidate", offerPeer.handleCandidateReceived)

	// A HTTP handler that processes a SessionDescription given to us from the other Pion process
	http.HandleFunc("/sdp", offerPeer.handleSessionDescriptionReceived)

	// Start HTTP server that accepts requests from the answer process
	go func() { panic(http.ListenAndServe(offerAddr, nil)) }()

	// Create a datachannel with label 'data'
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(offerPeer.handleStateChange)

	// Register channel opening handling
	dataChannel.OnOpen(offerPeer.handleDataChannelMessage)

	// Register text message handling
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message from DataChannel '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
	})
	fmt.Println("handlers registered.")
	fmt.Println("sending offer...")

	// Create an offer to send to the other process
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	// Note: this will start the gathering of ICE candidates
	if err = peerConnection.SetLocalDescription(offer); err != nil {
		panic(err)
	}

	offerChan <- offer // send the offer back to the main thread

	// Block forever
	select {}
}

func (peer *offerPeer) handleICECandidate(c *webrtc.ICECandidate) {
	if c == nil {
		return
	}

	peer.Lock()
	defer peer.Unlock()

	desc := peer.connection.RemoteDescription()
	if desc == nil {
		peer.pendingCandidates = append(peer.pendingCandidates, c)
	} else if onICECandidateErr := signalCandidate(answerAddr, c); onICECandidateErr != nil {
		panic(onICECandidateErr)
	}
}

func (peer *offerPeer) handleCandidateReceived(w http.ResponseWriter, r *http.Request) {
	candidate, candidateErr := io.ReadAll(r.Body)
	if candidateErr != nil {
		panic(candidateErr)
	}
	iceCandidate := webrtc.ICECandidateInit{Candidate: string(candidate)}
	if candidateErr := peer.connection.AddICECandidate(iceCandidate); candidateErr != nil {
		panic(candidateErr)
	}
}

func (peer *offerPeer) handleSessionDescriptionReceived(w http.ResponseWriter, r *http.Request) {
	sdp := webrtc.SessionDescription{}
	if sdpErr := json.NewDecoder(r.Body).Decode(&sdp); sdpErr != nil {
		panic(sdpErr)
	}

	if sdpErr := peer.connection.SetRemoteDescription(sdp); sdpErr != nil {
		panic(sdpErr)
	}

	peer.Lock()
	defer peer.Unlock()

	for _, c := range peer.pendingCandidates {
		if onICECandidateErr := signalCandidate(answerAddr, c); onICECandidateErr != nil {
			panic(onICECandidateErr)
		}
	}
}

func (peer *offerPeer) handleStateChange(s webrtc.PeerConnectionState) {
	fmt.Printf("Peer Connection State has changed: %s\n", s.String())

	if s == webrtc.PeerConnectionStateFailed {
		// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
		// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
		// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
		fmt.Println("Peer Connection has gone to failed exiting")
		os.Exit(0)
	}
}

func (peer *offerPeer) handleDataChannelMessage() {
	// fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", dataChannel.Label(), dataChannel.ID())

	// for range time.NewTicker(5 * time.Second).C {
	// 	rand.Seed(time.Now().Unix())
	// 	randInt := rand.Intn(99999999)
	// 	message := fmt.Sprintf("message %d", randInt)
	// 	fmt.Printf("Sending '%s'\n", message)

	// 	// Send the message as text
	// 	sendTextErr := dataChannel.SendText(message)
	// 	if sendTextErr != nil {
	// 		panic(sendTextErr)
	// 	}
	// }
}
