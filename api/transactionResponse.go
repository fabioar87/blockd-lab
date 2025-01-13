package main

import (
	"api/transaction"
	"encoding/json"
	"time"
)

type transactionResponse struct {
	Results transaction.Batch `json:"results"`
}

func (r *transactionResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      transaction.Batch `json:"results"`
		Date         int64             `json:"date"`
		TotalResults int               `json:"total_results"`
	}{
		Results:      r.Results,
		Date:         time.Now().Unix(),
		TotalResults: len(r.Results),
	}

	return json.Marshal(resp)
}
