package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Gorilla mux
/*
func PostData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	log.Printf("routed to /data")
	if err != nil {
		log.Printf("error occurred while parsing request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := fmt.Sprintf("I got message:\n%s", string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
*/

// httpRouter
func PostData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	log.Printf("routed to /data")
	if err != nil {
		log.Printf("error occurred while parsing request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := fmt.Sprintf("I got message:\n%s", string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
