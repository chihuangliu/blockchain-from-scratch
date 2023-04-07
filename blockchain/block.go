package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int64
	prevHash     string
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int64, prevHash string) *Block {
	timestamp := time.Now().Unix()
	return &Block{
		nonce:     nonce,
		prevHash:  prevHash,
		timestamp: timestamp,
	}
}

func (block *Block) jsonMarshal() (string, error) {
	blockMap := map[string]interface{}{
		"nonce":        block.nonce,
		"prevHash":     block.prevHash,
		"timestamp":    block.timestamp,
		"transactions": block.transactions,
	}

	jsonBytes, err := json.Marshal(blockMap)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes)

	return jsonString, nil
}

func (block *Block) GenerateHash() (string, error) {
	blockToJson, err := block.jsonMarshal()
	if err != nil {
		return "", err
	}
	jsonBytes := []byte(blockToJson)
	hashBytes := sha256.Sum256(jsonBytes)
	hashString := hex.EncodeToString(hashBytes[:])

	return hashString, nil
}

func (block *Block) PrintBlock() {
	fmt.Printf("nonce: %d\nprevHash: %s\ntimestamp: %d\ntransactions:\n", block.nonce, block.prevHash, block.timestamp)
	for _, tr := range block.transactions {
		tr.PrintTransaction()
	}
}

func (block *Block) ProofOfWork(difficulty int) int64 {
	n := int64(0)
	for {
		success := block.miningNonce(difficulty, n)
		if success {
			break
		}
		n += 1
	}

	return n
}

func (block *Block) miningNonce(difficulty int, nonce int64) bool {
	zeros := strings.Repeat("0", difficulty)

	transactions := block.transactions
	prevHash := block.prevHash
	blockMined := Block{nonce: nonce, prevHash: prevHash, timestamp: 0, transactions: transactions}

	hash, _ := blockMined.GenerateHash()
	subHash := hash[:difficulty]

	if zeros == subHash {
		return true
	} else {
		return false
	}
}
