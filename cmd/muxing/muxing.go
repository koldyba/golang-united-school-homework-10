package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GolangUnited/helloweb/rest/handlers"
	"github.com/julienschmidt/httprouter"
)

func GetBad2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
/*
func Start(host string, port int) {
		router := mux.NewRouter()

		log.Println(fmt.Printf("Starting API server using Gorilla mux on %s:%d\n", host, port))

		router.HandleFunc("/bad", handlers.GetBad).Methods("GET")
		router.HandleFunc("/name/{PARAM}", handlers.GetName).Methods("GET")
		router.HandleFunc("/data", handlers.PostData).Methods("POST")
		router.HandleFunc("/headers", handlers.PostHeaders).Methods("POST")

		if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
				log.Fatal(err)
			}

		}
*/

// httpRouter
// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := httprouter.New()

	log.Println(fmt.Printf("Starting API server using httpRouter on %s:%d\n", host, port))

	router.GET("/bad", handlers.GetBad)
	router.GET("/name/:param", handlers.GetName)
	router.POST("/data", handlers.PostData)
	router.POST("/headers", handlers.PostHeaders)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
