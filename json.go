package main

import (
	"encoding/json"
	"log"
	"net/http"
);

func respondWithJSON(w http.ResponseWriter, code int , payload interface{}){
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("failed to marshal JSON: %v", payload)
		w.WriteHeader(500) // Internal Server Error
		return 
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)

	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int , msg string){
	if code > 499{
		log.Println("5XX server error: ", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}