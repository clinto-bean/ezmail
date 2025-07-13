package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type API struct {
	port string
}

func (a *API) Start() error {
	log.Println("Initiating API")
	r := mux.NewRouter()
	r.Use(middlewareCORS)
	r.Use(middlewareLogging)
	// status checker handlers
	r.HandleFunc("/v1/shutdown", a.gracefulShutdown)
	r.HandleFunc("/v1/ping", a.checkAPIStatus)
	// user handlers
	r.HandleFunc("/v1/users/new", a.handleUser)
	srv := &http.Server{
		Addr: ":" + a.port,
		Handler: r,
	}
	return srv.ListenAndServe()
}

func (a *API) gracefulShutdown (w http.ResponseWriter, r *http.Request) {
	log.Println("Shutdown initiated")
	os.Exit(0)
}

func (a *API) checkAPIStatus(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, 200, "pong")
}

func (a *API) checkAPIError(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, 200, "success")
}