package blockchain

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int64
	prevHash     string
	timestamp    int64
	transactions []*Transcation
}

func NewBlock(nonce int64, prevHash string) *Block {
	timestamp := time.Now().Unix()
	return &Block{
		nonce:     nonce,
		prevHash:  prevHash,
		timestamp: timestamp,
	}
}

func (block *Block) GenerateHash() string {
	return Sha256Hash()
}

func (block *Block) PrintBlock() {
	fmt.Printf("nonce: %d\nprevHash: %s\ntimestamp: %d\ntransactions:\n", block.nonce, block.prevHash, block.timestamp)
	for _, tr := range block.transactions {
		tr.PrintTransction()
	}
}
