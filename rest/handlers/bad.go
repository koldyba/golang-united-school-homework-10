package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Gorilla mux
/*
func GetBad(w http.ResponseWriter, r *http.Request) {
	log.Printf("routed to /bad")
	w.WriteHeader(http.StatusInternalServerError)
}
*/

// httpRouter
func GetBad(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("routed to /bad")
	w.WriteHeader(http.StatusInternalServerError)
}
