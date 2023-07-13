package models

import (
	"encoding/hex"
	"fmt"
	"hash-chain/src/services"
)

type Block struct {
	Data      BlockData
	Number    int
	PrevBlock *Block
	Hash      []byte
}

func CreateBlock(strData string, prvBlock Block) Block {
	newBlock := &Block{
		Data:      CreateBlockData(strData),
		Number:    prvBlock.Number + 1,
		PrevBlock: &prvBlock,
		Hash:      services.HashBlock(strData, prvBlock.Hash),
	}
	return *newBlock
}

func (block *Block) GetPrevBlock() Block {
	return *block.PrevBlock
}

func (block *Block) ToString() string {
	str := fmt.Sprintf("Block Data: %s\n", block.Data.ExportToString())
	str += fmt.Sprintf("Block Number %d\n", block.Number)
	if block.Number != 1 {
		str += fmt.Sprintf("PreviousBlock Hash: %s\n", hex.EncodeToString(block.PrevBlock.Hash))
	} else {
		str += fmt.Sprintf("PreviousBlock Hash: GenesisBlock\n")

	}
	str += fmt.Sprintf("Block Hash: %s\n\n", hex.EncodeToString(block.Hash))
	return str
}
