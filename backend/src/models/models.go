package models

import "github.com/ethereum/go-ethereum/common"

type NFTCollection struct {
	Address common.Address `json:"address"`
	Name    string         `json:"name"`
	Symbol  string         `json:"symbol"`
	TxHash  common.Hash    `json:"txHash"`
}

type NFT struct {
	Collection common.Address `json:"collection"`
	TokenURI   string         `json:"tokenURI"`
	TokenID    string         `json:"tokenID"`
	Owner      common.Address `json:"owner"`
	TxHash     common.Hash    `json:"txHash"`
}
