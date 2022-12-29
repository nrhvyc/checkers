package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pion/stun"
	"github.com/rs/cors"

	"github.com/nrhvyc/checkers/internal/client/ui"
	serverAPI "github.com/nrhvyc/checkers/internal/server/api"
)

func main() {
	time.Sleep(time.Second * 5)
	fmt.Println("starting...")
	// go matchmaker.RunMatchMaker()

	appHandler := &app.Handler{
		Title:  "Checkers",
		Styles: []string{"/web/styles.css"},
	}

	app.Route("/", &ui.Game{})

	app.RunWhenOnBrowser()

	mux := http.NewServeMux()

	/*
		Register Server API Routes
	*/
	mux.HandleFunc("/server/api/game/state", serverAPI.GameStateHandler)
	mux.HandleFunc("/server/api/game/new", serverAPI.NewGameHandler)
	mux.HandleFunc("/server/api/game/play-again", serverAPI.PlayAgainHandler)

	mux.HandleFunc("/server/api/checker/possible-moves", serverAPI.PossibleMovesHandler)
	mux.HandleFunc("/server/api/checker/move", serverAPI.CheckerMoveHandler)

	mux.HandleFunc("/server/api/match/add", serverAPI.AddToMatchQueueHandler)

	/*
		Register WASM Routes
	*/
	mux.Handle("/", appHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	router := c.Handler(mux)

	addr, _ := serverAddress()
	fmt.Printf("addr: %s", addr.IP.String())

	if err := http.ListenAndServe(":7790", router); err != nil {
		panic(err)
	}
}

func serverAddress() (stun.XORMappedAddress, error) {
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
