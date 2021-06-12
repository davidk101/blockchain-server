package blockchain

// blockchain - public DB that doesn't rely on trust between nodes of data
// here, blockchain nodes could produce incorrect/corrupted data and the DB can fix itself

type BlockChain struct {
	Blocks []*Block // array of pointers to Blocks
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

/* unnecessary now since hash derived in pow
// takes in a pointer to a block as a parameter
func (b *Block) DeriveHash() {

	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// trivial hashing algorithm using sha256 temporarily
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
*/

// outputs a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {

	block := &Block{[]byte{}, []byte(data), prevHash, 0} // this is a reference to Block

	// run pow algorithm on each block
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// add a block to the chain
func (chain *BlockChain) AddBlock(data string) {

	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash) // creating the current block
	chain.Blocks = append(chain.Blocks, new) // appending the current block with the previous block

}

// the first block that will be created
func Genesis() *Block {

	return CreateBlock("Genesis", []byte{})

}

// building the init blockchain using the Genesis block
func InitBlockChain() *BlockChain {

	return &BlockChain{[]*Block{Genesis()}} // returning a reference to BlockChain
}
