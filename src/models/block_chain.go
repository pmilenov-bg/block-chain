package models

import (
	"fmt"
	"hash-chain/src/services"
)

type BlockChain struct {
	latestBlock Block
}

func GenerateBlockChain() *BlockChain {
	genesisBlock := &Block{
		Data:      CreateBlockData("Genesis"),
		Index:     1,
		PrevBlock: nil,
		Hash:      services.HashFromString("Genesis"),
	}

	return &BlockChain{
		latestBlock: *genesisBlock,
	}
}

func (blockChain *BlockChain) GetLatestBlock() Block {
	return blockChain.latestBlock
}

func (blockChain *BlockChain) AddBlock(strData string) BlockChain {
	block := CreateBlock(strData, blockChain.latestBlock)
	blockChain.latestBlock = block
	return *blockChain
}

func (blockChain *BlockChain) String() string {
	currentBlock := blockChain.latestBlock
	result := fmt.Sprintf("\nBlockChain blocks: %d\n\n", currentBlock.Index)

	for {
		result += currentBlock.ToString() + "\n"
		if currentBlock.PrevBlock == nil {
			break
		}
		currentBlock = *currentBlock.PrevBlock
	}

	result += "End!"
	return result
}

//recursive display of blockchain
