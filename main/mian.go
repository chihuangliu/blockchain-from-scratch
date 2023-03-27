package main

import "blockchain-from-sratch/blockchain"

func main() {
	bc := blockchain.InitBlockchain()

	newHash := bc.GenerateHashOfLastBlock()
	newBlock := blockchain.NewBlock(20, newHash)
	bc.AddBlock(newBlock)

	newTransaction1 := blockchain.NewTransaction("A", "B", 80.8888)
	bc.AddTransactionToPool(newTransaction1)

	newTransaction2 := blockchain.NewTransaction("X", "Y", 50.1111)
	bc.AddTransactionToPool(newTransaction2)

	bc.PrintChain()
	bc.PrintTranscationPool()

	bc.TransferTransactionsFromPoolToBlock()
	bc.PrintChain()
	bc.PrintTranscationPool()
}
