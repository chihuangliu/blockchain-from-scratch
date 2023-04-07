package blockchain

import (
	"fmt"
	"time"
)

type Transaction struct {
	senderAdress   string
	receiverAdress string
	amount         float64
	timestamp      int64
}

func (tr *Transaction) PrintTransaction() {
	fmt.Printf("--------\n")
	fmt.Printf(" Sender Adress: %s\n", tr.senderAdress)
	fmt.Printf(" Reciever Adress: %s\n", tr.receiverAdress)
	fmt.Printf(" Amount: %.4f\n", tr.amount)
	fmt.Printf(" timestamp: %d\n", tr.timestamp)
}

func NewTransaction(senderAdress string, receiverAdress string, amount float64) *Transaction {
	timestamp := time.Now().Unix()
	return &Transaction{
		senderAdress:   senderAdress,
		receiverAdress: receiverAdress,
		amount:         amount,
		timestamp:      timestamp,
	}
}
