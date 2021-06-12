package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// proof of work concept
// want to secure blockchains to force networks to do computational work to sign block on blockchain

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block   // the block inside the blockchain
	Target *big.Int // the number mentioned in the requirement derived from difficulty
}

// pairing block to target and creating new hash
func NewProof(b *Block) *ProofOfWork {

	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // there are 256 bits in our blockchain hash

	pow := &ProofOfWork{b, target}
	return pow
}

// nonce: counter to ensure data transmitted is unique and not replayed
func (pow *ProofOfWork) InitData(nonce int) []byte {

	data := bytes.Join([][]byte{pow.Block.PrevHash, pow.Block.Data, ToHex(int64(nonce)), ToHex(int64(Difficulty))}, []byte{})

	return data
}

//
func (pow *ProofOfWork) Run() (int, []byte) {

	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {

		data := pow.InitData(nonce) // prepare the byte slice
		hash = sha256.Sum256(data)  // hash the data

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:]) // convert hash to bigInt

		if intHash.Cmp(pow.Target) == -1 {
			break // hash is less than pow Target i.e. we signed the block
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:] // returning a tuple

}

// add Difficulty into slice of bytes
// add nonce into slice of bytes
func ToHex(num int64) []byte {

	buff := new(bytes.Buffer)                        // creating a new bytes buffer
	err := binary.Write(buff, binary.BigEndian, num) // writes the binary representation into buffer
	if err != nil {

		log.Panic(err)
	}

	return buff.Bytes()
}
