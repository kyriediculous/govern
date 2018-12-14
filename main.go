package main

import (
	"fmt"
	"strconv"

	"./blockchain"
)

func main() {
	chain := blockchain.Genesis()

	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("Valid proof: %s\n", strconv.FormatBool(pow.Validate()))
	}
}
