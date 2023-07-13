package main

import (
	"bufio"
	"fmt"
	"hash-chain/src/models"
	"os"
)

func printBlockChainGenerated() {
	fmt.Println("BlockChain Generated")
}
func getInput() string {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("data:")
	reader.Scan()
	return reader.Text()
}

func main() {
	blockChain := models.GenerateBlockChain()
	printBlockChainGenerated()
	for {
		inputString := getInput()
		if inputString == "exit" {
			break
		}
		blockChain.AddBlock(inputString)
	}
	fmt.Print(blockChain)
}
