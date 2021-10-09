package main

import (
	"MaQuina1995/blockchain/model"
	"fmt"
	"strconv"
)

func main() {
	chain := model.InitBlockChain()

	for i := 1; i < 4; i++ {
		chain.AddBlock(fmt.Sprintf("%vº Block Despues del Genesis", i))
	}

	for _, block := range chain.Blocks {
		fmt.Printf("Bloque:\n\tHash Anterior: %x\n\tDato Del Bloque: %s\n\tHash: %x\n\n", block.PrevHash, block.Data, block.Hash)

		pow := model.NewProof(block)
		fmt.Printf("\tPrueba de Trabajo: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}
