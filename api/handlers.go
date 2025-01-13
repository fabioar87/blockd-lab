package main

import (
	"errors"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	content := "API service index\n"
	replyContentText(w, r, http.StatusOK, content)
}

var (
	ErrNotFound    = errors.New("not found")
	ErrInvalidData = errors.New("invalid data")
)

// func todoRouter(todoFile string, l sync.Locker) http.HandlerFunc {}
