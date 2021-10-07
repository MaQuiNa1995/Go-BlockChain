package model

type blockChain struct {
	Blocks []*block
}

func (chain *blockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockChain() *blockChain {
	return &blockChain{[]*block{Genesis()}}
}
