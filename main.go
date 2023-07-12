package main

import (
	"bufio"
	"fmt"
	"hash-chain/src/models"
	hashGenerator "hash-chain/src/services"
	print "hash-chain/src/services"
	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	// blockChain := make(map[string]string)
	var blockChain []models.Block

	fmt.Println("data:")
	reader.Scan()
	genesisData := models.UserData{
		Data: reader.Text(),
	}

	genesisBlock := models.Block{Data: genesisData.ExportToString(),
		PrevBlockHash: nil,
		Hash:          []byte(hashGenerator.HashFromString(genesisData.ExportToString())),
	}
	blockChain = append(blockChain, genesisBlock)

	for {
		fmt.Println("data:")
		reader.Scan()
		inputData := models.UserData{
			Data: reader.Text(),
		}

		if inputData.ExportToString() == "exit" {
			break
		}

		previousElementIndex := len(blockChain) - 1

		newBlock := models.Block{
			Data:          inputData.ExportToString(),
			PrevBlockHash: blockChain[previousElementIndex].Hash,
			Hash:          []byte(hashGenerator.HashBlock(genesisData.ExportToString(), []byte(blockChain[previousElementIndex].Hash))),
		}
		blockChain = append(blockChain, newBlock)
	}

	print.PrintFormattedBlocks(blockChain)

}
