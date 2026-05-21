package blockchain

func mining(block Block) {
	for {
		block.Header.Nonce++

		hash := block.CalculateHash()

		if HasValidDifficulty(hash, block.Header.Difficulty) {
			block.Hash = hash
			break
		}
	}
}
