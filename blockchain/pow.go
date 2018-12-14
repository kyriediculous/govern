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

//Difficulty determines proof-of-work hash leading 0's
//This simple implementation uses static difficulty

//PoW struct represents our proof-of-work
// The difficulty is encoded as a target which is essentially a 256 bit number.
// Since block hashes are produced by SHA-256, they're also a string of 256 bits.
// If a block candidate's hash interpreted as a number is numerically smaller than the target, the block candidate is a valid block.

type PoW struct {
	Block  *Block
	Target *big.Int
}

// NewProof initiates a blank proof for a specific block
// The proof contains the block and the target
func NewProof(b *Block) *PoW {
	//Initialize target as bigInt
	target := big.NewInt(1)
	//Leftshit bit operation
	target.Lsh(target, uint(256-b.Difficulty))
	//return the pointer
	pow := &PoW{b, target}
	return pow
}

//Work starts the work on a block
func (pow *PoW) Work() (int, []byte) {
	//initiate intHash and hash variables
	var intHash big.Int
	var hash []byte
	//Set nonce to 0 to start the work
	nonce := 0

	//Start incrementing nonce and rehashing the block.
	for nonce < math.MaxInt64 {
		hash = pow.deriveHash(nonce)

		fmt.Printf("\r%x", hash)
		//convert hash into big.Int
		intHash.SetBytes(hash)

		if intHash.Cmp(pow.Target) == -1 {
			//bigInt of blockhash is smaller than target, means valid hash
			break
		} else {
			//Not a valid hash , increment nonce and repeat
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash
}

//Validate the proof
//Take the big int of the block hash
//Compare it to our bigInt target
//Data hash has to be smaller
func (pow *PoW) Validate() bool {
	var intHash big.Int
	hash := pow.deriveHash(pow.Block.Nonce)
	//convert hash to bigInt and compare with target
	return intHash.SetBytes(hash).Cmp(pow.Target) == -1
}

func (pow *PoW) deriveHash(nonce int) []byte {
	//Concat Data and PrevHash
	// Join takes in a multi-dimensional bytes slice
	// And concats the elements
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			toHex(int64(nonce)),
			toHex(int64(pow.Block.Difficulty)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(data)
	return hash[:]
}

func toHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
