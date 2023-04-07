package blockchain

import "fmt"

type Blockchain struct {
	TransactionPool []*Transaction
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

func (bc *Blockchain) AddTransactionToPool(tr *Transaction) {
	bc.TransactionPool = append(bc.TransactionPool, tr)
}

func (bc *Blockchain) PrintChain() {
	for i, block := range bc.Chain {
		fmt.Printf("Block %d\n", i)
		block.PrintBlock()
		fmt.Printf("\n")
	}
}

func (bc *Blockchain) PrintTransactionPool() {
	fmt.Printf("Transacations in Pools:\n")
	for _, tc := range bc.TransactionPool {
		tc.PrintTransaction()
		fmt.Printf("\n")
	}
}

func (bc *Blockchain) lastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) GenerateHashOfLastBlock() string {
	lastBlock := bc.lastBlock()
	newHash, _ := lastBlock.GenerateHash()
	return newHash
}

func (bc *Blockchain) ProofOfWorkOfLastBlock(difficulty int) int64 {
	lastBlock := bc.lastBlock()
	return lastBlock.ProofOfWork(difficulty)
}

func (bc *Blockchain) TransferTransactionsFromPoolToBlock() {
	fmt.Printf("Transform transactions from the pool to last block.....\n")
	lastBlock := bc.Chain[len(bc.Chain)-1]
	for _, tr := range bc.TransactionPool {
		lastBlock.transactions = append(lastBlock.transactions, tr)
	}
	bc.TransactionPool = []*Transaction{}
}
