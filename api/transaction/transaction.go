package transaction

type transaction struct {
	TransactionId string
	Amount        float64
	TimeStamp     string
}

type Batch []transaction

//func (b *Batch) AddTransaction(tID string, amount float64) {
//	t := transaction{
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
