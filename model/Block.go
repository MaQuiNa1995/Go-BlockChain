package model

import (
	"bytes"
	"crypto/sha256"
)

type block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *block) CalculateHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *block {

	block := &block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}

	block.CalculateHash()
	return block
}

func Genesis() *block {
	return CreateBlock("Genesis", []byte{})
}
