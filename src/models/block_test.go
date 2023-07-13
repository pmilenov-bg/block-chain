package models

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	prevBlock := Block{
		Data:   CreateBlockData("Previous Block"),
		Number: 1,
		Hash:   []byte("Previous Hash"),
	}

	newBlock := CreateBlock("New Block", prevBlock)

	if newBlock.Data.ExportToString() != "New Block" {
		t.Errorf("Expected Data to be 'New Block', but got '%s'", newBlock.Data.ExportToString())
	}

	if newBlock.Number != 2 {
		t.Errorf("Expected Number to be 2, but got %d", newBlock.Number)
	}

	if newBlock.PrevBlock != &prevBlock {
		t.Errorf("Expected PrevBlock to point to the previous block, but it doesn't match")
	}

	// Test the hash calculation
	expectedHash := "3d51d316bd0e31f285b67c25d09ee7507b83ffbb9e88dd2bef7956a92aa7f80c"
	actualHash := hex.EncodeToString(newBlock.Hash)

	if actualHash != expectedHash {
		t.Errorf("Expected Hash to be '%s', but got '%s'", expectedHash, actualHash)
	}
}

func TestGetPrevBlock(t *testing.T) {
	prevBlock := Block{
		Data:   CreateBlockData("Previous Block"),
		Number: 1,
		Hash:   []byte("Previous Hash"),
	}

	block := Block{
		Data:      CreateBlockData("Block"),
		Number:    2,
		PrevBlock: &prevBlock,
		Hash:      []byte("Block Hash"),
	}

	prevBlockReturned := block.GetPrevBlock()

	if !bytes.Equal(prevBlockReturned.Hash, prevBlock.Hash) {
		t.Errorf("Expected GetPrevBlock() to return the previous block, but it doesn't match")
	}
}

func TestToString(t *testing.T) {
	block := Block{
		Data:   CreateBlockData("Block Data"),
		Number: 1,
		PrevBlock: &Block{
			Data:   CreateBlockData("Previous Block Data"),
			Number: 0,
			Hash:   []byte("Previous Hash"),
		},
		Hash: []byte("Block Hash"),
	}

	expectedString := "Block Data: Block Data\nBlock Number 1\nPreviousBlock Hash: GenesisBlock\nBlock Hash: 426c6f636b2048617368\n\n"
	actualString := block.ToString()

	if actualString != expectedString {
		t.Errorf("Expected ToString() to return:\n%s\n\nBut got:\n%s\n", expectedString, actualString)
	}
}
