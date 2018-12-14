package blockchain

import (
	"bytes"
	"testing"
)

func TestGenesis(t *testing.T) {
	bc := Genesis()
	if !bytes.Equal(bc.Blocks[0].Data[:], []byte("genesis")[:]) {
		t.Errorf("Did not find genesis block")
	}
}

func TestAddBlock(t *testing.T) {
	bc := Genesis()
	bc.AddBlock("second block")
	if !bytes.Equal(bc.Blocks[1].Data[:], []byte("second block")[:]) {
		t.Errorf("Second block not found")
	}
}
