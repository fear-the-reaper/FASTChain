package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	Chain []*Block
}

// Calculating the Hash of the block
func (block *Block) CalcHash() {
	// Joining the hashes
	data := bytes.Join(
		[][]byte{block.Data, block.PrevHash},
		[]byte{},
	)
	// making a hash of the combined Hashes
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func (block *Block) Display() {
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
	fmt.Printf("Data -> %s\nHash -> %x\nHash of PrevBlock %x\n", block.Data, block.Hash, block.PrevHash)
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
	fmt.Printf("\n")
}

// Making a new block with parameters of the data to enter and the hash of the prev block
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
	}
	block.CalcHash()
	return block
}

// Adding a new block to the blockchain
func (blockChain *Blockchain) AddBlock(data string) {
	prevBlock := blockChain.Chain[len(blockChain.Chain)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	blockChain.Chain = append(blockChain.Chain, newBlock)
}

// Makes the blockchain with the gensis block
func Gensis() *Block {
	return NewBlock(
		"Gensis",
		[]byte{},
	)
}

func Init() *Blockchain {
	gensisBlock := Gensis()
	return &Blockchain{
		[]*Block{
			gensisBlock,
		},
	}
}

func main() {
	blockChain := Init()
	blockChain.AddBlock("Block 1")
	blockChain.AddBlock("NOT Block 2")
	blockChain.AddBlock("Block 3")
	blockChain.AddBlock("Block 4")

	for _, block := range blockChain.Chain {
		block.Display()
	}

}
