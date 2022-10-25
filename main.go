package main

import "blockchain/blockchain"

func Init() *blockchain.Blockchain {
	gensisBlock := blockchain.Gensis()
	return &blockchain.Blockchain{
		[]*blockchain.Block{
			gensisBlock,
		},
	}
}

func main() {
	blockChain := Init()
	blockChain.AddBlock("Block 1")
	blockChain.AddBlock("Block 2")
	blockChain.AddBlock("Block 3")
	blockChain.AddBlock("Block 4")

	for _, block := range blockChain.Chain {
		block.Display()
	}

}
