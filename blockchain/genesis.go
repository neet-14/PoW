package blockchain

import (
 "time"
)

// CreateGenesisBlock creates the first block in the blockchain (genesis block).
func CreateGenesisBlock(difficulty int) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
		Index:        0,
		Timestamp:    timestamp,
		PrevHash:     "0",
		Data:         "Genesis Block",
		Nonce:        0,
		Difficulty:   difficulty,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
return genesisBlock
}