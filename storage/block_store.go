package storage

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/KONshougun/Blockchain/blockchain"
)

type BlockIndex struct {
	FileID uint32
	Offset int64
	Size   uint32
	Height uint64
}

func AppendBlcok(block blockchain.Block) error {

	//PRENDO IL PREV_BLOCK
	if _, err := os.Stat("storage/block/blk0000.dat"); err == nil {
		return fmt.Errorf("genesis already exists")
	}
	//CONTROLLO LA VALIDICITà DEL NUOVO BLOCCO

	return nil
}

func CreateGenesisBlock() error {
	/*
		if _, err := os.Stat("storage/block/blk0000.dat"); err == nil {
			return fmt.Errorf("genesis already exists")
		}

		var genesisBlock = blockchain.Block{
			Header: blockchain.BlockHeader{
				Index:      0,
				Timestamp:  time.Now().Unix(),
				PrevHash:   [32]byte{},
				Difficulty: 1,
				Nonce:      rand.Uint64(),
			},
		}
		genesisBlock.Hash = genesisBlock.CalculateHash()

		file, err := os.Create("storage/block/blk0000.dat")
		if err != nil {
			return err
		}
		defer file.Close()

		var buf bytes.Buffer

		enc := gob.NewEncoder(&buf)
		err = enc.Encode(genesisBlock)
		if err != nil {
			return err
		}

		blockBytes := buf.Bytes()

		size := uint32(len(blockBytes))

		err = binary.Write(file, binary.LittleEndian, size)
		if err != nil {
			return err
		}

		_, err = file.Write(blockBytes)
		if err != nil {
			return err
		}

		//	ADESSO MODIFICO IL FILE INDEX
		buf.Reset()
		blockIdx := BlockIndex{
			FileID: 0,
			Offset: 0,
			Size:   size,
			Height: 0,
		}
		fileIdx, err := os.OpenFile(
			"storage/index.dat",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644,
		)
		if err != nil {
			return err
		}
		defer fileIdx.Close()
		err = binary.Write(fileIdx, binary.LittleEndian, blockIdx.FileID)
		if err != nil {
			return err
		}

		err = binary.Write(fileIdx, binary.LittleEndian, blockIdx.Offset)
		if err != nil {
			return err
		}

		err = binary.Write(fileIdx, binary.LittleEndian, blockIdx.Size)
		if err != nil {
			return err
		}

		err = binary.Write(fileIdx, binary.LittleEndian, blockIdx.Height)
		if err != nil {
			return err
		}

		return nil
	*/
	return fmt.Errorf("Funzione deprecata")
}