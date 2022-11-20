package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/rs/cors"

	"github.com/nrhvyc/checkers/internal/api"
	"github.com/nrhvyc/checkers/internal/ui"
)

func main() {
	appHandler := &app.Handler{
		Title:  "Checkers",
		Styles: []string{"/web/styles.css"},
	}

	app.Route("/", &ui.Game{})

	app.RunWhenOnBrowser()

	mux := http.NewServeMux()

	// Register API Routes
	mux.HandleFunc("/api/game/state", api.GameStateHandler)
	mux.HandleFunc("/api/game/play-again", api.PlayAgainHandler)
	mux.HandleFunc("/api/checker/possible-moves", api.PossibleMovesHandler)
	mux.HandleFunc("/api/checker/move", api.CheckerMoveHandler)

	// Register WASM Routes
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
