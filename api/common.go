package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Package wide error messages
const (
	EmptyBody         = "body of the request is empty"
	InvalidBody       = "body of the request is invalid"
	UnableToParseBody = "enable to parse request body"
	NoParamsFound     = "no parameters found on request"
)

func openCors(handler http.Handler, domain string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", domain)
		handler.ServeHTTP(w, r)
	})
}

func writeJSONResponse(writer http.ResponseWriter, returnStatus int, responseBody interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(returnStatus)
	if err := json.NewEncoder(writer).Encode(responseBody); err != nil {
		log.Printf("Problem encoding new response body: %s", err.Error())
	}
}
