package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, msg string) {
	type response struct {
		Error string `json:"error"`
	}
	jsonResponse(w, code, response{
		Error: msg,
	})
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}){
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf(`JSON Error: %s`, err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func decodeJSON(payload io.Reader, params interface{}) (error) {
	decoder := json.NewDecoder(payload)
	err := decoder.Decode(params)
	if err != nil {
		return err
	}
	return nil
}