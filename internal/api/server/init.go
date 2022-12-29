package server

import (
	"fmt"
	"net/http"
)

func handleErr(w http.ResponseWriter, err error) {
	fmt.Printf("we got an error: %s\n", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
