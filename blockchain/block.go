package blockchain

//Block type
type Block struct {
	Hash       []byte //Blockhash Sha3(Data+PrevHash)
	Data       []byte //Data of the block
	PrevHash   []byte //Previous block hash
	Difficulty int
	Nonce      int
}

//CreateBlock craetes a new block
func CreateBlock(data string, prevHash []byte) *Block {
	//Initialize Block struct
	//And get the pointer
	block := &Block{
		Hash:       []byte{},
		Data:       []byte(data),
		PrevHash:   prevHash,
		Difficulty: 20,
		Nonce:      0,
	}
	//Initialize new proof for current block
	// sets static bigInt target and takes in the block
	pow := NewProof(block)
	nonce, hash := pow.Work()
	block.Nonce = nonce
	block.Hash = hash[:]
	//return the block pointer
	return block
}
