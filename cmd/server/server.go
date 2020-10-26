package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:  "Checkers",
		Styles: []string{"/web/styles.css"},
	}

	if err := http.ListenAndServe(":7777", h); err != nil {
		panic(err)
	}
}
