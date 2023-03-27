package blockchain

import (
	"fmt"
	"time"
)

type Transcation struct {
	senderAdress   string
	receiverAdress string
	amount         float64
	timestamp      int64
}

func (tr *Transcation) PrintTransction() {
	fmt.Printf("--------\n")
	fmt.Printf(" Sender Adress: %s\n", tr.senderAdress)
	fmt.Printf(" Reciever Adress: %s\n", tr.receiverAdress)
	fmt.Printf(" Amount: %.4f\n", tr.amount)
	fmt.Printf(" timestamp: %d\n", tr.timestamp)
}

func NewTransaction(senderAdress string, receiverAdress string, amount float64) *Transcation {
	timestamp := time.Now().Unix()
	return &Transcation{
		senderAdress:   senderAdress,
		receiverAdress: receiverAdress,
		amount:         amount,
		timestamp:      timestamp,
	}
}
