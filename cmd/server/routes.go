package main

import (
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/rs/cors"

	serverAPI "github.com/nrhvyc/checkers/internal/api/server"
	"github.com/nrhvyc/checkers/internal/matchmaker"
	"github.com/nrhvyc/checkers/internal/ui"
)

func main() {
	time.Sleep(time.Second * 10)
	go matchmaker.RunMatchMaker()

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

	if err := http.ListenAndServe(":7790", router); err != nil {
		panic(err)
	}
}
