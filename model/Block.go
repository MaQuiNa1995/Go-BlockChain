package model

type Block struct {
	Index             int       `json:"index"`
	CreationTime      int64     `json:"creation_time"`
	UpdatedTime       int64     `json:"updated_time"`
	Loteries          []Lottery `json:"lotery_number"`
	Nonce             int       `json:"nonce"`
	Hash              string    `json:"hash"`
	PreviousBlockHash string    `json:"previous_block_hash"`
}
