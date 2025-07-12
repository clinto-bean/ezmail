package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	port string
}

func (a *API) Start() {
	log.Println("Initiating API")
	r := mux.NewRouter()
	r.Use(middlewareCORS)
	r.Use(middlewareLogging)
	r.HandleFunc("/ping", a.checkAPIStatus)
	r.HandleFunc("/users/new", a.handlerCreateUser)
	err := http.ListenAndServe(":"+a.port, r)
	if err != nil {
		panic(err)
	}
}

func (a *API) checkAPIStatus(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, 200, "pong")
}

func (a *API) checkAPIError(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, 200, "success")
}