package model

type BlockChain struct {
	Blocks []*block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*block{Genesis()}}
}
