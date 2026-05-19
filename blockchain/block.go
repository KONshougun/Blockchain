package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type BlockHeader struct {
	Index      uint64
	Timestamp  int64

	PrevHash   [32]byte
	MerkleRoot [32]byte
	Nonce      uint64

	Difficulty uint32
}

type Block struct {
	Header       BlockHeader
	Transactions []Transaction
	Hash         [32]byte
}

type Transaction struct {
	ID        [32]byte
	Timestamp int64

	StudentID string
	CourseID  string
	Grade     uint8

	ProfessorSignature []byte
}

func (b *Block) CalculateHash() [32]byte{
	
	hash := sha256.New()
	fmt.Printf("hash: %v\n", hash)
	return [32]byte{}
}