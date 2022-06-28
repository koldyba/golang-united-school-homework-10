package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Gorilla mux
/*
func PostHeaders(w http.ResponseWriter, r *http.Request) {
	log.Printf("routed to /headers")
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		log.Printf("wrong value in a header")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Printf("wrong value in b header")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := strconv.Itoa(a + b)
	w.Header().Add("a+b", result)
	w.WriteHeader(http.StatusOK)
}
*/

// httpRouter
func PostHeaders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("routed to /headers")
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		log.Printf("wrong value in a header")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Printf("wrong value in b header")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := strconv.Itoa(a + b)
	w.Header().Add("a+b", result)
	w.WriteHeader(http.StatusOK)
}
