package blockchain

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	b := CreateBlock("hello block", []byte{})
	hash := sha256.Sum256(bytes.Join([][]byte{[]byte("hello block"), []byte{}}, []byte{}))
	if !bytes.Equal(b.Hash[:], hash[:]) {
		t.Errorf("Blockhash is not correct")
	}
}
