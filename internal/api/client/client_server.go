package client

import (
	"net/http"

	"github.com/rs/cors"
)

func ClientServer() {
	mux := http.NewServeMux()

	// Register API Routes
	mux.HandleFunc("/client/match/start", EstablishPeerConnectionHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	router := c.Handler(mux)

	if err := http.ListenAndServe(":7791", router); err != nil {
		panic(err)
	}
}
