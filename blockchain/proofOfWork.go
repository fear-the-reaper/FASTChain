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

// What this does:

// TODO: Take data from the block

// TODO: create a counter/random/salt/nounce which starts from 0 and goes to infinity (theoretically)

// TODO: create a Hash of the data + counter/random/salt/nounce

// TODO: check the hash to see if it meets our requirements

// * Requirements
// * The first few bytes of the hash must contain 0s

const DIFFICULTY = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func getBytes(num int64) []byte {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, num); err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()

}

func (proofOfWork *ProofOfWork) Run() (int, []byte) {
	var randNumber big.Int
	var hash [32]byte
	counter := 0

	found := false

	for counter < math.MaxInt64 && !found {
		data := proofOfWork.CombineData(counter)
		hash = sha256.Sum256(data)

		fmt.Printf("Salt: %d | Hash: %x\n", counter, hash)
		randNumber.SetBytes(hash[:])

		if randNumber.Cmp(proofOfWork.Target) == -1 {
			found = true
		} else {
			counter++
		}

	}

	return counter, hash[:]

}

func (proofOfWork *ProofOfWork) CombineData(counter int) []byte {
	data := bytes.Join(
		[][]byte{
			proofOfWork.Block.PrevHash,
			proofOfWork.Block.Data,
			getBytes(int64(counter)),
			getBytes(int64(DIFFICULTY)),
		},
		[]byte{},
	)

	return data

}

func (proofOfWork *ProofOfWork) Validate() bool {
	var gottenHash big.Int

	data := proofOfWork.CombineData(proofOfWork.Block.Nounce)
	hash := sha256.Sum256(data)
	gottenHash.SetBytes(hash[:])

	return gottenHash.Cmp(proofOfWork.Target) == -1

}

func NewProof(block *Block) *ProofOfWork {

	target := big.NewInt(1)
	target.Lsh(target, uint(256-DIFFICULTY))

	fmt.Println("")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("The target for this block is ", target)
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("")

	proofOfWork := &ProofOfWork{
		block,
		target,
	}

	return proofOfWork
}
