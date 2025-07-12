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
	r.HandleFunc("/ping", a.CheckAPIStatus)
	err := http.ListenAndServe(":"+a.port, r)
	if err != nil {
		panic(err)
	}
}

func (a *API) CheckAPIStatus(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, 200, "pong")
}

func (a *API) CheckAPIError(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, 200, "success")
}