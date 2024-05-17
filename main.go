package main

import(
	"fmt"
	"consensus/blockchain"
)

func main(){
	fmt.Println("Hello world")
	difficulty := 3 // Adjust the difficulty level as needed

 // Test Proof of Work
 powBlockchain := []blockchain.Block{blockchain.CreateGenesisBlock(difficulty)}

 // Generate a new block with some data using Proof of Work
 newBlockData := "Block Data"
 newPowBlock := blockchain.GenerateNewBlockWithPoW(powBlockchain[len(powBlockchain)-1], newBlockData, difficulty)
 powBlockchain = append(powBlockchain, newPowBlock)

 // Print the Proof of Work blockchain
 fmt.Println("Proof of Work Blockchain:")
 for _, block := range powBlockchain {
  fmt.Printf("Index: %d\n", block.Index)
  fmt.Printf("Timestamp: %d\n", block.Timestamp)
  fmt.Printf("PrevHash: %s\n", block.PrevHash)
  fmt.Printf("Data: %s\n", block.Data)
  fmt.Printf("Nonce: %d\n", block.Nonce)
  fmt.Printf("Difficulty: %d\n", block.Difficulty)
  fmt.Printf("Hash: %s\n\n", block.Hash)
 }
}