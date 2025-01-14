package main

import (
	"api/models"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type DBClient struct {
	db *gorm.DB
}

type transactionResponse struct {
	Transaction models.Transaction `json:"transaction"`
	Data        interface{}        `json:"data"`
}

func main() {
	host := flag.String("h", "0.0.0.0", "API hostname")
	port := flag.Int("p", 8080, "API service port")
	flag.Parse()

	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}

	dbClient := &DBClient{db: db}
	router := mux.NewRouter()
	router.HandleFunc("/api/transactions", dbClient.GetTransactions).Methods("GET")
	router.HandleFunc("/api/transaction", dbClient.CreateTransaction).Methods("POST")
	router.HandleFunc("/api/health", dbClient.GetHealth).Methods("GET")

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
