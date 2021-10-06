package model

type Lottery struct {
	Owner   string `json:"playername"`
	Number  string `json:"number"`
	BuyTime int64  `json:"buy_time"`
}
