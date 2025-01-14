package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Transaction struct {
	gorm.Model
	TransactionId string  `json:"transactionId"`
	Amount        float64 `json:"amount"`
	TimeStamp     string  `json:"timestamp"`
}

func InitDB() (*gorm.DB, error) {
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("root:%s@tcp(db:3306)/transaction", dbPassword)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if !db.Migrator().HasTable(&Transaction{}) {
		db.Migrator().CreateTable(&Transaction{})
	}

	return db, nil
}

//type Batch []models

//func (b *Batch) AddTransaction(tID string, amount float64) {
//	t := models{
//		TransactionId: tID,
//		Amount:        amount,
//		TimeStamp:     time.Now().Format("YYYY-MM-DDTHHmmssZ"),
//	}
//
//	*b = append(*b, t)
//}
//
//func (b *Batch) DeleteTransaction(tID string) error {
//	bs := *b
//
//}
