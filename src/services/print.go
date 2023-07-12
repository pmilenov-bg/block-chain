package services

import (
	"encoding/hex"
	"fmt"
	. "hash-chain/src/models"
)

func PrintFormattedBlocks(blocks []Block) {
	fmt.Printf("\nBlockChain blocks: %d\n\n", len(blocks))
	for _, block := range blocks {
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("PreviousBlock Hash: %s\n", hex.EncodeToString(block.PrevBlockHash))
		fmt.Printf("Block Hash: %s\n\n", hex.EncodeToString(block.Hash))
	}
}
