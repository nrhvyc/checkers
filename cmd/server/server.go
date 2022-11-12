package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/rs/cors"

	"github.com/nrhvyc/checkers/internal/api"
	"github.com/nrhvyc/checkers/internal/ui"
)

// type hello struct {
// 	app.Compo
// }

// // The Render method is where the component appearance is defined. Here, a
// // "Hello World!" is displayed as a heading.
// func (h *hello) Render() app.UI {
// 	return app.H1().Text("Hello World!")
// }

func main() {
	appHandler := &app.Handler{
		Title:  "Checkers",
		Styles: []string{"/web/styles.css"},
	}

	// g := ui.NewGame()

	app.Route("/", &ui.Game{})
	// app.Route("/", &hello{})

	app.RunWhenOnBrowser()

	mux := http.NewServeMux()

	// Register API Routes
	mux.HandleFunc("/api/game/state", api.GameStateHandler)
	mux.HandleFunc("/api/checker/possible-moves", api.PossibleMovesHandler)

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
