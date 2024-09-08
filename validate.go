package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type errorVal struct {
		errorMsg string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		// w.Write(errorVal{errorMsg: err})
		return
	}
}
