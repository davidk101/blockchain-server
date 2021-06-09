package main

import (
	"bytes"
	"crypto/sha256"
)

// blockchain - public DB that doesn't rely on trust between nodes of data
// here, blockchain nodes could produce incorrect data and the DB can fix itself

type BlockChain struct {
	blocks []*Block // array of pointers to Blocks
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// takes in a pointer to a block as a parameter
func (b *Block) DeriveHash() {

	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte)
	// trivial hashing algorithm using sha256 temporarily
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// outputs a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {

	block := &Block{[]byte{}, []byte(data), prevHash} // this is a reference to Block
	block.DeriveHash()
	return block
}

// add a block to the chain
func (chain *BlockChain) AddBlock(data string) {

	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash) // creating the current block
	chain.blocks = append(chain.blocks, new) // appending the current block with the previous block

}

// the first block that will be created
func Genesis() *Block {

	return CreateBlock("Genesis", []byte{})

}

// building the init blockchain using the Genesis block
func InitBlockChain() *BlockChain {

	return &BlockChain{[]*Block{Genesis()}} // returning a reference to BlockChain
}

func main() {

	chain := InitBlockChain()

	// hashes for the block will be determined by info inside the block and the previous hash
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

}
