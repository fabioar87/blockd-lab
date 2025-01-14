package main

import (
	"api/models"
	"encoding/json"
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

func (driver *DBClient) GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	driver.db.Find(&transactions)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(transactions)
	w.Write(respJSON)
}

func (driver *DBClient) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid JSON payload"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := driver.db.Save(&transaction).Error; err != nil {
		http.Error(w, "Failed to insert record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// func todoRouter(todoFile string, l sync.Locker) http.HandlerFunc {}
