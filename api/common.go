package api

import (
	"net/http"
)

// Package wide error messages
const (
	EmptyBody         = "Body of the request is empty"
	InvalidBody       = "Body of the request is invalid"
	UnableToParseBody = "Unable to parse request body"
	NoParamsFound     = "No parameters found on request"
)

func openCors(handler http.Handler, domain string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", domain)
		handler.ServeHTTP(w, r)
	})
}
