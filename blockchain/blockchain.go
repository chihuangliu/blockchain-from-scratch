package blockchain

import "fmt"

type Blockchain struct {
	TransactionPool []*Transcation
	Chain           []*Block
}

func (bc *Blockchain) AddBlock(block *Block) {
	bc.Chain = append(bc.Chain, block)
}

func InitBlockchain() *Blockchain {
	bc := new(Blockchain)
	initialBlock := NewBlock(10, Sha256Hash())

	bc.AddBlock(initialBlock)
	return bc
}

func (bc *Blockchain) AddTransactionToPool(tr *Transcation) {
	bc.TransactionPool = append(bc.TransactionPool, tr)
}

func (bc *Blockchain) PrintChain() {
	for i, block := range bc.Chain {
		fmt.Printf("Block %d\n", i)
		block.PrintBlock()
		fmt.Printf("\n")
	}
}

func (bc *Blockchain) PrintTranscationPool() {
	fmt.Printf("Transcations in Pools:\n")
	for _, tc := range bc.TransactionPool {
		tc.PrintTransction()
		fmt.Printf("\n")
	}
}

func (bc *Blockchain) GenerateHashOfLastBlock() string {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newHash := lastBlock.GenerateHash()
	return newHash
}

func (bc *Blockchain) TransferTransactionsFromPoolToBlock() {
	fmt.Printf("Transform transactions from the pool to last block.....\n")
	lastBlock := bc.Chain[len(bc.Chain)-1]
	for _, tr := range bc.TransactionPool {
		lastBlock.transactions = append(lastBlock.transactions, tr)
	}
	bc.TransactionPool = []*Transcation{}
}
