package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("something went wrong", message)
	}
	type empty struct{}

	jsonResponseWriter(w, code,  message, empty{})
}
type ApiResponse[T any]struct{
		Message string `json:"message"`
		Data T `json:"data,omitempty"`
	}
func jsonResponseWriter[T any](w http.ResponseWriter, code int, message string,  payload T) {
	
	apiRes := ApiResponse[T]{
		Message: message,
		Data: payload,
	}
	data, err := json.Marshal(apiRes)
	if err != nil {
		log.Println("error occured in json reponse writer", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}