package model

type Lottery struct {
	Owner   string `json:"owner"`
	Number  string `json:"number"`
	BuyTime int64  `json:"buy_time"`
}

type Block struct {
	Index             int       `json:"index"`
	Timestamp         int64     `json:"timestamp"`
	Lotteries         []Lottery `json:"lotteries"`
	Nonce             int       `json:"nonce"`
	Hash              string    `json:"hash"`
	PreviousBlockHash string    `json:"previous_block_hash"`
}

type Blockchain struct {
	Chain            []Block   `json:"chain"`
	PendingLotteries []Lottery `json:"pending_lotteries"`
	NetworkNodes     []string  `json:"network_nodes"`
}

type BlockData struct {
	Index     string
	Lotteries []Lottery
}
