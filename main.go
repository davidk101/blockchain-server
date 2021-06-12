package main

import (
	"fmt"
	"github.com/davidk101/golang-blockchain/blockchain"
	"strconv"
)

func main() {

	chain := blockchain.InitBlockChain()

	// hashes for the block will be determined by info inside the block and the previous hash
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// iterate through each block and print the fields in each block
	for _, block := range chain.Blocks {

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block) // pass in block to POW algorithm
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
