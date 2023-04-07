package main

import "blockchain-from-sratch/blockchain"

const DIFFICULTY = 3

func main() {
	bc := blockchain.InitBlockchain()

	newHash := bc.GenerateHashOfLastBlock()
	newNonce := bc.ProofOfWorkOfLastBlock(DIFFICULTY)
	newBlock := blockchain.NewBlock(newNonce, newHash)
	bc.AddBlock(newBlock)

	newTransaction1 := blockchain.NewTransaction("A", "B", 80.8888)
	bc.AddTransactionToPool(newTransaction1)

	newTransaction2 := blockchain.NewTransaction("X", "Y", 50.1111)
	bc.AddTransactionToPool(newTransaction2)

	// bc.PrintChain()
	// bc.PrintTransactionPool()

	bc.TransferTransactionsFromPoolToBlock()
	// bc.PrintChain()
	// bc.PrintTransactionPool()

	newHash2 := bc.GenerateHashOfLastBlock()
	newNonce2 := bc.ProofOfWorkOfLastBlock(DIFFICULTY)
	newBlock2 := blockchain.NewBlock(newNonce2, newHash2)
	bc.AddBlock(newBlock2)

	bc.PrintChain()
}
