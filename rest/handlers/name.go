package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Gorilla mux
/*
func GetName(w http.ResponseWriter, r *http.Request) {
	log.Printf("routed to /data")
	vars := mux.Vars(r)
	resp := fmt.Sprintf("Hello, %s!", vars["PARAM"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
*/

// httpRouter
func GetName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("routed to /data")
	resp := fmt.Sprintf("Hello, %s!", p.ByName("param"))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
