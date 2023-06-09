package main

import "blockchain-from-sratch/blockchain"

const DIFFICULTY = 3
const MINER_REWARD = 10.0
const BLOCKCHAIN_ADDRESS = "chi-huang-blockchain.io"

func main() {
	bc := blockchain.InitBlockchain(BLOCKCHAIN_ADDRESS)

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
	bc.RewardMiner("Amo", MINER_REWARD)

	bc.PrintChain()
	bc.PrintTransactionPool()
}
