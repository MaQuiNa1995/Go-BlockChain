package main

import (
	"MaQuina1995/blockchain/model"
	"fmt"
)

func main() {
	chain := model.InitBlockChain()

	for i := 1; i < 4; i++ {
		chain.AddBlock(fmt.Sprintf("%vº Block Despues del Genesis", i))
	}

	for _, block := range chain.Blocks {
		fmt.Printf("Bloque:\n\tPrevious Hash: %x\n\tData in Block: %s\n\tHash: %x\n\n", block.PrevHash, block.Data, block.Hash)
	}
}
