package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"time"
)

type BlockHeader struct {
	Index     uint64
	Timestamp int64

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
	ID [32]byte

	StudentID   string
	CourseID    string
	ProfessorID string
	Timestamp   int64

	Ciphertext []byte
	Nonce      []byte

	ProfessorSignature []byte
}

func (tx *Transaction) CalculateHash() [32]byte {
	data := fmt.Appendf(nil, "%x|%s|%s|%s|%d|%x|%x",
		tx.ID,
		tx.StudentID,
		tx.CourseID,
		tx.ProfessorID,
		tx.Timestamp,
		tx.Ciphertext,
		tx.Nonce,
	)

	return sha256.Sum256(data)
}

func (block *Block) CalculateHash() [32]byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, block.Header.Index)
	binary.Write(&buf, binary.BigEndian, block.Header.Timestamp)

	buf.Write(block.Header.PrevHash[:])
	buf.Write(block.Header.MerkleRoot[:])

	binary.Write(&buf, binary.BigEndian, block.Header.Nonce)
	binary.Write(&buf, binary.BigEndian, block.Header.Difficulty)
	return sha256.Sum256(buf.Bytes())
}
func (block *Block) IsHashValid() bool {
	return block.Hash == block.CalculateHash()
}

func HasValidDifficulty(hash [32]byte, difficulty uint32) bool {

	hexHash := fmt.Sprintf("%x", hash)

	for i := range difficulty {
		if hexHash[i] != '0' {
			return false
		}
	}

	return true
}

func MerkleRoot(txs []Transaction) [32]byte {

	if len(txs) == 0 {
		return [32]byte{}
	}

	var hashes [][32]byte

	for _, tx := range txs {
		hashes = append(hashes, tx.CalculateHash())
	}

	for len(hashes) > 1 {

		var next [][32]byte

		for i := 0; i < len(hashes); i += 2 {

			// Se dispari, duplica ultimo hash
			if i+1 >= len(hashes) {
				hashes = append(hashes, hashes[i])
			}

			combined := append(hashes[i][:], hashes[i+1][:]...)

			next = append(next,
				sha256.Sum256(combined),
			)
		}

		hashes = next
	}

	return hashes[0]
}

func ValidateBlock(block Block, prev Block) error {

	//Controllo che il previous hash sia corretto
	if block.Header.PrevHash != prev.Hash {
		return fmt.Errorf("Previous hash errato")
	}

	//Controllo la correttezza dell'hash
	if block.Hash != block.CalculateHash() {
		return fmt.Errorf("Hash non valido")
	}

	//Controllo la validità dei timestamp
	if block.Header.Timestamp <= prev.Header.Timestamp {
		return fmt.Errorf("timestamp non valido")
	}
	if block.Header.Timestamp > time.Now().Unix()+300 {
		return fmt.Errorf("timestamp troppo nel futuro")
	}

	//Controllo la validità dell'hash in base alla difficulty
	if !block.IsHashValid() {
		return fmt.Errorf("Merkle root non valido")
	}

	//Controllo la validità Merkle root
	if block.Header.MerkleRoot != MerkleRoot(block.Transactions) {
		return fmt.Errorf("Merkle root non valido")
	}

	//Controllo la validità dell'hash in base alla difficulty
	if !HasValidDifficulty(block.Hash, block.Header.Difficulty) {
		return fmt.Errorf("Hash non valido rispetto alla difficulty")
	}

	return nil
}

func CreateGenesisBlock() Block {

	var genesisBlock = Block{
		Header: BlockHeader{
			Index:      0,
			Timestamp:  time.Now().Unix(),
			PrevHash:   [32]byte{},
			Difficulty: 1,
			Nonce:      rand.Uint64(),
		},
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()

	return genesisBlock
}