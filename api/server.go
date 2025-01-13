package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func customMux() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", indexHandler)

	return m
}

func replyContentText(w http.ResponseWriter, r *http.Request, statusCode int, contentText string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write([]byte(contentText))
}

func replyContentJSON(w http.ResponseWriter, r *http.Request, statusCode int, resp *transactionResponse) {
	body, err := json.Marshal(resp)
	if err != nil {
		replyErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func replyErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Printf("%s %s Error: %d %s", r.URL, r.Method, statusCode, message)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
