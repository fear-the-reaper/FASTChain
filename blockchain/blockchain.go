package blockchain

import (
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nounce   int
}

type Blockchain struct {
	Chain []*Block
}

func (block *Block) Display() {

	proofOfWork := NewProof(block)
	valid := proofOfWork.Validate()
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
	fmt.Printf("Data -> %s\nHash -> %x\nHash of PrevBlock %x\nNounce -> %d\nValid -> %t\n", block.Data, block.Hash, block.PrevHash, block.Nounce, valid)
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
	fmt.Printf("\n")
}

// Making a new block with parameters of the data to enter and the hash of the prev block
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
		0,
	}
	proofOfWork := NewProof(block)
	nounce, hash := proofOfWork.Run()
	block.Hash = hash[:]
	block.Nounce = nounce
	// block.CalcHash()
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
