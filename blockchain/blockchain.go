package blockchain

//Blockchain type
type Blockchain struct {
	Blocks []*Block
}

//AddBlock has a blockchain Pointer as receiver
//And data as argument
//Adds a new block with that data
func (chain *Blockchain) AddBlock(data string) {
	//Get previous block (we need the hash)
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// Create new block from data and previous blockhash
	// CreateBlock returns a pointer to a Block struct
	// Which are the type of elements in our blockchain array
	newBlock := CreateBlock(data, prevBlock.Hash)
	// Append the new block to the chain
	// Similar to array.push() in javascript
	chain.Blocks = append(chain.Blocks, newBlock)
}

//Genesis spawns a new blockchain with a genesis block
func Genesis() *Blockchain {
	genesisBlock := CreateBlock("genesis", []byte{})
	return &Blockchain{[]*Block{genesisBlock}}
	//Remember CreateBlock returns a pointer
	// ALternatively we can thus do it in one line
	// If we initialize the struct and refer the pointer directly
	// return &Blockchain{[]*Block{CreateBlock("genesis", []byte{})}}
}
